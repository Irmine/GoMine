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

/**
 * Handles chatting of players.
 */
func (handler TextHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if textPacket, ok := packet.(*packets.TextPacket); ok {
		if textPacket.TextType != packets.TextChat {
			return false
		}

		for _, receiver := range server.GetPlayerFactory().GetPlayers() {
			pk := packets.NewTextPacket()
			pk.Message = textPacket.Message
			pk.TextType = textPacket.TextType
			pk.SourceName = textPacket.SourceName
			pk.SourceDisplayName = player.GetDisplayName()
			pk.SourcePlatform = textPacket.SourcePlatform
			pk.UnknownString = ""
			pk.XUID = player.GetXUID()

			receiver.SendPacket(pk)
		}

		server.GetLogger().LogChat("<" + player.GetDisplayName() + "> " + textPacket.Message)

		return true
	}
	return false
}
