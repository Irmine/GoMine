package commands

import (
	"errors"
)

type Manager struct {
	commands map[string]*Command
	aliases  map[string]*Command
}

// NewManager returns a new Manager struct.
func NewManager() *Manager {
	return &Manager{make(map[string]*Command), make(map[string]*Command)}
}

// IsCommandRegistered checks if the command has been registered.
// Also checks for aliases.
func (holder *Manager) IsCommandRegistered(commandName string) bool {
	var _, exists = holder.GetCommand(commandName)
	return exists == nil
}

// DeregisterCommand deregisters a command from the command holder.
// Also deregisters all command aliases.
func (holder *Manager) DeregisterCommand(commandName string) bool {
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
func (holder *Manager) GetCommand(commandName string) (*Command, error) {
	var command, err = holder.GetCommandByName(commandName)
	if err != nil {
		command, err = holder.GetCommandByAlias(commandName)
	}
	return command, err
}

// GetCommandByAlias returns a command by alias, and an error if none was found.
func (holder *Manager) GetCommandByAlias(aliasName string) (*Command, error) {
	if !holder.AliasExists(aliasName) {
		return nil, errors.New("command alias " + aliasName + " not found")
	}
	return holder.aliases[aliasName], nil
}

// GetCommandByName returns a command by name, and an error if none was found.
func (holder *Manager) GetCommandByName(commandName string) (*Command, error) {
	var _, exists = holder.commands[commandName]
	if !exists {
		return nil, errors.New("command " + commandName + " not found")
	}
	return holder.commands[commandName], nil
}

// RegisterCommand registers a command in the command holder with the including aliases.
func (holder *Manager) RegisterCommand(command *Command) {
	holder.commands[command.GetName()] = command
	for _, alias := range command.GetAliases() {
		holder.registerAlias(alias, command)
	}
}

// AliasExists checks if the given alias exists or not.
func (holder *Manager) AliasExists(aliasName string) bool {
	var _, exists = holder.aliases[aliasName]
	return exists
}

// registerAlias registers a new alias for the given command.
func (holder *Manager) registerAlias(aliasName string, command *Command) {
	holder.aliases[aliasName] = command
}

// DeregisterAlias deregisters an alias.
func (holder *Manager) deregisterAlias(aliasName string) {
	delete(holder.aliases, aliasName)
}
