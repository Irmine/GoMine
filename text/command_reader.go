package text

import (
	"bufio"
	"io"
	"strings"
)

// CommandReader implements command reading from io.Readers.
// CommandReader continuously processes incoming commands,
// and executes all associated functions with it.
type CommandReader struct {
	// reader is the bufio.Reader encapsulating the input reader.
	reader *bufio.Reader
	// LineReadFunctions are all line read functions.
	// These functions get executed every time a line gets read.
	LineReadFunctions []func(line string)
}

// NewCommandReader returns a new CommandReader.
// The input io.Reader is encapsulated by a bufio.Reader
// and further used to continuously read from.
func NewCommandReader(inputReader io.Reader) *CommandReader {
	reader := &CommandReader{bufio.NewReader(inputReader), []func(string){}}
	go func() {
		for {
			reader.readLine()
		}
	}()
	return reader
}

// AddReadFunc adds a new line read function to the command reader.
// The function passed will get called with the line read as argument,
// every time a command gets read from the input reader.
// Example:
// func(line string) { os.Stdout.Write([]byte("You wrote: " + line)) }
func (reader *CommandReader) AddReadFunc(outputFunc func(string)) {
	reader.LineReadFunctions = append(reader.LineReadFunctions, outputFunc)
}

// readLine continuously reads lines from the input reader.
// Every time a line gets read from the input reader,
// all LineReadFunctions are executed with the line read.
func (reader *CommandReader) readLine() {
	command, _ := reader.reader.ReadString('\n')
	command = strings.Trim(command, "\n")
	for _, f := range reader.LineReadFunctions {
		f(command)
	}
}
