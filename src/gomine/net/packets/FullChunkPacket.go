package packets

import (
	"gomine/interfaces"
	"gomine/net/info"
)

type FullChunkPacket struct {
	*Packet
	ChunkX int32
	ChunkZ int32
	Chunk interfaces.IChunk
}

func NewFullChunkPacket() *FullChunkPacket {
	return &FullChunkPacket{Packet: NewPacket(info.FullChunkDataPacket)}
}

func (pk *FullChunkPacket) Encode() {
	pk.PutVarInt(pk.ChunkX)
	pk.PutVarInt(pk.ChunkZ)
	var bytes = pk.Chunk.ToBinary()
	pk.PutVarInt(int32(len(bytes)))
	pk.PutBytes(bytes)
}

func (pk *FullChunkPacket) Decode() {

}
