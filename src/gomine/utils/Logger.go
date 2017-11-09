package utils

import (
	"os"
	"fmt"
	"strings"
)

const (
	Debug    = "debug"
	Info     = "info"
	Alert    = "alert"
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
 * Logs the given message with the given log level.
 */
func (logger Logger) Log(message string, logLevel string) {
	if logLevel == Debug && !logger.debugMode {
		return
	}

	var line = "[" + logger.prefix + "][" + strings.Title(logLevel) + "] " + message
	fmt.Println(line)

	go logger.write(line)
}

/**
 * Logs a debug message.
 */
func (logger Logger) Debug(message string) {
	logger.Log(message, Debug)
}

/**
 * Logs an info message.
 */
func (logger Logger) Info(message string) {
	logger.Log(message, Info)
}

/**
 * Logs an alert.
 */
func (logger Logger) Alert(message string) {
	logger.Log(message, Alert)
}

/**
 * Logs a warning message.
 */
func (logger Logger) Warning(message string) {
	logger.Log(message, Warning)
}

/**
 * Logs a critical warning message.
 */
func (logger Logger) Critical(message string) {
	logger.Log(message, Critical)
}

/**
 * Writes the given line to the log and appends a new line.
 */
func (logger Logger) write(line string) {
	logger.file.WriteString(line + "\n")
	logger.file.Sync()
}
