package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/utils"
)

type ChunkRadiusRequestHandler struct {
	*PacketHandler
}

func NewChunkRadiusRequestHandler() ChunkRadiusRequestHandler {
	return ChunkRadiusRequestHandler{NewPacketHandler(info.RequestChunkRadiusPacket)}
}

/**
 * Handles the chunk radius requests.
 */
func (handler ChunkRadiusRequestHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if chunkRadiusPacket, ok := packet.(*packets.ChunkRadiusRequestPacket); ok {

		player.SetViewDistance(chunkRadiusPacket.Radius)

		var radiusUpdated = packets.NewChunkRadiusUpdatedPacket()
		radiusUpdated.Radius = player.GetViewDistance()
		player.SendPacket(radiusUpdated)

		var hasChunksInUse = player.HasAnyChunkInUse()

		server.GetDefaultLevel().GetDefaultDimension().RequestChunks(player, 10)

		if !hasChunksInUse {
			var playerList = packets.NewPlayerListPacket()
			playerList.Players = server.GetPlayerFactory().GetPlayers()
			playerList.ListType = packets.ListTypeAdd
			player.SendPacket(playerList)

			for _, receiver := range server.GetPlayerFactory().GetPlayers() {
				if player != receiver {
					var list = packets.NewPlayerListPacket()
					list.ListType = packets.ListTypeAdd
					list.Players = map[string]interfaces.IPlayer{player.GetName(): player}
					receiver.SendPacket(list)

					receiver.SpawnPlayerTo(player)
				}
			}

			player.SpawnPlayerToAll()

			server.BroadcastMessage(utils.Yellow + player.GetName() + " has joined the server")
		}

		var playStatus = packets.NewPlayStatusPacket()
		playStatus.Status = 3
		player.SendPacket(playStatus)

		return true
	}

	return false
}