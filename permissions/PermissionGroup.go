package permissions

import (
	"github.com/irmine/gomine/interfaces"
)

type PermissionGroup struct {
	name        string
	level       int
	permissions map[string]interfaces.IPermission
}

/**
 * Creates a new permission group.
 */
func NewPermissionGroup(name string, level int) *PermissionGroup {
	return &PermissionGroup{name, level, make(map[string]interfaces.IPermission)}
}

/**
 * Returns the name of this group.
 */
func (group *PermissionGroup) GetName() string {
	return group.name
}

/**
 * Returns all permissions added to this group.
 */
func (group *PermissionGroup) GetPermissions() map[string]interfaces.IPermission {
	return group.permissions
}

/**
 * Checks if the group has the given permission.
 */
func (group *PermissionGroup) HasPermission(permission string) bool {
	var _, ok = group.permissions[permission]
	return ok
}

/**
 * Adds a new permission to this group.
 * Returns true if the permission was overwritten.
 */
func (group *PermissionGroup) AddPermission(permission interfaces.IPermission) bool {
	var hasPermission = group.HasPermission(permission.GetName())

	group.permissions[permission.GetName()] = permission

	return hasPermission
}

/**
 * Removes an existing permission from this group.
 * Returns true if the removal was successful.
 */
func (group *PermissionGroup) RemovePermission(permission string) bool {
	var hasPermission = group.HasPermission(permission)

	delete(group.permissions, permission)

	return hasPermission
}

/**
 * Inherits all permissions from the given group.
 */
func (group *PermissionGroup) InheritGroup(inheritedGroup interfaces.IPermissionGroup) {
	for _, permission := range inheritedGroup.GetPermissions() {
		group.AddPermission(permission)
	}
}
