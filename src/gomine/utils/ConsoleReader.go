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
func (reader *ConsoleReader) ReadLine(holder interfaces.ICommandHolder) string {
	var command, _ = reader.reader.ReadString('\n')
	command = strings.Trim(command, "\n")

	if command == "" {
		return command
	}
	reader.attemptReadCommand(command, holder)
	return command
}

/**
 * Attempts to execute the command entered in the console.
 */
func (reader *ConsoleReader) attemptReadCommand(commandName string, holder interfaces.ICommandHolder) bool {
	if !holder.IsCommandRegistered(commandName) {
		return false
	}
	var command, _ = holder.GetCommand(commandName)
	command.Execute(commandName)
	return true
}
