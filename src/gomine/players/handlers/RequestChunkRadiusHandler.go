package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/utils"
	"gomine/net/packets/p200"
	"gomine/net/packets/data"
)

type RequestChunkRadiusHandler struct {
	*PacketHandler
}

func NewRequestChunkRadiusHandler() RequestChunkRadiusHandler {
	return RequestChunkRadiusHandler{NewPacketHandler()}
}

/**
 * Handles the chunk radius requests and initial spawns.
 */
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

			server.BroadcastMessage(utils.Yellow + player.GetName() + " has joined the server")
		}

		player.SendPlayStatus(data.StatusSpawn)

		return true
	}

	return false
}