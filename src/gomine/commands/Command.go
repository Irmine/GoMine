package commands

import (
	"fmt"
	"gomine/interfaces"
)

type Command struct {
	command interfaces.ICommand
	name string
	permission string
	aliases []string
}

func NewCommand(command interfaces.ICommand, name string, permission string, aliases []string) *Command {
	return &Command{command, name, permission, aliases}
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

func (command *Command) Parse(commandText string) {
	command.command.Execute(commandText)
	fmt.Println("Command executed!")
}
