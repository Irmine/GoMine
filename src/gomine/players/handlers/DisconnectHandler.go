package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/info"
	"gomine/utils"
	"gomine/net/packets"
)

type DisconnectHandler struct {
	*PacketHandler
}

func NewDisconnectHandler() DisconnectHandler {
	return DisconnectHandler{NewPacketHandler(info.DisconnectPacket)}
}

/**
 * The disconnect handler is a special case. It does not follow the rules of the other handlers.
 */
func (handler DisconnectHandler) Handle(player interfaces.IPlayer, session *server.Session, server interfaces.IServer) {
	if player.GetSession() == nil {
		return
	}

	server.GetPlayerFactory().RemovePlayer(player)

	for _, online := range server.GetPlayerFactory().GetPlayers() {
		if online != player {
			var list = packets.NewPlayerListPacket()
			list.ListType = packets.ListTypeRemove
			list.Players = map[string]interfaces.IPlayer{player.GetName(): player}
			online.SendPacket(list)
		}
	}

	server.BroadcastMessage(utils.Yellow + player.GetName() + " has left the server")
}
