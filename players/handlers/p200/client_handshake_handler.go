package p200

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/players/handlers"
	"github.com/irmine/goraklib/server"
)

type ClientHandshakeHandler struct {
	*handlers.PacketHandler
}

func NewClientHandshakeHandler() ClientHandshakeHandler {
	return ClientHandshakeHandler{handlers.NewPacketHandler()}
}

// Handle handles the client handshake, given to indicate that the client has enabled encryption.
func (handler ClientHandshakeHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if _, ok := packet.(*p200.ClientHandshakePacket); ok {
		player.SendPlayStatus(data.StatusLoginSuccess)

		player.SendResourcePackInfo(server.GetConfiguration().ForceResourcePacks, server.GetPackManager().GetResourceStack().GetPacks(), server.GetPackManager().GetBehaviorStack().GetPacks())

		return true
	}

	return false
}
