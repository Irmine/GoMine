package p200

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/players/handlers"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/goraklib/server"
)

type RequestChunkRadiusHandler struct {
	*handlers.PacketHandler
}

func NewRequestChunkRadiusHandler() RequestChunkRadiusHandler {
	return RequestChunkRadiusHandler{handlers.NewPacketHandler()}
}

// Handle handles the chunk radius requests and initial spawn.
func (handler RequestChunkRadiusHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if chunkRadiusPacket, ok := packet.(*p200.RequestChunkRadiusPacket); ok {

		player.SetViewDistance(chunkRadiusPacket.Radius)

		player.SendChunkRadiusUpdated(player.GetViewDistance())

		var hasChunksInUse = player.HasAnyChunkInUse()

		server.GetDefaultLevel().GetDefaultDimension().RequestChunks(player, 10)

		if !hasChunksInUse {
			player.SetSpawned(true)

			var players = server.GetPlayerFactory().GetPlayers()
			for name, pl := range players {
				if !pl.HasSpawned() {
					delete(players, name)
				}
			}
			player.SendPlayerList(data.ListTypeAdd, players)

			for _, receiver := range server.GetPlayerFactory().GetPlayers() {
				if player != receiver {
					receiver.SendPlayerList(data.ListTypeAdd, map[string]interfaces.IPlayer{player.GetName(): player})

					receiver.SpawnTo(player)
					receiver.SpawnPlayerTo(player)
				}
			}

			player.SpawnToAll()
			player.SpawnPlayerToAll()

			player.UpdateAttributes()
			player.SendSetEntityData(player, player.GetEntityData())

			server.BroadcastMessage(utils.Yellow + player.GetDisplayName() + " has joined the server")
		}

		player.SendPlayStatus(data.StatusSpawn)

		return true
	}

	return false
}
