package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/worlds/blocks"
)

const (
	DataLayerNormal = iota
	DataLayerLiquid
)

type UpdateBlockPacket struct {
	*packets.Packet
	Position    blocks.Position
	BlockRuntimeId   uint32
	Flags       uint32
	DataLayerId uint32
}

func NewUpdateBlockPacket() *UpdateBlockPacket {
	return &UpdateBlockPacket{Packet: packets.NewPacket(info.PacketIds[info.UpdateBlockPacket]), Flags: 0x02, DataLayerId: DataLayerNormal}
}

func (pk *UpdateBlockPacket) Encode() {
	pk.PutBlockPosition(pk.Position)
	pk.PutUnsignedVarInt(pk.BlockRuntimeId)
	pk.PutUnsignedVarInt(pk.Flags)
	pk.PutUnsignedVarInt(pk.DataLayerId)
}

func (pk *UpdateBlockPacket) Decode() {
	pk.Position = pk.GetBlockPosition()
	pk.BlockRuntimeId = pk.GetUnsignedVarInt()
	pk.Flags = pk.GetUnsignedVarInt()
	pk.DataLayerId = pk.GetUnsignedVarInt()
}
