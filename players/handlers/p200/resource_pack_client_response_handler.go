package p200

import (
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/players/handlers"
	"github.com/irmine/gomine/vectors"
	"github.com/irmine/goraklib/server"
)

type ResourcePackClientResponseHandler struct {
	*handlers.PacketHandler
}

func NewResourcePackClientResponseHandler() ResourcePackClientResponseHandler {
	return ResourcePackClientResponseHandler{handlers.NewPacketHandler()}
}

// Handle handles the resource pack client response.
func (handler ResourcePackClientResponseHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if response, ok := packet.(*p200.ResourcePackClientResponsePacket); ok {
		switch response.Status {
		case data.StatusRefused:
			// TODO: Kick the player. We can't kick yet.
			return false

		case data.StatusSendPacks:
			for _, packUUID := range response.PackUUIDs {
				if !server.GetPackManager().IsPackLoaded(packUUID) {
					// TODO: Kick the player. We can't kick yet.
					return false
				}
				var pack = server.GetPackManager().GetPack(packUUID)

				player.SendResourcePackDataInfo(pack)
			}

		case data.StatusHaveAllPacks:
			player.SendResourcePackStack(server.GetConfiguration().ForceResourcePacks, server.GetPackManager().GetResourceStack().GetPacks(), server.GetPackManager().GetBehaviorStack().GetPacks())

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
