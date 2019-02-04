package gomine

import (
	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/net"
	"github.com/irmine/gomine/text"
	"strconv"
)

func NewTest(server *Server) *commands.Command {
	cmd := commands.NewCommand("chunk", "Lists the current chunk", "none", []string{}, func(sender commands.Sender) {
		if session, ok := sender.(*net.MinecraftSession); ok {
			text.DefaultLogger.Debug(session.GetPlayer().GetChunk().X, session.GetPlayer().GetChunk().Z)
			session.SendMessage(session.GetPlayer().GetChunk().X, session.GetPlayer().GetChunk().Z)
		}
	})
	cmd.ExemptFromPermissionCheck(true)
	return cmd
}

func NewList(server *Server) *commands.Command {
	var list = commands.NewCommand("list", "Lists all players online", "gomine.list", []string{}, func(sender commands.Sender) {
		var s = "s"
		if len(server.SessionManager.GetSessions()) == 1 {
			s = ""
		}

		var playerList = text.BrightGreen + "-----" + text.White + " Player List (" + strconv.Itoa(len(server.SessionManager.GetSessions())) + " Player" + s + ") " + text.BrightGreen + "-----\n"
		for name, player := range server.SessionManager.GetSessions() {
			playerList += text.BrightGreen + name + ": " + text.Yellow + text.Bold + strconv.Itoa(int(player.GetPing())) + "ms" + text.Reset + "\n"
		}
		sender.SendMessage(playerList)
	})
	list.ExemptFromPermissionCheck(true)
	return list
}

func NewPing() *commands.Command {
	var ping = commands.NewCommand("ping", "Returns your latency", "gomine.ping", []string{}, func(sender commands.Sender) {
		if session, ok := sender.(*net.MinecraftSession); ok {
			session.SendMessage(text.Yellow+"Your current latency/ping is:", session.GetPing())
		} else {
			sender.SendMessage(text.Red + "Please run this command as a player.")
		}
	})
	ping.ExemptFromPermissionCheck(true)
	return ping
}

func NewStop(server *Server) *commands.Command {
	return commands.NewCommand("stop", "Stops the server", "gomine.stop", []string{"shutdown"}, func() {
		for _, session := range server.SessionManager.GetSessions() {
			session.Kick("Server Stopped", false, true)
		}

		server.Shutdown()
	})
}
