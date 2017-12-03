package players

import (
	"gomine/interfaces"
	"errors"
	"goraklib/server"
	"strconv"
)

type PlayerFactory struct {
	server interfaces.IServer
	players map[string]interfaces.IPlayer
	playersAddress map[string]interfaces.IPlayer
}

/**
 * Returns a new player factory, used for managing players on the server.
 */
func NewPlayerFactory(server interfaces.IServer) *PlayerFactory {
	return &PlayerFactory{server, make(map[string]interfaces.IPlayer), make(map[string]interfaces.IPlayer)}
}

/**
 * Adds a player to the player factory.
 */
func (factory *PlayerFactory) AddPlayer(player interfaces.IPlayer, session *server.Session) {
	factory.players[player.GetName()] = player
	factory.playersAddress[session.GetAddress() + strconv.Itoa(int(session.GetPort()))] = player
}

/**
 * Returns a player from the player map.
 * If the player does not exist, returns an error.
 */
func (factory *PlayerFactory) GetPlayerByName(name string) (interfaces.IPlayer, error) {
	var player Player
	if _, ok := factory.players[name]; ok {
		return factory.players[name], nil
	}
	return &player, errors.New("player does not exist")
}

/**
 * Returns a player by its GoRakLib session.
 */
func (factory *PlayerFactory) GetPlayerBySession(address string, port uint16) (interfaces.IPlayer, error) {
	var player Player
	var key = address + strconv.Itoa(int(port))

	if _, ok := factory.playersAddress[key]; ok {
		return factory.playersAddress[key], nil
	}
	return &player, errors.New("player with session does not exist")
}

/**
 * Returns all players online in a name => player map.
 */
func (factory *PlayerFactory) GetPlayers() map[string]interfaces.IPlayer {
	return factory.players
}

/**
 * Returns the count of all players online.
 */
func (factory *PlayerFactory) GetPlayerCount() uint {
	return uint(len(factory.players))
}