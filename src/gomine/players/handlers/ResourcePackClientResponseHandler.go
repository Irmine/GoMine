package handlers

import (
	"gomine/players"
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
)

type ResourcePackClientResponseHandler struct {
	*players.PacketHandler
}

func NewResourcePackClientResponseHandler() ResourcePackClientResponseHandler {
	return ResourcePackClientResponseHandler{players.NewPacketHandler(info.ResourcePackClientResponsePacket)}
}

/**
 * Handles the resource pack client response.
 */
func (handler ResourcePackClientResponseHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if response, ok := packet.(*packets.ResourcePackClientResponsePacket); ok {
		if response.Status == 3 {
			var pk = packets.NewResourcePackStackPacket()
			server.GetRakLibAdapter().SendPacket(pk, session)
			return true
		}

		var pk4 = packets.NewStartGamePacket()
		server.GetRakLibAdapter().SendPacket(pk4, session)

		var pk = packets.NewPlayerListPacket()
		pk.Players = append(pk.Players, player)
		pk.ListType = packets.ListTypeAdd
		server.GetRakLibAdapter().SendPacket(pk, session)

		var pk3 = packets.NewCraftingDataPacket()
		server.GetRakLibAdapter().SendPacket(pk3, session)
	}

	return true
}
