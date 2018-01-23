package interfaces

import "goraklib/server"

type IPacketHandler interface {
	GetId() int
	Handle(IPacket, IPlayer, *server.Session, IServer) bool
	SetPriority(int) bool
	GetPriority() int
}

type IPacket interface {
	SetBuffer([]byte)
	GetBuffer() []byte
	GetId() int
	EncodeHeader()
	Encode()
	DecodeHeader()
	Decode()
	ResetStream()
	GetOffset() int
	SetOffset(int)
	Discard()
	IsDiscarded() bool
}

type INetworkAdapter interface {
	GetSession(string, uint16) *server.Session
	SendBatch(IMinecraftPacketBatch, *server.Session, byte)
	SendPacket(IPacket, IMinecraftSession, byte)
	Tick()
	GetRakLibServer() *server.GoRakLibServer
	IsPacketRegistered(int) bool
	RegisterPacket(int, func() IPacket)
	GetPacket(int) IPacket
	RegisterPacketHandler(int, IPacketHandler, int) bool
	GetPacketHandlers(int) [][]IPacketHandler
	DeregisterPacketHandlers(int, int)
	DeletePacket(id int)
}

type IMinecraftPacketBatch interface {
	GetPackets() []IPacket
	AddPacket(IPacket)
	Encode()
	Decode()
	GetBuffer() []byte
}

type IMinecraftSession interface {
	GetServer() IServer
	GetSession() *server.Session
	GetPing() uint64
	GetUUID() utils.UUID
	GetXUID() string
	SetLanguage(string)
	GetLanguage() string
	GetClientId() int
	GetEncryptionHandler() *utils.EncryptionHandler
	UsesEncryption() bool
	EnableEncryption()
	IsXBOXLiveAuthenticated() bool
	SetXBOXLiveAuthenticated(bool)
	SendPacket(IPacket)
	SendBatch(IMinecraftPacketBatch)
	IsInitialized() bool
}