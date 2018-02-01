package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type ChunkRadiusUpdatedPacket struct {
	*packets.Packet
	Radius int32
}

func NewChunkRadiusUpdatedPacket() *ChunkRadiusUpdatedPacket {
	return &ChunkRadiusUpdatedPacket{packets.NewPacket(info.ChunkRadiusUpdatedPacket), 0}
}

func (pk *ChunkRadiusUpdatedPacket) Encode()  {
	pk.PutVarInt(pk.Radius)
}

func (pk *ChunkRadiusUpdatedPacket) Decode()  {

}