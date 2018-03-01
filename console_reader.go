package gomine

import (
	"bufio"
	"os"
	"strings"
)

type ConsoleReader struct {
	reader  *bufio.Reader
	reading bool
}

// NewConsoleReader returns a new Console Reader.
func NewConsoleReader(server *Server) *ConsoleReader {
	var reader = &ConsoleReader{bufio.NewReader(os.Stdin), false}
	reader.StartReading()

	go func() {
		for {
			if reader.reading {
				reader.ReadLine(server)
			}
		}
	}()

	return reader
}

// StartReading makes the console reader start reading.
func (reader *ConsoleReader) StartReading() {
	reader.reading = true
}

// StopReading makes the console reader stop reading.
func (reader *ConsoleReader) StopReading() {
	reader.reading = false
}

// IsReading checks if the console reader is currently reading.
func (reader *ConsoleReader) IsReading() bool {
	return reader.reading
}

// ReadLine reads any commands if entered.
// Reading lines is blocking, and other goroutines should always be used.
func (reader *ConsoleReader) ReadLine(server *Server) string {
	var command, _ = reader.reader.ReadString('\n')
	command = strings.Trim(command, "\n")

	if command == "" {
		return command
	}
	reader.attemptReadCommand(command, server)
	return command
}

// attemptReadCommand attempts to execute the command entered in the console.
func (reader *ConsoleReader) attemptReadCommand(commandText string, server *Server) bool {
	var args = strings.Split(commandText, " ")
	var commandName = args[0]
	var i = 1
	for !server.GetCommandManager().IsCommandRegistered(commandName) {
		if i == len(args) {
			break
		}
		commandName += " " + args[i]
		i++
	}

	var manager = server.GetCommandManager()

	if !manager.IsCommandRegistered(commandName) {
		server.GetLogger().Error("Command could not be found.")
		return false
	}
	args = args[i:]

	var command, _ = manager.GetCommand(commandName)
	command.Execute(server, args)

	return true
}
