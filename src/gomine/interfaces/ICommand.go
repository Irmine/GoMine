package interfaces

type ICommand interface {
	GetName() string
	GetPermission() string
	GetAliases() []string
	Execute(string) bool
	Parse(string) (string, bool)
}
