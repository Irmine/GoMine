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
}

/**
 * Returns a new logger with the given prefix and output file.
 */
func NewLogger(prefix string, outputDir string) Logger {
	var path = outputDir + "gomine.log"
	var _, error = os.Stat(path)
	var file, fileError = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if fileError != nil {
		panic(error)
	}

	return Logger{prefix, path, file}
}

/**
 * Logs the given message with the given log level.
 */
func (logger Logger) Log(message string, logLevel string) {
	var line = "[" + logger.prefix + "]"

	switch logLevel {
	case Debug:
		// TODO: Add a configuration file to enable debug.
		break
	default:
		line += "[" + strings.Title(logLevel) + "] "
		break
	}

	line += message

	fmt.Println(line)

	go logger.write(line)
}

/**
 * Writes the given line to the log and appends a new line.
 */
func (logger Logger) write(line string) {
	logger.file.WriteString(line + "\n")
	logger.file.Sync()
}
