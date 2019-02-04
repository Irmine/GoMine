package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type RequestChunkRadiusPacket struct {
	*packets.Packet
	Radius int32
}

func NewRequestChunkRadiusPacket() *RequestChunkRadiusPacket {
	return &RequestChunkRadiusPacket{packets.NewPacket(info.PacketIds[info.RequestChunkRadiusPacket]), 0}
}

func (pk *RequestChunkRadiusPacket) Encode() {

}

func (pk *RequestChunkRadiusPacket) Decode() {
	pk.Radius = pk.GetVarInt()
}
