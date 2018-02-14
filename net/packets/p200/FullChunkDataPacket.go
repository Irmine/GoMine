package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type FullChunkDataPacket struct {
	*packets.Packet
	ChunkX int32
	ChunkZ int32
	ChunkData []byte
}

func NewFullChunkDataPacket() *FullChunkDataPacket {
	return &FullChunkDataPacket{Packet: packets.NewPacket(info.PacketIds200[info.FullChunkDataPacket])}
}

func (pk *FullChunkDataPacket) Encode() {
	pk.PutVarInt(pk.ChunkX)
	pk.PutVarInt(pk.ChunkZ)
	pk.PutLengthPrefixedBytes(pk.ChunkData)
}

func (pk *FullChunkDataPacket) Decode() {

}