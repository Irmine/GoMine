package p200

import (
	"strings"

	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/players/handlers"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/goraklib/server"
)

type CommandRequestHandler struct {
	*handlers.PacketHandler
}

func NewCommandRequestHandler() CommandRequestHandler {
	return CommandRequestHandler{handlers.NewPacketHandler()}
}

/**
 * Handles commands issues by players.
 */
func (handler CommandRequestHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {
	if pk, ok := packet.(*p200.CommandRequestPacket); ok {
		pk.CommandText = pk.CommandText[1:]
		var args = strings.Split(pk.CommandText, " ")

		var commandName = strings.TrimSpace(args[0])
		var holder = server.GetCommandHolder()

		if !holder.IsCommandRegistered(commandName) {
			player.SendMessage(utils.BrightRed + "Command could not be found")
			return false
		}

		var command, _ = holder.GetCommand(commandName)
		var parsedInput, valid = command.Parse(player, args[1:], server)

		if valid {
			commands.ParseIntoInputAndExecute(player, command, parsedInput)
		}

		return true
	}

	return false
}
