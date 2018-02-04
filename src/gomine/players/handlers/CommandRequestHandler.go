package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"strings"
	"gomine/commands"
	"gomine/utils"
	"gomine/net/packets/p200"
)

type CommandRequestHandler struct {
	*PacketHandler
}

func NewCommandRequestHandler() CommandRequestHandler {
	return CommandRequestHandler{NewPacketHandler()}
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