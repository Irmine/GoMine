package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Debug    = "debug"
	Info     = "info"
	Notice   = "notice"
	Alert    = "alert"
	Error    = "error"
	Warning  = "warning"
	Critical = "critical"
	Chat     = "chat"
)

type Logger struct {
	prefix    string
	path      string
	file      *os.File
	debugMode bool

	terminalQueue []string
	fileQueue     []string

	terminated bool
}

// NewLogger returns a new logger with the given prefix and output file.
func NewLogger(prefix string, outputDir string, debugMode bool) *Logger {
	var path = outputDir + "gomine.log"
	var file, fileError = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if fileError != nil {
		panic(fileError)
	}

	var logger = &Logger{prefix, path, file, debugMode, []string{}, []string{}, false}

	go func() {
		var ticker = time.NewTicker(time.Second / 40)
		for range ticker.C {
			if !logger.terminated {
				logger.ProcessQueue(false)
			}
		}
	}()

	return logger
}

// ProcessQueue continuously processes the queue of log messages.
func (logger *Logger) ProcessQueue(force bool) {
	for _, message := range logger.terminalQueue {
		if logger.terminated && !force {
			return
		}

		if len(logger.terminalQueue) > 0 {
			logger.terminalQueue = logger.terminalQueue[1:]
		}

		os.Stdout.Write([]byte(message))
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

// Log logs the given message with the given log level and color.
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

	var line = prefix + color + level + message + AnsiReset + "\n"

	logger.fileQueue = append(logger.fileQueue, line)
	logger.terminalQueue = append(logger.terminalQueue, ConvertMcpeColorsToAnsi(line))
}

// Notice logs a notice message.
func (logger *Logger) Notice(messages ...interface{}) {
	logger.Log(Notice, Yellow, messages)
}

// Debug logs a debug message.
func (logger *Logger) Debug(messages ...interface{}) {
	logger.Log(Debug, Orange, messages)
}

// Info logs an info message.
func (logger *Logger) Info(messages ...interface{}) {
	logger.Log(Info, BrightCyan, messages)
}

// Alert logs an alert.
func (logger *Logger) Alert(messages ...interface{}) {
	logger.Log(Alert, BrightRed, messages)
}

// Warning logs a warning message.
func (logger *Logger) Warning(messages ...interface{}) {
	logger.Log(Warning, BrightRed+Bold, messages)
}

// Critical logs a critical warning message.
func (logger *Logger) Critical(messages ...interface{}) {
	logger.Log(Critical, BrightRed+Underlined+Bold, messages)
}

// Error logs an error message.
func (logger *Logger) Error(messages ...interface{}) {
	logger.Log(Error, Red, messages)
}

// LogChat logs a chat message to the logger.
func (logger *Logger) LogChat(messages ...interface{}) {
	logger.Log(Chat, BrightCyan, messages)
}

// LogError logs an actual error, or nothing if the error is nil.
func (logger *Logger) LogError(err error) {
	if err == nil {
		return
	}
	logger.Error(err.Error())
}

// write writes the given line to the log and appends a new line.
func (logger *Logger) write(line string) {
	logger.file.WriteString(StripAllColors(line))
}

// Sync synchronizes the file.
func (logger *Logger) Sync() {
	logger.file.Sync()
}

// Terminate terminates the logger and stops processing the queue.
func (logger *Logger) Terminate() {
	logger.terminated = true
}
