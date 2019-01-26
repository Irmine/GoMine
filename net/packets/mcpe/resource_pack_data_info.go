package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type ResourcePackDataInfoPacket struct {
	*packets.Packet
	PackUUID           string
	MaxChunkSize       int32
	ChunkCount         int32
	CompressedPackSize int64
	Sha256             string
}

func NewResourcePackDataInfoPacket() *ResourcePackDataInfoPacket {
	return &ResourcePackDataInfoPacket{packets.NewPacket(info.PacketIds[info.ResourcePackDataInfoPacket]), "", 0, 0, 0, ""}
}

func (pk *ResourcePackDataInfoPacket) Encode() {
	pk.PutString(pk.PackUUID)
	pk.PutLittleInt(pk.MaxChunkSize)
	pk.PutLittleInt(pk.ChunkCount)
	pk.PutLittleLong(pk.CompressedPackSize)
	pk.PutString(pk.Sha256)
}

func (pk *ResourcePackDataInfoPacket) Decode() {

}
