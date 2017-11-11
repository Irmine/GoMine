package utils

import (
	"bufio"
	"os"
	"gomine/interfaces"
	"strings"
)

type ConsoleReader struct {
	reader *bufio.Reader
}

/**
 * Returns a new Console Reader.
 */
func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{bufio.NewReader(os.Stdin)}
}

/**
 * Reads any commands if entered.
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
	var commandName = args[0]
	var holder = server.GetCommandHolder()

	if !holder.IsCommandRegistered(commandName) {
		return false
	}

	var command, _ = holder.GetCommand(commandName)
	var parsedInput, valid = command.Parse(args[1:], server)

	if valid {
		command.Execute(server, parsedInput)
	}
	return true
}
