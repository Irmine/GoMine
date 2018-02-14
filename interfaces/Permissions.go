package interfaces

type IPermissionManager interface {
	GetServer() IServer
	GetDefaultGroup() IPermissionGroup
	SetDefaultGroup(IPermissionGroup)
	AddGroup(IPermissionGroup) bool
	GroupExists(string) bool
	RemoveGroup(string) bool
	GetGroup(string) (IPermissionGroup, error)
	GetPermission(string) (IPermission, error)
	IsPermissionRegistered(string) bool
	RegisterPermission(IPermission) bool
}

type IPermissionGroup interface {
	GetName() string
	GetPermissions() map[string]IPermission
	HasPermission(string) bool
	AddPermission(IPermission) bool
	RemovePermission(string) bool
	InheritGroup(IPermissionGroup)
}

type IPermission interface {
	GetName() string
	GetDefaultLevel() int
	SetDefaultLevel(int)
	GetChildren() map[string]IPermission
	AddChild(IPermission) bool
	HasChild(string) bool
}

type IPermissible interface {
	HasPermission(string) bool
	AddPermission(IPermission) bool
	RemovePermission(string) bool
}
