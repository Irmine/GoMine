package p200

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/players/handlers"
	"github.com/irmine/goraklib/server"
)

type ResourcePackChunkRequestHandler struct {
	*handlers.PacketHandler
}

func NewResourcePackChunkRequestHandler() ResourcePackChunkRequestHandler {
	return ResourcePackChunkRequestHandler{handlers.NewPacketHandler()}
}

// Handle handles resource pack chunk requests, returning chunks of resource pack data to the client.
func (handler ResourcePackChunkRequestHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if request, ok := packet.(*p200.ResourcePackChunkRequestPacket); ok {
		if !server.GetPackManager().IsPackLoaded(request.PackUUID) {
			// TODO: Kick the player. We can't kick yet.
			return false
		}

		var pack = server.GetPackManager().GetPack(request.PackUUID)
		player.SendResourcePackChunkData(request.PackUUID, request.ChunkIndex, int64(data.ResourcePackChunkSize*request.ChunkIndex), pack.GetChunk(int(data.ResourcePackChunkSize*request.ChunkIndex), data.ResourcePackChunkSize))

		return true
	}

	return false
}
