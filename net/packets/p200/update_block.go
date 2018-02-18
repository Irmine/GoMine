package p200

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/golang/geo/r3"
)

type UpdateBlockPacket struct {
	*packets.Packet
	X, Z                          int32
	Y                             uint32
	BlockId, BlockMetadata, Flags uint32
}

func NewUpdateBlockPacket() *UpdateBlockPacket {
	return &UpdateBlockPacket{packets.NewPacket(info.PacketIds200[info.UpdateBlockPacket]), 0, 0, 0, 0, 0, 0}
}

func (pk *UpdateBlockPacket) Encode() {
	pk.PutBlockPos(r3.Vector{float64(pk.X), float64(pk.Y), float64(pk.Z)})
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
