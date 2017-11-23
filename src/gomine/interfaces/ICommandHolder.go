package interfaces

type ICommandHolder interface {
	IsCommandRegistered(string) bool
	UnRegisterCommand(string) bool
	GetCommand(string) (ICommand, error)
	GetCommandByAlias(string) (ICommand, error)
	GetCommandByName(string) (ICommand, error)
	RegisterCommand(ICommand)
	AliasExists(string) bool
}
