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
	if _, ok := packet.(*packets.ChunkRadiusRequestPacket); ok {

		player.SetViewDistance(2/*chunkRadiusPacket.Radius*/)

		var pk = packets.NewChunkRadiusUpdatedPacket()
		pk.Radius = player.GetViewDistance()
		server.GetRakLibAdapter().SendPacket(pk, session)

		server.GetDefaultLevel().GetDefaultDimension().RequestChunks(player)

		pk2 := packets.NewPlayStatusPacket()
		pk2.Status = 3
		server.GetRakLibAdapter().SendPacket(pk2, session)
	}

	return true
}