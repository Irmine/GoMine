package interfaces

import (
	"goraklib/server"
	"github.com/Irmine/GoMine/src/gomine/interfaces"
)

type IGoRakLibAdapter interface {
	GetSession(string, uint16) *server.Session
	SendBatch(IMinecraftPacketBatch, *server.Session, byte)
	SendPacket(IPacket, *server.Session, byte)
	Tick()
	GetRakLibServer() *server.GoRakLibServer
	IsPacketRegistered(int) bool
	RegisterPacket(int, func() IPacket)
	GetPacket(int) IPacket
	RegisterPacketHandler(int, IPacketHandler, int) bool
	GetPacketHandlers(int) map[int][]IPacketHandler
}