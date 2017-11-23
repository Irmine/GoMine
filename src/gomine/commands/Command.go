package commands

import (
	"gomine/interfaces"
	"reflect"
	"strings"
	"strconv"
)

type Command struct {
	name string
	description string
	permission string
	aliases []string
	arguments []interfaces.ICommandArgument
	usage string
}

/**
 * Returns a new base command.
 */
func NewCommand(name string, description string, permission string, aliases []string) *Command {
	return &Command{name: name, permission: permission, aliases: aliases, description: description}
}

/**
 * Returns the usage of this command.
 */
func (command *Command) GetUsage() string {
	command.parseUsage()
	return command.usage
}

/**
 * Returns the command name.
 */
func (command *Command) GetName() string {
	return command.name
}

/**
 * Returns the command description.
 */
func (command *Command) GetDescription() string {
	return command.description
}

/**
 * Sets the description of the command.
 */
func (command *Command) SetDescription(description string) {
	command.description = description
}

/**
 * Sets the permission of the command.
 */
func (command *Command) SetPermission(permission string) {
	command.permission = permission
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
 * Parses the usage into a readable and clear one.
 */
func (command *Command) parseUsage() {
	if command.usage == "" {
		var usage = "Usage: /" + command.GetName() + " "
		for _, argument := range command.GetArguments() {
			if argument.IsOptional() {
				usage += "["
			} else {
				usage += "<"
			}
			var argName = strings.ToLower(reflect.TypeOf(argument).Elem().Name())
			var argType = strings.Replace(strings.Replace(argName, "enum", "", -1), "arg", "", -1)

			usage += argument.GetName() + ": " + argType
			if argument.GetInputAmount() > 1 && argType != "string" {
				usage += "(" + strconv.Itoa(argument.GetInputAmount()) + ")"
			}

			if argument.IsOptional() {
				usage += "]"
			} else {
				usage += ">"
			}
			usage += " "
		}
		command.usage = usage
	}
}

/**
 * Checks and parses the values of a command.
 */
func (command *Command) Parse(sender interfaces.ICommandSender, commandArgs []string, server interfaces.IServer) ([]interfaces.ICommandArgument, bool) {
	var stringIndex = 0
	if len(commandArgs) == 0 {
		if len(command.GetArguments()) == 0 {
			return command.GetArguments(), true
		}
		sender.SendMessage(command.GetUsage())
		return nil, false
	}
	for _, argument := range command.GetArguments() {
		var i = 0
		var output = []string{}

		for i < argument.GetInputAmount() {
			if len(commandArgs) < stringIndex + i + 1 {
				if !argument.IsOptional() {
					sender.SendMessage(command.GetUsage())
					return nil, false
				}
			} else {
				commandArgs[stringIndex + i] = strings.TrimSpace(commandArgs[stringIndex + i])

				if !argument.IsValidValue(commandArgs[stringIndex + i], server) {
					sender.SendMessage(command.GetUsage())
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

		if argument.ShouldMerge() {
			argument.SetOutput(strings.Join(output, " "))
		} else {
			if len(processedOutput) == 1 {
				argument.SetOutput(processedOutput[0])
			} else {
				argument.SetOutput(processedOutput)
			}
		}
	}
	return command.GetArguments(), true
}
