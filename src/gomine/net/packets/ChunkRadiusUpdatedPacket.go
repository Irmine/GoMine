package packets

import (
	"gomine/net/info"
)

type ChunkRadiusUpdatedPacket struct {
	*Packet
	Radius int32
}

func NewChunkRadiusUpdatedPacket() *ChunkRadiusUpdatedPacket {
	return &ChunkRadiusUpdatedPacket{NewPacket(info.ChunkRadiusUpdatedPacket), 0}
}

func (pk *ChunkRadiusUpdatedPacket) Encode()  {
	pk.PutVarInt(12/*pk.Radius*/)
}

func (pk *ChunkRadiusUpdatedPacket) Decode()  {

}
