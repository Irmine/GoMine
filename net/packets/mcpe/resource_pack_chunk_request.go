package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type ResourcePackChunkRequestPacket struct {
	*packets.Packet
	PackUUID   string
	ChunkIndex int32
}

func NewResourcePackChunkRequestPacket() *ResourcePackChunkRequestPacket {
	return &ResourcePackChunkRequestPacket{packets.NewPacket(info.PacketIds[info.ResourcePackChunkRequestPacket]), "", 0}
}

func (pk *ResourcePackChunkRequestPacket) Encode() {

}

func (pk *ResourcePackChunkRequestPacket) Decode() {
	pk.PackUUID = pk.GetString()
	pk.ChunkIndex = pk.GetLittleInt()
}
