package p200

import (
	"gomine/interfaces"
	"gomine/net/info"
	"gomine/net/packets"
)

type FullChunkDataPacket struct {
	*packets.Packet
	Chunk interfaces.IChunk
}

func NewFullChunkDataPacket() *FullChunkDataPacket {
	return &FullChunkDataPacket{Packet: packets.NewPacket(info.FullChunkDataPacket)}
}

func (pk *FullChunkDataPacket) Encode() {
	pk.PutVarInt(pk.Chunk.GetX())
	pk.PutVarInt(pk.Chunk.GetZ())
	var bytes = pk.Chunk.ToBinary()
	pk.PutUnsignedVarInt(uint32(len(bytes)))
	pk.PutBytes(bytes)
}

func (pk *FullChunkDataPacket) Decode() {

}