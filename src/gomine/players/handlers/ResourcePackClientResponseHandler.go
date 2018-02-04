package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/vectors"
	"gomine/entities/math"
	"gomine/net/packets/p200"
	"gomine/net/packets/data"
)

type ResourcePackClientResponseHandler struct {
	*PacketHandler
}

func NewResourcePackClientResponseHandler() ResourcePackClientResponseHandler {
	return ResourcePackClientResponseHandler{NewPacketHandler()}
}

/**
 * Handles the resource pack client response.
 */
func (handler ResourcePackClientResponseHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if response, ok := packet.(*p200.ResourcePackClientResponsePacket); ok {
		switch response.Status {
		case data.StatusRefused:
			// TODO: Kick the player. We can't kick yet.
			return false

		case data.StatusSendPacks:
			for _, packUUID := range response.PackUUIDs {
				if !server.GetPackHandler().IsPackLoaded(packUUID) {
					// TODO: Kick the player. We can't kick yet.
					return false
				}
				var pack = server.GetPackHandler().GetPack(packUUID)

				player.SendResourcePackDataInfo(pack)
			}

		case data.StatusHaveAllPacks:
			player.SendResourcePackStack(server.GetConfiguration().ForceResourcePacks, server.GetPackHandler().GetResourceStack().GetPacks(), server.GetPackHandler().GetBehaviorStack().GetPacks())

		case data.StatusCompleted:
			player.PlaceInWorld(vectors.NewTripleVector(0, 20, 0), math.NewRotation(0, 0, 0), server.GetDefaultLevel(), server.GetDefaultLevel().GetDefaultDimension())
			player.SetFinalized()

			player.SendStartGame(player)

			player.SendCraftingData()
		}
		return true
	}

	return false
}