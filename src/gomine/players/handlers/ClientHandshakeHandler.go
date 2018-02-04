package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets/p200"
	"gomine/net/packets/data"
)

type ClientHandshakeHandler struct {
	*PacketHandler
}

func NewClientHandshakeHandler() ClientHandshakeHandler {
	return ClientHandshakeHandler{NewPacketHandler()}
}

/**
 * Handles the client handshake, given to indicate that the client has enabled encryption.
 */
func (handler ClientHandshakeHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if _, ok := packet.(*p200.ClientHandshakePacket); ok {
		player.SendPlayStatus(data.StatusLoginSuccess)

		player.SendResourcePackInfo(server.GetConfiguration().ForceResourcePacks, server.GetPackHandler().GetResourceStack().GetPacks(), server.GetPackHandler().GetBehaviorStack().GetPacks())

		return true
	}

	return false
}
