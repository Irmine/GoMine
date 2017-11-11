package commands

import (
	"gomine/interfaces"
	"fmt"
)

type Command struct {
	name string
	permission string
	aliases []string
	arguments []interfaces.ICommandArgument
}

/**
 * Returns a new base command.
 */
func NewCommand(name string, permission string, aliases []string) *Command {
	return &Command{name, permission, aliases, []interfaces.ICommandArgument{}}
}

/**
 * Returns the command name.
 */
func (command *Command) GetName() string {
	return command.name
}

/**
 * Returns the command permission string.
 */
func (command *Command) GetPermission() string {
	return command.permission
}

/**
 * Returns the aliases of this command.
 */
func (command *Command) GetAliases() []string {
	return command.aliases
}

/**
 * Returns a slice with all arguments.
 */
func (command *Command) GetArguments() []interfaces.ICommandArgument {
	return command.arguments
}

/**
 * Sets the command arguments.
 */
func (command *Command) SetArguments(arguments []interfaces.ICommandArgument) {
	command.arguments = arguments
}

/**
 * Adds one argument to the command.
 */
func (command *Command) AppendArgument(argument interfaces.ICommandArgument) {
	command.arguments = append(command.arguments, argument)
}

/**
 * Checks and parses the values of a command.
 */
func (command *Command) Parse(commandArgs []string, server interfaces.IServer) ([]interfaces.ICommandArgument, bool) {
	var stringIndex = 0
	if len(commandArgs) == 0 {
		if len(command.GetArguments()) == 0 {
			return command.GetArguments(), true
		}
		return nil, false
	}
	for _, argument := range command.GetArguments() {
		var i = 0
		var output = []string{}

		for i < argument.GetInputAmount() {
			if len(commandArgs) < stringIndex + i + 1 {
				if !argument.IsOptional() {
					fmt.Println("My line is too short")
					return nil, false
				}
			} else {
				if !argument.IsValidValue(commandArgs[stringIndex + i], server) {
					fmt.Println("My value isn't valid")
					return nil, false
				}
				output = append(output, commandArgs[stringIndex + i])
			}
			i++
		}
		stringIndex += i
		var processedOutput = []interface{}{}

		for _, value := range output {
			processedOutput = append(processedOutput, argument.ConvertValue(value, server))
		}

		if len(processedOutput) == 1 {
			argument.SetOutput(processedOutput[0])
		} else {
			argument.SetOutput(processedOutput)
		}
	}
	return command.GetArguments(), true
}
