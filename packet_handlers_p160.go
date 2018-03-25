package gomine

import (
	"github.com/irmine/gomine/net"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p160"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/text"
)

func NewTextHandler_160(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if textPacket, ok := packet.(*p160.TextPacket); ok {
			if textPacket.TextType != data.TextChat {
				return false
			}
			for _, receiver := range server.GetSessionManager().GetSessions() {
				receiver.SendText(types.Text{Message: textPacket.Message, SourceName: session.GetPlayer().GetName(), SourceDisplayName: session.GetPlayer().GetDisplayName(), SourceXUID: session.GetXUID(), TextType: data.TextChat})
			}
			text.DefaultLogger.LogChat("<" + session.GetPlayer().GetDisplayName() + "> " + textPacket.Message)
			return true
		}
		return false
	})
}
