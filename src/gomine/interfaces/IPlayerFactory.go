package interfaces

import (
	"goraklib/server"
)

type IPlayerFactory interface {
	AddPlayer(IPlayer, *server.Session)
	GetPlayers() map[string]IPlayer
	GetPlayerByName(string) (IPlayer, error)
	GetPlayerBySession(*server.Session) (IPlayer, error)
	GetPlayerCount() uint
	RemovePlayer(player IPlayer)
}