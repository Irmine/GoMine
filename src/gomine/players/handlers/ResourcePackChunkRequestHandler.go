package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets/p200"
	"gomine/net/packets/data"
)

type ResourcePackChunkRequestHandler struct {
	*PacketHandler
}

func NewResourcePackChunkRequestHandler() ResourcePackChunkRequestHandler {
	return ResourcePackChunkRequestHandler{NewPacketHandler()}
}

/**
 * Handles resource pack chunk requests, returning chunks of resource pack data to the client.
 */
func (handler ResourcePackChunkRequestHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if request, ok := packet.(*p200.ResourcePackChunkRequestPacket); ok {
		if !server.GetPackHandler().IsPackLoaded(request.PackUUID) {
			// TODO: Kick the player. We can't kick yet.
			return false
		}

		var pack = server.GetPackHandler().GetPack(request.PackUUID)
		player.SendResourcePackChunkData(request.PackUUID, request.ChunkIndex, int64(data.ResourcePackChunkSize * request.ChunkIndex), pack.GetChunk(int(data.ResourcePackChunkSize * request.ChunkIndex), data.ResourcePackChunkSize))

		return true
	}

	return false
}
