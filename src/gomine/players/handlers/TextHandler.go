package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets/p200"
	"gomine/net/packets/data"
	"gomine/net/packets/types"
)

type TextHandler struct {
	*PacketHandler
}

func NewTextHandler() TextHandler {
	return TextHandler{NewPacketHandler()}
}

/**
 * Handles chatting of players.
 */
func (handler TextHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if textPacket, ok := packet.(*p200.TextPacket); ok {
		if textPacket.TextType != data.TextChat {
			return false
		}

		for _, receiver := range server.GetPlayerFactory().GetPlayers() {
			var text = types.Text{}
			text.Message = textPacket.Message
			text.TextType = data.TextChat
			text.SourceName = textPacket.SourceName
			text.SourceDisplayName = player.GetDisplayName()
			text.SourcePlatform = player.GetPlatform()
			text.SourceXUID = player.GetXUID()

			receiver.SendText(text)
		}

		server.GetLogger().LogChat("<" + player.GetDisplayName() + "> " + textPacket.Message)

		return true
	}
	return false
}
