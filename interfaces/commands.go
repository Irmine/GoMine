package interfaces

type ICommand interface {
	GetName() string
	GetPermission() string
	GetAliases() []string
	Parse(ICommandSender, []string, IServer) ([]ICommandArgument, bool)
}

type ICommandArgument interface {
	GetName() string
	SetName(string)
	IsOptional() bool
	SetOptional(bool)
	GetInputAmount() int
	SetOutput(interface{})
	GetOutput() interface{}
	IsValidValue(string, IServer) bool
	ConvertValue(string, IServer) interface{}
	ShouldMerge() bool
}

type ICommandHolder interface {
	IsCommandRegistered(string) bool
	UnRegisterCommand(string) bool
	GetCommand(string) (ICommand, error)
	GetCommandByAlias(string) (ICommand, error)
	GetCommandByName(string) (ICommand, error)
	RegisterCommand(ICommand)
	AliasExists(string) bool
}

type ICommandSender interface {
	HasPermission(string) bool
	SendMessage(string)
}
