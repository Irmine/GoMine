package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets/p200"
)

type MovePlayerHandler struct {
	*PacketHandler
}

func NewMovePlayerHandler() MovePlayerHandler {
	return MovePlayerHandler{NewPacketHandler()}
}

/**
 * Handles the synchronization of player movement server sided.
 */
func (handler MovePlayerHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if pk, ok := packet.(*p200.MovePlayerPacket); ok {
		if !player.HasSpawned() {
			return false
		}

		player.SyncMove(pk.Position.X, pk.Position.Y, pk.Position.Z, pk.Rotation.Pitch, pk.Rotation.Yaw, pk.Rotation.HeadYaw, pk.OnGround)
		player.GetDimension().RequestChunks(player, player.GetViewDistance())

		for _, player2 := range player.GetViewers() {
			player2.SendPacket(packet)
		}

		return true
	}

	return false
}