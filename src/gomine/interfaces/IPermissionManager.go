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
