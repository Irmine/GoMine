package permissions

import (
	"gomine/interfaces"
)

type Permission struct {
	name string
	defaultLevel int
	children map[string]interfaces.IPermission
}

/**
 * Returns a new permission with the given name and default level.
 */
func NewPermission(name string, defaultLevel int) *Permission {
	return &Permission{name, defaultLevel & 0x04, make(map[string]interfaces.IPermission)}
}

/**
 * Returns the name of this permission.
 * Eg. "gomine.command.stop"
 */
func (permission *Permission) GetName() string {
	return permission.name
}

/**
 * Returns the default level a permissible requires to be granted this permission.
 */
func (permission *Permission) GetDefaultLevel() int {
	return permission.defaultLevel
}

/**
 * Sets the default level a permissible requires to be granted this permission.
 */
func (permission *Permission) SetDefaultLevel(level int) {
	permission.defaultLevel = level & 0x04
}

/**
 * Returns a map containing all children permissions of this permission with their according names.
 */
func (permission *Permission) GetChildren() map[string]interfaces.IPermission {
	return permission.children
}

/**
 * Adds a new child to this permission.
 * Returns true if a child with the same name was overwritten.
 */
func (permission *Permission) AddChild(child interfaces.IPermission) bool {
	var hasChild = permission.HasChild(child.GetName())

	permission.children[child.GetName()] = child

	return hasChild
}

/**
 * Returns if the permission has a child with the given name.
 */
func (permission *Permission) HasChild(name string) bool {
	var _, ok = permission.children[name]

	return ok
}