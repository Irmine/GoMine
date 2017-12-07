package interfaces

import (
	"goraklib/server"
)

type IGoRakLibAdapter interface {
	GetSession(string, uint16) *server.Session
	SendBatch(IMinecraftPacketBatch, *server.Session)
	SendPacket(IPacket, *server.Session)
	Tick()
	GetRakLibServer() *server.GoRakLibServer
}