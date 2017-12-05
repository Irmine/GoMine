package interfaces

import (
	"gomine/utils"
)

type IMinecraftPacketBatch interface {
	GetPackets() []IPacket
	AddPacket(IPacket)
	Encode()
	Decode(ILogger)
	GetStream() *utils.BinaryStream
}