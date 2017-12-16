package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/info"
	"gomine/net/packets"
)

type MovePlayerHandler struct {
	*PacketHandler
}

func NewMovePlayerHandler() MovePlayerHandler {
	return MovePlayerHandler{NewPacketHandler(info.MovePlayerPacket)}
}

func (handler MovePlayerHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {

	if pk, ok := packet.(*packets.MovePlayerPacket); ok {
		player.Move(pk.Position.X, pk.Position.Y, pk.Position.Z, pk.Rotation.Pitch, pk.Rotation.Yaw, pk.Rotation.HeadYaw)
	}

	return true
}