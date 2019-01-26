package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type ChunkRadiusUpdatedPacket struct {
	*packets.Packet
	Radius int32
}

func NewChunkRadiusUpdatedPacket() *ChunkRadiusUpdatedPacket {
	return &ChunkRadiusUpdatedPacket{packets.NewPacket(info.PacketIds[info.ChunkRadiusUpdatedPacket]), 0}
}

func (pk *ChunkRadiusUpdatedPacket) Encode() {
	pk.PutVarInt(pk.Radius)
}

func (pk *ChunkRadiusUpdatedPacket) Decode() {

}
