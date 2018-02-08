package p200

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets/p200"
	"gomine/net/packets/data"
	"gomine/net/packets/types"
	"gomine/players/handlers"
)

type TextHandler struct {
	*handlers.PacketHandler
}

func NewTextHandler() TextHandler {
	return TextHandler{handlers.NewPacketHandler()}
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
			receiver.SendText(types.Text{Message: textPacket.Message, SourceName: textPacket.SourceName, SourceDisplayName: textPacket.SourceDisplayName, SourcePlatform: textPacket.SourcePlatform, SourceXUID: player.GetXUID(), TextType: data.TextChat})
		}

		server.GetLogger().LogChat("<" + player.GetDisplayName() + "> " + textPacket.Message)

		return true
	}
	return false
}
