package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/info"
	"gomine/net/packets"
)

type MoveHandler struct {
	*PacketHandler
}

func NewMoveHandler() MoveHandler {
	return MoveHandler{NewPacketHandler(info.MovePlayerPacket)}
}

func (handler MoveHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {

	if pk, ok := packet.(*packets.MovePlayerPacket); ok {
		player.Move(pk.Position.X, pk.Position.Y, pk.Position.Z, pk.Rotation.Pitch, pk.Rotation.Yaw, pk.Rotation.HeadYaw)
	}

	return true
}