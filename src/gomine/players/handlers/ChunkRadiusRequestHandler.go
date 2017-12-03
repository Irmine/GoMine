package handlers

import (
	"gomine/players"
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/worlds/chunks"
	"fmt"
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

		fmt.Println("Setting view distance:", chunkRadiusPacket.Radius)
		player.SetViewDistance(chunkRadiusPacket.Radius)

		var pk = packets.NewChunkRadiusUpdatedPacket()
		pk.Radius = player.GetViewDistance()
		server.GetRakLibAdapter().SendPacket(pk, session)

		var pk3 = packets.NewFullChunkPacket()
		pk3.ChunkX = 0
		pk3.ChunkZ = 0

		var chunk = chunks.NewChunk(256, 0, 0, make(map[int]interfaces.ISubChunk, 5), false, false, [256]byte{}, [4096]byte{})
		chunk.SetSubChunk(0, chunks.NewSubChunk())
		for x := 0; x < 16; x++ {
			for z := 0; z < 16; z++ {
				chunk.SetBlockId(x, 0, z, 1)
			}
		}

		pk3.Chunk = chunk
		server.GetRakLibAdapter().SendPacket(pk3, session)
	}

	return true
}
