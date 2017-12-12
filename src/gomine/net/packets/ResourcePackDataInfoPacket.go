package packets

import "gomine/net/info"

type ResourcePackDataInfoPacket struct {
	*Packet
	PackUUID string
	MaxChunkSize int32
	ChunkCount int32
	CompressedPackSize int64
	Sha256 string
}

func NewResourcePackDataInfoPacket() *ResourcePackDataInfoPacket {
	return &ResourcePackDataInfoPacket{NewPacket(info.ResourcePackDataInfoPacket), "", 0, 0, 0, ""}
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