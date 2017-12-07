package utils

import (
	"os"
	"fmt"
	"strings"
)

const (
	Debug    = "debug"
	Info     = "info"
	Notice   = "notice"
	Alert    = "alert"
	Error	 = "error"
	Warning  = "warning"
	Critical = "critical"
)

type Logger struct {
	prefix string
	path   string
	file   *os.File
	debugMode bool

	terminalQueue []string
	fileQueue []string

	terminated bool
}

/**
 * Returns a new logger with the given prefix and output file.
 */
func NewLogger(prefix string, outputDir string, debugMode bool) *Logger {
	var path = outputDir + "gomine.log"
	var file, fileError = os.OpenFile(path, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

	if fileError != nil {
		panic(fileError)
	}

	var logger = &Logger{prefix, path, file, debugMode, []string{}, []string{}, false}

	go func() {
		for !logger.terminated {
			logger.ProcessQueue(false)
		}
	}()

	return logger
}

/**
 * Continuously processes the queue of log messages.
 */
func (logger *Logger) ProcessQueue(force bool) {
	for _, message := range logger.terminalQueue {
		if logger.terminated && !force {
			return
		}

		if len(logger.terminalQueue) > 0 {
			logger.terminalQueue = logger.terminalQueue[1:]
		}

		fmt.Println(message)
	}

	for _, message := range logger.fileQueue {
		if logger.terminated && !force {
			return
		}

		if len(logger.fileQueue) > 0 {
			logger.fileQueue = logger.fileQueue[1:]
		}

	 	logger.write(message)
	}
}

/**
 * Logs the given message with the given log level and color.
 */
func (logger *Logger) Log(message string, logLevel string, color string) {
	if logLevel == Debug && !logger.debugMode {
		return
	}

	var prefix = "[" + logger.prefix + "]"
	var level = "[" + strings.Title(logLevel) + "] "

	var line = prefix + color + level + message + AnsiReset

	logger.fileQueue = append(logger.fileQueue, line)
	logger.terminalQueue = append(logger.terminalQueue, ConvertMcpeColorsToAnsi(line))
}

/**
 * Logs a notice message.
 */
func (logger *Logger) Notice(message string) {
	logger.Log(message, Notice, Yellow)
}

/**
 * Logs a debug message.
 */
func (logger *Logger) Debug(message string) {
	logger.Log(message, Debug, Orange)
}

/**
 * Logs an info message.
 */
func (logger *Logger) Info(message string) {
	logger.Log(message, Info, BrightCyan)
}

/**
 * Logs an alert.
 */
func (logger *Logger) Alert(message string) {
	logger.Log(message, Alert, BrightRed)
}

/**
 * Logs a warning message.
 */
func (logger *Logger) Warning(message string) {
	logger.Log(message, Warning, BrightRed + Bold)
}

/**
 * Logs a critical warning message.
 */
func (logger *Logger) Critical(message string) {
	logger.Log(message, Critical, BrightRed + Underlined + Bold)
}

/**
 * Logs an error.
 */
func (logger *Logger) Error(message string) {
	logger.Log(message, Error, Red)
}

/**
 * Writes the given line to the log and appends a new line.
 */
func (logger *Logger) write(line string) {
	logger.file.WriteString(StripAllColors(line + "\n"))
	logger.file.Sync()
}

/**
 * Terminates the logger and stops processing the queue.
 */
func (logger *Logger) Terminate() {
	logger.terminated = true
}