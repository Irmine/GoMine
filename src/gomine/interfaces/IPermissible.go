package interfaces

type IPermissible interface {
	HasPermission(string) bool
	AddPermission(IPermission) bool
	RemovePermission(string) bool
}