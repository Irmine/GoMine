package interfaces

import (
	"goraklib/server"
)

type IGoRakLibAdapter interface {
	GetSession(string, uint16) *server.Session
	SendBatch(IMinecraftPacketBatch, *server.Session, byte)
	SendPacket(IPacket, *server.Session, byte)
	Tick()
	GetRakLibServer() *server.GoRakLibServer
}