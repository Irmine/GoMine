package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type ResourcePackChunkRequestPacket struct {
	*packets.Packet
	PackUUID string
	ChunkIndex int32
}

func NewResourcePackChunkRequestPacket() *ResourcePackChunkRequestPacket {
	return &ResourcePackChunkRequestPacket{packets.NewPacket(info.ResourcePackChunkRequestPacket), "", 0}
}

func (pk *ResourcePackChunkRequestPacket) Encode() {

}

func (pk *ResourcePackChunkRequestPacket) Decode() {
	pk.PackUUID = pk.GetString()
	pk.ChunkIndex = pk.GetLittleInt()
}
