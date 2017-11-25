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

	return &Logger{prefix, path, file, debugMode}
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

	var line = prefix + level + message

	go fmt.Println(prefix + color + level + message + AnsiReset)
	go logger.write(line)
}

/**
 * Logs a notice message.
 */
func (logger *Logger) Notice(message string) {
	logger.Log(message, Notice, AnsiYellow)
}

/**
 * Logs a debug message.
 */
func (logger *Logger) Debug(message string) {
	logger.Log(message, Debug, AnsiBrightYellow)
}

/**
 * Logs an info message.
 */
func (logger *Logger) Info(message string) {
	logger.Log(message, Info, AnsiBrightCyan)
}

/**
 * Logs an alert.
 */
func (logger *Logger) Alert(message string) {
	logger.Log(message, Alert, AnsiBrightRed)
}

/**
 * Logs a warning message.
 */
func (logger *Logger) Warning(message string) {
	logger.Log(message, Warning, AnsiBrightRed + AnsiBold)
}

/**
 * Logs a critical warning message.
 */
func (logger *Logger) Critical(message string) {
	logger.Log(message, Critical, AnsiBrightRed + AnsiUnderlined + AnsiBold)
}

/**
 * Logs an error.
 */
func (logger *Logger) Error(message string) {
	logger.Log(message, Error, AnsiRed)
}

/**
 * Writes the given line to the log and appends a new line.
 */
func (logger *Logger) write(line string) {
	logger.file.WriteString(line + "\n")
	logger.file.Sync()
}
