package p200

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/utils"
	"gomine/net/packets/data"
	"gomine/players/handlers"
)

type DisconnectHandler struct {
	*handlers.PacketHandler
}

func NewDisconnectHandler() DisconnectHandler {
	return DisconnectHandler{handlers.NewPacketHandler()}
}

/**
 * The disconnect handler is a special case. It does not follow the rules of the other handlers, and has no own packet.
 */
func (handler DisconnectHandler) Handle(player interfaces.IPlayer, session *server.Session, server interfaces.IServer) {
	if player == nil {
		return
	}
	if player.GetSession() == nil {
		return
	}

	server.GetPlayerFactory().RemovePlayer(player)

	if player.HasSpawned() {
		for _, online := range server.GetPlayerFactory().GetPlayers() {
			online.SendPlayerList(data.ListTypeRemove, map[string]interfaces.IPlayer{player.GetName(): player})
		}

		player.DespawnFromAll()

		player.Close()
		player.SetSpawned(false)

		server.BroadcastMessage(utils.Yellow + player.GetName() + " has left the server")
	}
}
