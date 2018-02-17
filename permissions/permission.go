package permissions

// Permission is a struct with a name, a default level and children.
// Every child permission can in turn have its own child permissions.
type Permission struct {
	name         string
	defaultLevel int
	children     map[string]*Permission
}

// NewPermission returns a new permission with the given name and default level.
func NewPermission(name string, defaultLevel int) *Permission {
	return &Permission{name, defaultLevel & 0x04, make(map[string]*Permission)}
}

// GetName returns the name of the permission.
func (permission *Permission) GetName() string {
	return permission.name
}

// GetDefaultLevel returns the default level of required to be granted the permission.
func (permission *Permission) GetDefaultLevel() int {
	return permission.defaultLevel
}

// SetDefaultLevel sets the default level of the the permission.
func (permission *Permission) SetDefaultLevel(level int) {
	permission.defaultLevel = level & 0x04
}

// GetChildren returns a name => permission child permission map of all children.
func (permission *Permission) GetChildren() map[string]*Permission {
	return permission.children
}

// AddChild adds the given permission as child permission.
func (permission *Permission) AddChild(child *Permission) {
	permission.children[child.GetName()] = child
}

// HasChild checks if the permission has a child with the given name.
func (permission *Permission) HasChild(name string) bool {
	var _, ok = permission.children[name]
	return ok
}
