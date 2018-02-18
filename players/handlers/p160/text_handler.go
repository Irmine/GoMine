package p160

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p160"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/players/handlers"
	"github.com/irmine/goraklib/server"
)

type TextHandler struct {
	*handlers.PacketHandler
}

func NewTextHandler() TextHandler {
	return TextHandler{handlers.NewPacketHandler()}
}

// Handle handles chatting of players.
func (handler TextHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if textPacket, ok := packet.(*p160.TextPacket); ok {
		if textPacket.TextType != data.TextChat {
			return false
		}

		for _, receiver := range server.GetPlayerFactory().GetPlayers() {
			receiver.SendText(types.Text{Message: textPacket.Message, SourceName: player.GetDisplayName(), SourceDisplayName: textPacket.SourceName, SourceXUID: player.GetXUID(), TextType: data.TextChat})
		}

		server.GetLogger().LogChat("<" + player.GetDisplayName() + "> " + textPacket.Message)

		return true
	}
	return false
}
