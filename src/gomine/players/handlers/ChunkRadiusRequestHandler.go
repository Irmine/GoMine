package handlers

import (
	"gomine/players"
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/worlds/chunks"
)

type ChunkRadiusRequestHandler struct {
	*players.PacketHandler
}

func NewChunkRadiusRequestHandler() ChunkRadiusRequestHandler {
	return ChunkRadiusRequestHandler{players.NewPacketHandler(info.RequestChunkRadiusPacket)}
}

/**
 * Handles the chunk radius requests.
 */
func (handler ChunkRadiusRequestHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if chunkRadiusPacket, ok := packet.(*packets.ChunkRadiusRequestPacket); ok {

		player.SetViewDistance(chunkRadiusPacket.Radius)

		var pk = packets.NewChunkRadiusUpdatedPacket()
		pk.Radius = player.GetViewDistance()
		server.GetRakLibAdapter().SendPacket(pk, session)

		for x := -1; x <= 1; x++ {
			for z := -1; z <= 1; z++ {
				var pk3 = packets.NewFullChunkPacket()
				pk3.ChunkX = int32(x)
				pk3.ChunkZ = int32(z)

				var chunk = chunks.NewChunk(256, 0, 0, make(map[int]interfaces.ISubChunk, 5), false, false, [256]byte{}, [4096]byte{})
				chunk.SetSubChunk(0, chunks.NewSubChunk())
				for x := 0; x < 16; x++ {
					for z := 0; z < 16; z++ {
						chunk.SetBlockId(x, 1, z, 1)
					}
				}

				pk3.Chunk = chunk
				server.GetRakLibAdapter().SendPacket(pk3, session)
			}
		}

		pk2 := packets.NewPlayStatusPacket()
		pk2.Status = 3
		server.GetRakLibAdapter().SendPacket(pk2, session)
	}

	return true
}
