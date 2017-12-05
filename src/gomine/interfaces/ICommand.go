package interfaces

type ICommand interface {
	GetName() string
	GetPermission() string
	GetAliases() []string
	Parse(ICommandSender, []string, IServer) ([]ICommandArgument, bool)
}
