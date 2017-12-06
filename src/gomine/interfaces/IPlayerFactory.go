package interfaces

import "goraklib/server"

type IPlayerFactory interface {
	AddPlayer(IPlayer, *server.Session)
	GetPlayers() map[string]IPlayer
	GetPlayerByName(string) (IPlayer, error)
	GetPlayerBySession(string, uint16) (IPlayer, error)
	GetPlayerCount() uint
}