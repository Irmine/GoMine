package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type ResourcePackChunkDataPacket struct {
	*packets.Packet
	PackUUID   string
	ChunkIndex int32
	Progress   int64
	ChunkData  []byte
}

func NewResourcePackChunkDataPacket() *ResourcePackChunkDataPacket {
	return &ResourcePackChunkDataPacket{packets.NewPacket(info.PacketIds[info.ResourcePackChunkDataPacket]), "", 0, 0, []byte{}}
}

func (pk *ResourcePackChunkDataPacket) Encode() {
	pk.PutString(pk.PackUUID)
	pk.PutLittleInt(pk.ChunkIndex)
	pk.PutLittleLong(pk.Progress)
	pk.PutLittleInt(int32(len(pk.ChunkData)))
	pk.PutBytes(pk.ChunkData)
}

func (pk *ResourcePackChunkDataPacket) Decode() {

}
