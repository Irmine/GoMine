package players

import (
	"gomine/interfaces"
)

type Player struct {
	interfaces.IEntity
	playerName  string
	displayName string
}

func NewPlayer() Player {
	return Player{}
}

func (Player *Player) getName() string {
	return Player.playerName
}

func (Player *Player) getDisplayName() string {
	return Player.displayName
}

func (Player *Player) setDisplayName(name string) {
	Player.displayName = name
}
