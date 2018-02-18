package p200

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/players/handlers"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/goraklib/server"
)

// The disconnect handler is a special case. It does not follow the rules of the other handlers, and has no own packet.
type DisconnectHandler struct {
	*handlers.PacketHandler
}

func NewDisconnectHandler() DisconnectHandler {
	return DisconnectHandler{handlers.NewPacketHandler()}
}

// Handle handles player disconnects.
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

		server.BroadcastMessage(utils.Yellow + player.GetDisplayName() + " has left the server")
	}
}
