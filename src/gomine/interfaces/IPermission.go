package interfaces

type IPermission interface {
	GetName() string
	GetDefaultLevel() int
	SetDefaultLevel(int)
	GetChildren() map[string]IPermission
	AddChild(IPermission) bool
	HasChild(string) bool
}
