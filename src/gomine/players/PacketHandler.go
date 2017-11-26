package players

import (
	"gomine/interfaces"
	"goraklib/server"
)

type IPacketHandler interface {
	GetId() int
	Handle(interfaces.IPacket, Player, *server.Session, interfaces.IServer) bool
}

type PacketHandler struct {
	id int
}

func NewPacketHandler(id int) *PacketHandler {
	return &PacketHandler{id}
}

func (handler *PacketHandler) GetId() int {
	return handler.id
}
