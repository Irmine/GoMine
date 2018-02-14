package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type ResourcePackChunkDataPacket struct {
	*packets.Packet
	PackUUID string
	ChunkIndex int32
	Progress int64
	ChunkData []byte
}

func NewResourcePackChunkDataPacket() *ResourcePackChunkDataPacket {
	return &ResourcePackChunkDataPacket{packets.NewPacket(info.PacketIds200[info.ResourcePackChunkDataPacket]), "", 0, 0, []byte{}}
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
