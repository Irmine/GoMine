package interfaces

import (
	"goraklib/server"
	"gomine/utils"
)

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