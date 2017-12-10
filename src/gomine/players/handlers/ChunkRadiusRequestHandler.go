package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
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

		server.GetDefaultLevel().GetDefaultDimension().RequestChunks(player)

		var playStatus = packets.NewPlayStatusPacket()
		playStatus.Status = 3
		player.SendPacket(playStatus)
	}

	return true
}