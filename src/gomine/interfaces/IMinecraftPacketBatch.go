package interfaces

type IMinecraftPacketBatch interface {
	GetPackets() []IPacket
	AddPacket(IPacket)
	Encode()
	Decode(IPlayer, ILogger)
	GetBuffer() []byte
}