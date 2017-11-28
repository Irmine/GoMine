package packets

import (
	"gomine/net/info"
)

type ChunkRadiusRequestPacket struct {
	*Packet
	Radius uint32
}

func NewChunkRadiusRequestPacket() *ChunkRadiusRequestPacket {
	return &ChunkRadiusRequestPacket{NewPacket(info.RequestChunkRadiusPacket), 0}
}

func (pk *ChunkRadiusRequestPacket) Encode()  {

}

func (pk *ChunkRadiusRequestPacket) Decode()  {
	pk.Radius = pk.GetUnsignedVarInt()
}
