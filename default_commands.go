package gomine

import (
	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/net"
	"github.com/irmine/gomine/utils"
	"strconv"
)

func NewList(server *Server) *commands.Command {
	var list = commands.NewCommand("list", "Lists all players online", "gomine.list", []string{}, func(sender commands.Sender) {
		var s = "s"
		if len(server.GetSessionManager().GetSessions()) == 1 {
			s = ""
		}

		var playerList = utils.BrightGreen + "-----" + utils.White + " Player List (" + strconv.Itoa(len(server.GetSessionManager().GetSessions())) + " Player" + s + ") " + utils.BrightGreen + "-----\n"
		for name, player := range server.GetSessionManager().GetSessions() {
			playerList += utils.BrightGreen + name + ": " + utils.Yellow + utils.Bold + strconv.Itoa(int(player.GetPing())) + "ms" + utils.Reset + "\n"
		}
		sender.SendMessage(playerList)
	})
	list.ExemptFromPermissionCheck(true)
	return list
}

func NewPing() *commands.Command {
	var ping = commands.NewCommand("ping", "Returns your latency", "gomine.ping", []string{}, func(sender commands.Sender) {
		if session, ok := sender.(*net.MinecraftSession); ok {
			session.SendMessage(utils.Yellow+"Your current latency/ping is:", session.GetPing())
		} else {
			sender.SendMessage(utils.Red + "Please run this command as a player.")
		}
	})
	ping.ExemptFromPermissionCheck(true)
	return ping
}

func NewStop(server *Server) *commands.Command {
	return commands.NewCommand("stop", "Stops the server", "gomine.stop", []string{"shutdown"}, func() {
		server.Shutdown()
	})
}
