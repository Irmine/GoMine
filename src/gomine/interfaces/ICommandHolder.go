package interfaces

type ICommandHolder interface {
	IsCommandRegistered(commandName string) bool
	UnRegisterCommand(commandName string) bool
	GetCommand(commandName string) (ICommand, error)
	GetCommandByAlias(aliasName string) (ICommand, error)
	GetCommandByName(commandName string) (ICommand, error)
	RegisterCommand(command ICommand)
	AliasExists(aliasName string) bool
}
