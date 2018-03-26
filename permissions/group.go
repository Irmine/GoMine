package permissions

// Group is a struct used for basic permission managing.
// Groups can be granted a set of permissions.
type Group struct {
	name        string
	level       int
	permissions map[string]*Permission
}

// NewGroup returns a new group with the given name and permission level.
func NewGroup(name string, level int) *Group {
	return &Group{name, level, make(map[string]*Permission)}
}

// GetName returns the name of the group.
func (group *Group) GetName() string {
	return group.name
}

// GetPermissions returns a name => permission map of all permissions of the group.
func (group *Group) GetPermissions() map[string]*Permission {
	return group.permissions
}

// HasPermission checks if the group has a permission with the name.
func (group *Group) HasPermission(permission string) bool {
	var _, ok = group.permissions[permission]
	return ok
}

// AddPermission adds a permission to the group.
func (group *Group) AddPermission(permission *Permission) {
	group.permissions[permission.GetName()] = permission
}

// RemovePermission removes a permission with the given name from the group.
func (group *Group) RemovePermission(permission string) {
	delete(group.permissions, permission)
}

// InheritGroup inherits all permissions from a group.
func (group *Group) InheritGroup(inheritedGroup *Group) {
	for _, permission := range inheritedGroup.GetPermissions() {
		group.AddPermission(permission)
	}
}
