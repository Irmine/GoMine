package players

import (
	"gomine/interfaces"
)

type Player struct {
	playerName  string
	displayName string

	permissions map[string]interfaces.IPermission
	permissionGroup interfaces.IPermissionGroup

	server interfaces.IServer

	language string

	uuid string
	xuid string
	clientId int

	viewDistance uint
}

func NewPlayer(server interfaces.IServer, name string, uuid string, xuid string, clientId int) *Player {
	var player = &Player{}
	player.playerName = name
	player.displayName = name

	player.uuid = uuid
	player.xuid = xuid
	player.clientId = clientId

	player.permissions = make(map[string]interfaces.IPermission)
	player.permissionGroup = server.GetPermissionManager().GetDefaultGroup()

	player.server = server

	return player
}

/**
 * Returns the UUID of this player.
 */
func (player *Player) GetUUID() string {
	return player.uuid
}

/**
 * Returns the XUID of this player.
 */
func (player *Player) GetXUID() string {
	return player.xuid
}

/**
 * Sets the language (locale) of this player.
 */
func (player *Player) SetLanguage(language string) {
	player.language = language
}

/**
 * Returns the language (locale) of this player.
 */
func (player *Player) GetLanguage() string {
	return player.language
}

/**
 * Returns the client ID of this player.
 */
func (player *Player) GetClientId() int {
	return player.clientId
}

/**
 * Sets the view distance of this player.
 */
func (player *Player) SetViewDistance(distance uint) {
	player.viewDistance = distance
}

/**
 * Returns the view distance of this player.
 */
func (player *Player) GetViewDistance() uint {
	return player.viewDistance
}

/**
 * Returns the main server.
 */
func (player *Player) GetServer() interfaces.IServer {
	return player.server
}

/**
 * Returns the username the player used to join the server.
 */
func (player *Player) GetName() string {
	return player.playerName
}

/**
 * Returns the name the player shows in-game.
 */
func (player *Player) GetDisplayName() string {
	return player.displayName
}

/**
 * Sets the name other players can see in-game.
 */
func (player *Player) SetDisplayName(name string) {
	player.displayName = name
}

/**
 * Returns the permission group this player is in.
 */
func (player *Player) GetPermissionGroup() interfaces.IPermissionGroup {
	return player.permissionGroup
}

/**
 * Sets the permission group of this player.
 */
func (player *Player) SetPermissionGroup(group interfaces.IPermissionGroup) {
	player.permissionGroup = group
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