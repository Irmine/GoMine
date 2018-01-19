package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
)

type ClientHandshakeHandler struct {
	*PacketHandler
}

func NewClientHandshakeHandler() ClientHandshakeHandler {
	return ClientHandshakeHandler{NewPacketHandler(info.ClientHandshakePacket)}
}

/**
 * Handles the client handshake, given to indicate that the client has enabled encryption.
 */
func (handler ClientHandshakeHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if _, ok := packet.(*packets.ClientHandshakePacket); ok {
		println("client handshake")

		playStatus := packets.NewPlayStatusPacket()
		playStatus.Status = 0
		player.SendPacket(playStatus)

		resourceInfo := packets.NewResourcePackInfoPacket()
		resourceInfo.MustAccept = server.GetConfiguration().ForceResourcePacks

		resourceInfo.ResourcePacks = server.GetPackHandler().GetResourceStack().GetPacks()
		resourceInfo.BehaviorPacks = server.GetPackHandler().GetBehaviorStack().GetPacks()

		player.SendPacket(resourceInfo)

		return true
	}

	return false
}
