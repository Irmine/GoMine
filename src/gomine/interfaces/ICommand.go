package interfaces

type ICommand interface {
	GetName() string
	GetPermission() string
	GetAliases() []string
	Execute(ICommandSender, []ICommandArgument) bool
	Parse(ICommandSender, []string, IServer) ([]ICommandArgument, bool)
}
