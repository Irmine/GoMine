package interfaces

import (
	"goraklib/server"
)

type IPacketHandler interface {
	GetId() int
	Handle(IPacket, IPlayer, *server.Session, IServer) bool
}