package packets

import "gomine/net/info"

type ResourcePackChunkRequestPacket struct {
	*Packet
	PackUUID string
	ChunkIndex int32
}

func NewResourcePackChunkRequestPacket() *ResourcePackChunkRequestPacket {
	return &ResourcePackChunkRequestPacket{NewPacket(info.ResourcePackChunkRequestPacket), "", 0}
}

func (pk *ResourcePackChunkRequestPacket) Encode() {

}

func (pk *ResourcePackChunkRequestPacket) Decode() {
	pk.PackUUID = pk.GetString()
	pk.ChunkIndex = pk.GetLittleInt()
}
