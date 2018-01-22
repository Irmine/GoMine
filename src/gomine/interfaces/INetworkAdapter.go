package interfaces

import (
	"goraklib/server"
)

type INetworkAdapter interface {
	GetSession(string, uint16) *server.Session
	SendBatch(IMinecraftPacketBatch, *server.Session, byte)
	SendPacket(IPacket, IPlayer, byte)
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