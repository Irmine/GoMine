package p200

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type RequestChunkRadiusPacket struct {
	*packets.Packet
	Radius int32
}

func NewRequestChunkRadiusPacket() *RequestChunkRadiusPacket {
	return &RequestChunkRadiusPacket{packets.NewPacket(info.PacketIds200[info.RequestChunkRadiusPacket]), 0}
}

func (pk *RequestChunkRadiusPacket) Encode() {

}

func (pk *RequestChunkRadiusPacket) Decode() {
	pk.Radius = pk.GetVarInt()
}
