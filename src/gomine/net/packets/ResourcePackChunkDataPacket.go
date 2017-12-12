package packets

import "gomine/net/info"

type ResourcePackChunkDataPacket struct {
	*Packet
	PackUUID string
	ChunkIndex int32
	Progress int64
	ChunkData []byte
}

func NewResourcePackChunkDataPacket() *ResourcePackChunkDataPacket {
	return &ResourcePackChunkDataPacket{NewPacket(info.ResourcePackChunkDataPacket), "", 0, 0, []byte{}}
}

func (pk *ResourcePackChunkDataPacket) Encode() {

}

func (pk *ResourcePackChunkDataPacket) Decode() {
	pk.PutString(pk.PackUUID)
	pk.PutLittleInt(pk.ChunkIndex)
	pk.PutLittleLong(pk.Progress)
	pk.PutLittleInt(int32(len(pk.ChunkData)))
	pk.PutBytes(pk.ChunkData)
}
