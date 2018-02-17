package permissions

// Permissible is an interface used to satisfy for permission holders.
type Permissible interface {
	HasPermission(string) bool
	RemovePermission(string)
	AddPermission(*Permission)
}
