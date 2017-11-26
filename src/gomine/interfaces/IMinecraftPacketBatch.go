package interfaces

import (
	"gomine/utils"
)

type IMinecraftPacketBatch interface {
	GetPackets() []IPacket
	AddPacket(IPacket)
	Encode()
	Decode()
	GetStream() *utils.BinaryStream
}