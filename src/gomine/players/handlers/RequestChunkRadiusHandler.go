package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/utils"
)

type RequestChunkRadiusHandler struct {
	*PacketHandler
}

func NewRequestChunkRadiusHandler() RequestChunkRadiusHandler {
	return RequestChunkRadiusHandler{NewPacketHandler(info.RequestChunkRadiusPacket)}
}

/**
 * Handles the chunk radius requests and initial spawns.
 */
func (handler RequestChunkRadiusHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if chunkRadiusPacket, ok := packet.(*packets.RequestChunkRadiusPacket); ok {

		player.SetViewDistance(chunkRadiusPacket.Radius)

		var radiusUpdated = packets.NewChunkRadiusUpdatedPacket()
		radiusUpdated.Radius = player.GetViewDistance()
		player.SendPacket(radiusUpdated)

		var hasChunksInUse = player.HasAnyChunkInUse()

		server.GetDefaultLevel().GetDefaultDimension().RequestChunks(player, 10)

		if !hasChunksInUse {
			player.SetSpawned(true)

			var playerList = packets.NewPlayerListPacket()
			playerList.Players = server.GetPlayerFactory().GetPlayers()
			for name, pl := range playerList.Players {
				if !pl.HasSpawned() {
					delete(playerList.Players, name)
				}
			}
			playerList.ListType = packets.ListTypeAdd
			player.SendPacket(playerList)

			for _, receiver := range server.GetPlayerFactory().GetPlayers() {
				if player != receiver {
					list := packets.NewPlayerListPacket()
					list.ListType = packets.ListTypeAdd
					list.Players = map[string]interfaces.IPlayer{player.GetName(): player}
					receiver.SendPacket(list)

					receiver.SpawnTo(player)
					receiver.SpawnPlayerTo(player)
				}
			}

			player.SpawnToAll()
			player.SpawnPlayerToAll()

			player.UpdateAttributes()
			player.GetLevel().GetEntityHelper().SendEntityData(player.(interfaces.IEntity), player)

			server.BroadcastMessage(utils.Yellow + player.GetName() + " has joined the server")
		}

		var playStatus = packets.NewPlayStatusPacket()
		playStatus.Status = 3
		player.SendPacket(playStatus)

		return true
	}

	return false
}