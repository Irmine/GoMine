package interfaces

import (
	"goraklib/server"
)

type IGoRakLibAdapter interface {
	SendBatch(IMinecraftPacketBatch, *server.Session)
	SendPacket(IPacket, *server.Session)
	Tick()
}
