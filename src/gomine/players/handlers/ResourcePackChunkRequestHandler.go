package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
)

type ResourcePackChunkRequestHandler struct {
	*PacketHandler
}

func NewResourcePackChunkRequestHandler() ResourcePackChunkRequestHandler {
	return ResourcePackChunkRequestHandler{NewPacketHandler(info.ResourcePackChunkRequestPacket)}
}

func (handler ResourcePackChunkRequestHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if request, ok := packet.(*packets.ResourcePackChunkRequestPacket); ok {
		if !server.GetPackHandler().IsPackLoaded(request.PackUUID) {
			// TODO: Kick the player. We can't kick yet.
			return false
		}

		var pack = server.GetPackHandler().GetPack(request.PackUUID)
		var packData = packets.NewResourcePackChunkDataPacket()
		packData.PackUUID = request.PackUUID
		packData.ChunkIndex = request.ChunkIndex
		packData.ChunkData = pack.GetChunk(int(ChunkSize * request.ChunkIndex), ChunkSize)
		packData.Progress = int64(ChunkSize * request.ChunkIndex)

		player.SendPacket(packData)
	}

	return true
}
