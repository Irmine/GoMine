package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
)

type TextHandler struct {
	*PacketHandler
}

func NewTextHandler() TextHandler {
	return TextHandler{NewPacketHandler(info.TextPacket)}
}

func (handler TextHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if textPacket, ok := packet.(*packets.TextPacket); ok {
		for _, receiver := range server.GetPlayerFactory().GetPlayers() {
			pk := packets.NewTextPacket()
			pk.Message = textPacket.Message
			pk.TextType = textPacket.TextType
			pk.TextSource = textPacket.TextSource
			pk.XUID = textPacket.XUID

			receiver.SendPacket(pk)
		}

		return true
	}
	return false
}
