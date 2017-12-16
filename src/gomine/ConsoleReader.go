package gomine

import (
	"bufio"
	"os"
	"gomine/interfaces"
	"strings"
	"gomine/commands"
)

type ConsoleReader struct {
	reader *bufio.Reader
	reading bool
}

/**
 * Returns a new Console Reader.
 */
func NewConsoleReader(server interfaces.IServer) *ConsoleReader {
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

/**
 * Makes the console reader start reading.
 */
func (reader *ConsoleReader) StartReading() {
	reader.reading = true
}

/**
 * Makes the console reader stop reading.
 */
func (reader *ConsoleReader) StopReading() {
	reader.reading = false
}

/**
 * Checks if the console reader is currently reading.
 */
func (reader *ConsoleReader) IsReading() bool {
	return reader.reading
}

/**
 * Reads any commands if entered.
 * Reading lines is blocking, and other goroutines should always be used.
 */
func (reader *ConsoleReader) ReadLine(server interfaces.IServer) string {
	var command, _ = reader.reader.ReadString('\n')
	command = strings.Trim(command, "\n")

	if command == "" {
		return command
	}
	reader.attemptReadCommand(command, server)
	return command
}

/**
 * Attempts to execute the command entered in the console.
 */
func (reader *ConsoleReader) attemptReadCommand(commandText string, server interfaces.IServer) bool {
	var args = strings.Split(commandText, " ")

	var commandName = strings.TrimSpace(args[0])
	var holder = server.GetCommandHolder()

	if !holder.IsCommandRegistered(commandName) {
		server.GetLogger().Error("Command could not be found.")
		return false
	}

	var command, _ = holder.GetCommand(commandName)
	var parsedInput, valid = command.Parse(server, args[1:], server)

	if valid {
		commands.ParseIntoInputAndExecute(server, command, parsedInput)
	}
	return true
}