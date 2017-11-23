package players

import (
	"gomine/interfaces"
)

type Player struct {
	playerName  string
	displayName string

	permissions map[string]interfaces.IPermission
	permissionGroup interfaces.IPermissionGroup
}

func NewPlayer(server interfaces.IServer, name string) *Player {
	var player = &Player{}
	player.playerName = name
	player.displayName = name
	player.permissions = make(map[string]interfaces.IPermission)

	player.permissionGroup = server.GetPermissionManager().GetDefaultGroup()

	return player
}

/**
 * Returns the username the player used to join the server.
 */
func (player *Player) getName() string {
	return player.playerName
}

/**
 * Returns the name the player shows in-game.
 */
func (player *Player) getDisplayName() string {
	return player.displayName
}

/**
 * Sets the name other players can see in-game.
 */
func (player *Player) setDisplayName(name string) {
	player.displayName = name
}

/**
 * Returns the permission group this player is in.
 */
func (player *Player) GetPermissionGroup() interfaces.IPermissionGroup {
	return player.permissionGroup
}

/**
 * Checks if this player has a permission.
 */
func (player *Player) HasPermission(permission string) bool {
	if player.GetPermissionGroup().HasPermission(permission) {
		return true
	}
	var _, exists = player.permissions[permission]
	return exists
}

/**
 * Adds a permission to the player.
 * Returns true if a permission with the same name was overwritten.
 */
func (player *Player) AddPermission(permission interfaces.IPermission) bool {
	var hasPermission = player.HasPermission(permission.GetName())

	player.permissions[permission.GetName()] = permission

	return hasPermission
}

/**
 * Deletes a permission from the player.
 * This does not delete the permission from the group the player is in.
 */
func (player *Player) RemovePermission(permission string) bool {
	if !player.HasPermission(permission) {
		return false
	}

	delete(player.permissions, permission)

	return true
}