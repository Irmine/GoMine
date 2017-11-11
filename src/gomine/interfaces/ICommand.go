package interfaces

type ICommand interface {
	GetName() string
	GetPermission() string
	GetAliases() []string
	Execute([]ICommandArgument) bool
	Parse([]string, IServer) ([]ICommandArgument, bool)
}
