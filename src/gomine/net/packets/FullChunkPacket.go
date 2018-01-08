package packets

import (
	"gomine/interfaces"
	"gomine/net/info"
)

type FullChunkPacket struct {
	*Packet
	Chunk interfaces.IChunk
}

func NewFullChunkPacket() *FullChunkPacket {
	return &FullChunkPacket{Packet: NewPacket(info.FullChunkDataPacket)}
}

func (pk *FullChunkPacket) Encode() {
	pk.PutVarInt(pk.Chunk.GetX())
	pk.PutVarInt(pk.Chunk.GetZ())
	var bytes = pk.Chunk.ToBinary()
	pk.PutUnsignedVarInt(uint32(len(bytes)))
	pk.PutBytes(bytes)
}

func (pk *FullChunkPacket) Decode() {

}