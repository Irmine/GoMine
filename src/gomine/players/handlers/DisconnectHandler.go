package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/info"
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
	server.GetPlayerFactory().RemovePlayer(player)
	server.GetLogger().Debug(player.GetName(), "has left the server.")
}
