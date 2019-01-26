package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/worlds/blocks"
)

type UpdateBlockPacket struct {
	*packets.Packet
	Position                      blocks.Position
	BlockId, BlockMetadata, Flags uint32
}

func NewUpdateBlockPacket() *UpdateBlockPacket {
	return &UpdateBlockPacket{Packet: packets.NewPacket(info.PacketIds[info.UpdateBlockPacket])}
}

func (pk *UpdateBlockPacket) Encode() {
	pk.PutBlockPosition(pk.Position)
	pk.PutUnsignedVarInt(pk.BlockId)
	pk.PutUnsignedVarInt((pk.Flags << 4) | pk.BlockMetadata)
}

func (pk *UpdateBlockPacket) Decode() {
	pk.Position = pk.GetBlockPosition()
	pk.BlockId = pk.GetUnsignedVarInt()
	v := pk.GetUnsignedVarInt()
	pk.BlockMetadata = v & 240
	pk.Flags = v >> 4
}
