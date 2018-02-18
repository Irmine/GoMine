package players

import (
	"errors"

	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/goraklib/server"
)

type PlayerFactory struct {
	server         interfaces.IServer
	players        map[string]interfaces.IPlayer
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
	factory.playersAddress[server.GetSessionIndex(session)] = player
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
func (factory *PlayerFactory) GetPlayerBySession(session *server.Session) (interfaces.IPlayer, error) {
	var player Player
	var key = server.GetSessionIndex(session)

	if _, ok := factory.playersAddress[key]; ok {
		return factory.playersAddress[key], nil
	}
	return &player, errors.New("player with session does not exist")
}

/**
 * Checks if a player with the given session exists.
 */
func (factory *PlayerFactory) PlayerExistsBySession(session *server.Session) bool {
	var _, ok = factory.playersAddress[server.GetSessionIndex(session)]
	return ok
}

/**
 * Checks if a player with the given name exists.
 */
func (factory *PlayerFactory) PlayerExists(name string) bool {
	var _, ok = factory.players[name]
	return ok
}

/**
 * Returns all players online in a name => player map.
 */
func (factory *PlayerFactory) GetPlayers() map[string]interfaces.IPlayer {
	return factory.players
}

/**
 * Removes a player from the player factory.
 */
func (factory *PlayerFactory) RemovePlayer(player interfaces.IPlayer) {
	delete(factory.players, player.GetName())
	delete(factory.playersAddress, server.GetSessionIndex(player.GetSession()))
}

/**
 * Returns the count of all players online.
 */
func (factory *PlayerFactory) GetPlayerCount() uint {
	return uint(len(factory.players))
}
