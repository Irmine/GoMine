package interfaces

type IMinecraftPacketBatch interface {
	GetPackets() []IPacket
	AddPacket(IPacket)
	Encode()
	Decode(ILogger)
	GetBuffer() []byte
}