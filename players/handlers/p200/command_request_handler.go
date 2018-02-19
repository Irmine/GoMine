package p200

import (
	"strings"

	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/players/handlers"
	"github.com/irmine/goraklib/server"
)

type CommandRequestHandler struct {
	*handlers.PacketHandler
}

func NewCommandRequestHandler() CommandRequestHandler {
	return CommandRequestHandler{handlers.NewPacketHandler()}
}

// Handle handles commands issues by players.
func (handler CommandRequestHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if pk, ok := packet.(*p200.CommandRequestPacket); ok {
		var args = strings.Split(pk.CommandText, " ")
		var commandName = args[0]
		var i = 1
		for !server.GetCommandManager().IsCommandRegistered(commandName) {
			commandName += " " + args[i]
			if i == len(args)-1 {
				break
			}
		}

		var manager = server.GetCommandManager()

		if !manager.IsCommandRegistered(commandName) {
			server.GetLogger().Error("Command could not be found.")
			return false
		}
		args = args[i:]

		var command, _ = manager.GetCommand(commandName)
		command.Execute(server, args)

		return true
	}

	return false
}
