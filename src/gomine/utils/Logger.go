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
func (logger *Logger) Log(logLevel string, color string, messages ...interface{}) {
	var params []string
	for range messages {
		params = append(params, "%v")
	}
	var parameterString = strings.Join(params, " ")
	var message = strings.Trim(fmt.Sprintf(parameterString, messages), "[]")

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
func (logger *Logger) Notice(messages ...interface{}) {
	logger.Log(Notice, Yellow, messages)
}

/**
 * Logs a debug message.
 */
func (logger *Logger) Debug(messages ...interface{}) {
	logger.Log(Debug, Orange, messages)
}

/**
 * Logs an info message.
 */
func (logger *Logger) Info(messages ...interface{}) {
	logger.Log(Info, BrightCyan, messages)
}

/**
 * Logs an alert.
 */
func (logger *Logger) Alert(messages ...interface{}) {
	logger.Log(Alert, BrightRed, messages)
}

/**
 * Logs a warning message.
 */
func (logger *Logger) Warning(messages ...interface{}) {
	logger.Log(Warning, BrightRed + Bold, messages)
}

/**
 * Logs a critical warning message.
 */
func (logger *Logger) Critical(messages ...interface{}) {
	logger.Log(Critical, BrightRed + Underlined + Bold, messages)
}

/**
 * Logs an error message.
 */
func (logger *Logger) Error(messages ...interface{}) {
	logger.Log(Error, Red, messages)
}

/**
 * Logs an actual error.
 */
func (logger *Logger) LogError(err error) {
	if err == nil {
		return
	}
	logger.Error(err.Error())
}

/**
 * Writes the given line to the log and appends a new line.
 */
func (logger *Logger) write(line string) {
	logger.file.WriteString(StripAllColors(line + "\n"))
}

/**
 * Synchronizes the file.
 */
func (logger *Logger) Sync() {
	logger.file.Sync()
}

/**
 * Terminates the logger and stops processing the queue.
 */
func (logger *Logger) Terminate() {
	logger.terminated = true
}