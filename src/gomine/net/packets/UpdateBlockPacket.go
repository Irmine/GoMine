package packets

import (
	"gomine/net/info"
	"gomine/vectors"
)

type UpdateBlockPacket struct{
	*Packet
	X, Z int32
	Y uint32
	BlockId, BlockMetadata, Flags uint32
}

func NewUpdateBlockPacket() *UpdateBlockPacket {
	return &UpdateBlockPacket{NewPacket(info.UpdateBlockPacket), 0, 0,0, 0, 0, 0}
}

func (pk *UpdateBlockPacket) Encode() {
	pk.PutBlockPos(vectors.TripleVector{float32(pk.X), float32(pk.Y), float32(pk.Z)})
	pk.PutUnsignedVarInt(pk.BlockId)
	pk.PutUnsignedVarInt((pk.Flags << 4) | pk.BlockMetadata)
}

func (pk *UpdateBlockPacket) Decode() {
	pk.X = pk.GetVarInt()
	pk.Y = pk.GetUnsignedVarInt()
	pk.Z = pk.GetVarInt()
	pk.BlockId = pk.GetUnsignedVarInt()
	v := pk.GetUnsignedVarInt()
	pk.BlockMetadata = v & 240
	pk.Flags = v >> 4
}