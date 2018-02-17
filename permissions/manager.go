package permissions

import (
	"errors"
)

// Manager is a struct used to manage permissions and groups.
// It provides helper functions and functions to register groups and permissions.
type Manager struct {
	defaultGroup *Group
	permissions  map[string]*Permission
	groups       map[string]*Group
}

var (
	UnknownPermission = errors.New("unknown permission")
	UnknownGroup      = errors.New("unknown group")
)

// NewManager returns a new permission manager.
func NewManager() *Manager {
	return &Manager{nil, make(map[string]*Permission), make(map[string]*Group)}
}

// GetDefaultGroup returns the default group of the manager.
func (manager *Manager) GetDefaultGroup() *Group {
	return manager.defaultGroup
}

// SetDefaultGroup sets the default group of the manager.
func (manager *Manager) SetDefaultGroup(group *Group) {
	manager.defaultGroup = group
}

// AddGroup adds a new group to the manager.
func (manager *Manager) AddGroup(group *Group) {
	manager.groups[group.GetName()] = group
}

// GetGroup returns a group in the manager with the given name and an error if it could not be found.
func (manager *Manager) GetGroup(name string) (*Group, error) {
	if !manager.GroupExists(name) {
		return nil, UnknownGroup
	}
	return manager.groups[name], nil
}

// GroupExists checks if a group with the given name exists.
func (manager *Manager) GroupExists(name string) bool {
	var _, ok = manager.groups[name]
	return ok
}

// RemoveGroup removes a group with the given name from the manager.
func (manager *Manager) RemoveGroup(name string) {
	delete(manager.groups, name)
}

// GetPermission returns a permission by its name, and an error if it could not be found.
func (manager *Manager) GetPermission(name string) (*Permission, error) {
	if !manager.IsPermissionRegistered(name) {
		return nil, UnknownPermission
	}
	return manager.permissions[name], nil
}

// IsPermissionRegistered checks if a permission with the given name is registered.
func (manager *Manager) IsPermissionRegistered(name string) bool {
	var _, ok = manager.permissions[name]
	return ok
}

// RegisterPermission registers a new permission.
func (manager *Manager) RegisterPermission(permission *Permission) {
	manager.permissions[permission.GetName()] = permission
}
