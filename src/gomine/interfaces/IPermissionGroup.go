package interfaces

type IPermissionGroup interface {
	GetName() string
	GetPermissions() map[string]IPermission
	HasPermission(string) bool
	AddPermission(IPermission) bool
	RemovePermission(string) bool
	InheritGroup(IPermissionGroup)
}