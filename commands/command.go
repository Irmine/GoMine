package commands

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/irmine/gomine/utils"
	"github.com/irmine/gomine/commands/arguments"
)

type Command struct {
	name              string
	description       string
	permission        string
	aliases           []string
	arguments         []*arguments.Argument
	argumentTypes     []string
	usage             string
	permissionExempt  bool
	executionFunction interface{}
}

// NewCommand returns a new command with the given command function.
// The permission used in the command should be registered in order to get correct output.
func NewCommand(name string, description string, permission string, aliases []string, function interface{}) *Command {
	if reflect.TypeOf(function).Kind() != reflect.Func {
		function = func() {}
	}
	return &Command{name: name, permission: permission, aliases: aliases, description: description, executionFunction: function}
}

// GetUsage returns the usage of this command.
// The usage will get parsed if it had not yet been.
func (command *Command) GetUsage() string {
	command.parseUsage()
	return command.usage
}

// ExemptFromPermissionCheck sets the command exempted from permission checking, allowing anybody to use it.
func (command *Command) ExemptFromPermissionCheck(value bool) {
	command.permissionExempt = value
}

// IsPermissionChecked checks if the user of this command is checked for the adequate permission.
func (command *Command) IsPermissionChecked() bool {
	return !command.permissionExempt
}

// GetName returns the command name.
func (command *Command) GetName() string {
	return command.name
}

// GetDescription returns the command description.
func (command *Command) GetDescription() string {
	return command.description
}

// SetDescription sets the description of the command.
func (command *Command) SetDescription(description string) {
	command.description = description
}

// SetPermission sets the permission of the command.
func (command *Command) SetPermission(permission string) {
	command.permission = permission
}

// GetPermission returns the command permission string.
func (command *Command) GetPermission() string {
	return command.permission
}

// GetAliases returns the aliases of this command.
func (command *Command) GetAliases() []string {
	return command.aliases
}

// GetArguments returns a slice with all arguments.
func (command *Command) GetArguments() []*arguments.Argument {
	return command.arguments
}

// SetArguments sets the command arguments.
func (command *Command) SetArguments(arguments []*arguments.Argument) {
	command.arguments = arguments
}

// AppendArgument adds one argument to the command.
func (command *Command) AppendArgument(argument *arguments.Argument) {
	command.argumentTypes = append(command.argumentTypes, reflect.TypeOf(argument.GetOutput()).Name())

	command.arguments = append(command.arguments, argument)
}

// parseUsage parses the usage into a readable and clear one.
func (command *Command) parseUsage() {
	if command.usage == "" {
		var usage = utils.Yellow + "Usage: /" + command.GetName() + " "
		for index, argument := range command.GetArguments() {
			if argument.IsOptional() {
				usage += "["
			} else {
				usage += "<"
			}

			usage += argument.GetName() + ": " + command.argumentTypes[index]
			if argument.GetInputAmount() > 1 && command.argumentTypes[index] != "string" {
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

// Execute executes the command with the given sender and command arguments.
func (command *Command) Execute(sender Sender, commandArgs []string) {
	if _, ok := command.parse(sender, commandArgs); !ok {
		return
	}
	command.parseArgsAndExecute(sender)
}

// Parse checks and parses the values of a command.
func (command *Command) parse(sender Sender, commandArgs []string) ([]*arguments.Argument, bool) {
	if command.IsPermissionChecked() && !sender.HasPermission(command.GetPermission()) {
		sender.SendMessage("You do not have permission to execute this command.")
		return []*arguments.Argument{}, false
	}

	var stringIndex = 0
	if len(commandArgs) == 0 {
		if len(command.GetArguments()) == 0 {
			return command.GetArguments(), true
		}
		sender.SendMessage(command.GetUsage())
		return nil, false
	}
	for _, argument := range command.arguments {
		var i = 0
		var output []string

		for i < argument.GetInputAmount() {
			if len(commandArgs) < stringIndex+i+1 {
				if !argument.IsOptional() {
					sender.SendMessage(command.GetUsage())
					return nil, false
				}
			} else {
				commandArgs[stringIndex+i] = strings.TrimSpace(commandArgs[stringIndex+i])

				if !argument.IsValidValue(commandArgs[stringIndex+i]) {
					sender.SendMessage(command.GetUsage())
					return nil, false
				}
				output = append(output, commandArgs[stringIndex+i])
			}
			i++
		}
		stringIndex += i
		var processedOutput []interface{}

		for _, value := range output {
			processedOutput = append(processedOutput, argument.ConvertValue(value))
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

// ParseArgsAndExecute parses the arguments into an output able to be typed against.
// After parsing, the command gets called.
func (command *Command) parseArgsAndExecute(sender Sender) {
	var method = reflect.ValueOf(command.executionFunction)
	var input = make([]reflect.Value, method.Type().NumIn())

	var argOffset = 0
	for i := 0; i < method.Type().NumIn(); i++ {

		if method.Type().In(i).String() == "commands.Sender" {
			input[i] = reflect.ValueOf(sender)
			continue
		}

		input[i] = reflect.ValueOf(command.arguments[argOffset].GetOutput())
		argOffset++
	}

	method.Call(input)
}
