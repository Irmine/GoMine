package interfaces

type IPermissionManager interface {
	GetServer() IServer
	GetDefaultGroup() IPermissionGroup
	SetDefaultGroup(group IPermissionGroup)
	AddGroup(group IPermissionGroup) bool
	GroupExists(string) bool
	RemoveGroup(string) bool
	GetPermission(string) (IPermission, error)
	IsPermissionRegistered(string) bool
	RegisterPermission(IPermission) bool
}
