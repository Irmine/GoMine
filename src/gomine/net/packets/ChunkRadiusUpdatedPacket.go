package packets

import "gomine/net/info"

type ChunkRadiusUpdatedPacket struct {
	*Packet
	Radius uint32
}

func NewChunkRadiusUpdatedPacket() *ChunkRadiusUpdatedPacket {
	return &ChunkRadiusUpdatedPacket{NewPacket(info.ChunkRadiusUpdatedPacket), 0}
}

func (pk *ChunkRadiusUpdatedPacket) Encode()  {
	pk.PutUnsignedVarInt(pk.Radius)
}

func (pk *ChunkRadiusUpdatedPacket) Decode()  {

}
