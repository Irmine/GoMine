package commands

import (
	"errors"

	"github.com/irmine/gomine/interfaces"
)

type CommandHolder struct {
	commands map[string]interfaces.ICommand
	aliases  map[string]interfaces.ICommand
}

// NewCommandHolder returns a new CommandHolder struct.
func NewCommandHolder() *CommandHolder {
	var holder = &CommandHolder{}
	holder.commands = make(map[string]interfaces.ICommand)
	holder.aliases = make(map[string]interfaces.ICommand)
	return holder
}

// IsCommandRegistered checks if the command has been registered.
// Also checks for aliases.
func (holder *CommandHolder) IsCommandRegistered(commandName string) bool {
	var _, exists = holder.GetCommand(commandName)
	return exists == nil
}

// DeregisterCommand deregisters a command from the command holder.
// Also deregisters all command aliases.
func (holder *CommandHolder) DeregisterCommand(commandName string) bool {
	if !holder.IsCommandRegistered(commandName) {
		return false
	}
	var command, _ = holder.GetCommand(commandName)

	for _, alias := range command.GetAliases() {
		holder.deregisterAlias(alias)
	}
	delete(holder.commands, commandName)
	return true
}

// GetCommand returns a command regardless whether it's an alias or the command name, or an error if none was found.
func (holder *CommandHolder) GetCommand(commandName string) (interfaces.ICommand, error) {
	var command, err = holder.GetCommandByName(commandName)
	if err != nil {
		command, err = holder.GetCommandByAlias(commandName)
	}
	return command, err
}

// GetCommandByAlias returns a command by alias, and an error if none was found.
func (holder *CommandHolder) GetCommandByAlias(aliasName string) (interfaces.ICommand, error) {
	var command interfaces.ICommand
	if !holder.AliasExists(aliasName) {
		return command, errors.New("command alias " + aliasName + " not found")
	}
	command = holder.aliases[aliasName]
	return command, nil
}

// GetCommandByName returns a command by name, and an error if none was found.
func (holder *CommandHolder) GetCommandByName(commandName string) (interfaces.ICommand, error) {
	var command interfaces.ICommand
	var _, exists = holder.commands[commandName]
	if !exists {
		return command, errors.New("command " + commandName + " not found")
	}
	return holder.commands[commandName], nil
}

// RegisterCommand registers a command in the command holder with the including aliases.
func (holder *CommandHolder) RegisterCommand(command interfaces.ICommand) {
	holder.commands[command.GetName()] = command
	for _, alias := range command.GetAliases() {
		holder.registerAlias(alias, command)
	}
}

// AliasExists checks if the given alias exists or not.
func (holder *CommandHolder) AliasExists(aliasName string) bool {
	var _, exists = holder.aliases[aliasName]
	return exists
}

// registerAlias registers a new alias for the given command.
func (holder *CommandHolder) registerAlias(aliasName string, command interfaces.ICommand) {
	holder.aliases[aliasName] = command
}

// DeregisterAlias deregisters an alias.
func (holder *CommandHolder) deregisterAlias(aliasName string) {
	delete(holder.aliases, aliasName)
}
