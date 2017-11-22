package players

import (
	"gomine/interfaces"
	"gomine/permissions"
)

type Player struct {
	playerName  string
	displayName string

	permissions map[string]interfaces.IPermission
	permissionGroup permissions.PermissionGroup
}

func NewPlayer(name string) *Player {
	var player = &Player{}
	player.playerName = name
	player.displayName = name
	player.permissions = make(map[string]interfaces.IPermission)
	//player.permissionGroup = PermissionMan
	return player
}

func (player *Player) getName() string {
	return player.playerName
}

func (player *Player) getDisplayName() string {
	return player.displayName
}

func (player *Player) setDisplayName(name string) {
	player.displayName = name
}

func (player *Player) HasPermission(permission string) bool {

}

func (player *Player) AddPermission(permission interfaces.IPermission) bool {

}

func (player *Player) RemovePermission(permission string) bool {

}
