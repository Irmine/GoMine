package text

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

const (
	Debug      = "[Debug]"
	Info       = "[Info]"
	Notice     = "[Notice]"
	Alert      = "[Alert]"
	Error      = "[Error]"
	Warning    = "[Warning]"
	Critical   = "[Critical]"
	Chat       = "[Chat]"
	StackTrace = "[Stack Trace]"
)

// Logger is a helper for writing log information to multiple
// locations at the same time on a different goroutine.
// Each logger has a prefix, which all messages will be
// prefixed with, and a debug mode, which if turned on will
// write debug messages too.
type Logger struct {
	// Prefix is the prefix of the logger.
	// Every message is prefixed with this string.
	// The prefix is enclosed in brackets, as such: [Prefix]
	Prefix string
	// DebugMode is the debug mode of the logger.
	// If true, writes debug messages.
	DebugMode bool
	// OutputFunctions contains all logger output functions.
	// Every output function gets called once a message gets logged.
	OutputFunctions []func(message []byte)
	// MessageQueue is the queue of messages to the processed.
	// These messages will be continuously processed on a different goroutine.
	MessageQueue chan string

	// waiting and waitRelease are used to manage the waiting state of the logger.
	// Both are used to notify the logger for waiting.
	waiting     bool
	waitRelease chan bool
}

// DefaultLogger is the default GoMine logger.
// It has the prefix `GoMine` and has debug turned off.
// The default logger will write only to Stdout.
var DefaultLogger = NewLogger("GoMine", false)

// init initializes the output of the default logger.
// It writes to Stdout by default.
func init() {
	DefaultLogger.AddOutput(func(message []byte) {
		os.Stdout.Write(message)
	})
}

// NewLogger returns a new logger with the given prefix and debug mode.
// Additional output functions can be added to the logger once an
// instance has been created using this function.
// The logger will be made to process immediately when creating a new logger.
func NewLogger(prefix string, debugMode bool) *Logger {
	logger := &Logger{prefix, debugMode, []func([]byte){}, make(chan string, 128), false, make(chan bool)}
	go logger.process()
	return logger
}

// AddOutput adds a new output function to the logger.
// The function passed will get called with the message
// provided as argument every time a message gets logged.
// Example:
// func(message []byte) { os.Stdout.Write(message) }
func (logger *Logger) AddOutput(f func(message []byte)) {
	logger.OutputFunctions = append(logger.OutputFunctions, f)
}

// Write writes a byte array to the logger.
// All Minecraft colors are first replaced with ANSI colors.
// after which they get added to the message queue.
// The message will then get processed on a different goroutine.
func (logger *Logger) Write(message []byte) {
	logger.MessageQueue <- ColoredString(string(message)).ToANSI() + AnsiReset + "\n"
}

// Write writes a string to the logger.
// All Minecraft colors are first replaced with ANSI colors.
// after which they get added to the message queue.
// The message will then get processed on a different goroutine.
func (logger *Logger) WriteString(message string) {
	logger.MessageQueue <- ColoredString(message).ToANSI() + AnsiReset + "\n"
}

// process continuously processes queued messages in the logger.
// Messages get fetched from the queue as soon as they're added,
// and will be ran through every output function.
func (logger *Logger) process() {
	for {
		if len(logger.MessageQueue) == 0 && logger.waiting {
			logger.waitRelease <- true
			return
		}
		message := "[" + logger.Prefix + "] " + <-logger.MessageQueue
		for _, f := range logger.OutputFunctions {
			f([]byte(message))
		}
	}
}

// Wait waits until the logger is done logging all messages
// currently in the message queue. The curent goroutine will be
// blocked until the logger is done processing all messages,
// and the writing goroutine will be stopped.
// After waiting, the writing process gets restarted.
func (logger *Logger) Wait() {
	logger.waiting = true
	<-logger.waitRelease
	logger.waiting = false
	go logger.process()
}

// Notice logs a notice message.
func (logger *Logger) Notice(messages ...interface{}) {
	logger.WriteString(Yellow + Notice + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

// Debug logs a debug message.
func (logger *Logger) Debug(messages ...interface{}) {
	logger.WriteString(Orange + Debug + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

// Info logs an info message.
func (logger *Logger) Info(messages ...interface{}) {
	logger.WriteString(BrightCyan + Info + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

// Alert logs an alert.
func (logger *Logger) Alert(messages ...interface{}) {
	logger.WriteString(BrightRed + Alert + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

// Warning logs a warning message.
func (logger *Logger) Warning(messages ...interface{}) {
	logger.WriteString(BrightRed + Bold + Warning + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

// Critical logs a critical warning message.
func (logger *Logger) Critical(messages ...interface{}) {
	logger.WriteString(BrightRed + Underlined + Bold + Critical + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

// Error logs an error message.
func (logger *Logger) Error(messages ...interface{}) {
	logger.WriteString(Red + Error + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

// LogChat logs a chat message to the logger.
func (logger *Logger) LogChat(messages ...interface{}) {
	logger.WriteString(BrightCyan + Chat + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

// LogStack logs the stack trace.
func (logger *Logger) LogStack() {
	logger.WriteString(Yellow + StackTrace + " " + string(debug.Stack()))
}

// LogError logs an actual error to the logger.
// A nil error may also be passed,
// which the logger will completely ignore.
func (logger *Logger) LogError(err error) {
	if err == nil {
		return
	}
	logger.Error(err.Error())
}
