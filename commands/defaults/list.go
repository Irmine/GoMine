package defaults

import (
	"strconv"

	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/utils"
)

func NewList(server interfaces.IServer) *commands.Command {
	var list = commands.NewCommand("list", "Lists all players online", "gomine.list", []string{}, func(sender commands.Sender) {
		var s = "s"
		if len(server.GetPlayerFactory().GetPlayers()) == 1 {
			s = ""
		}

		var playerList = utils.BrightGreen + "-----" + utils.White + " Player List (" + strconv.Itoa(len(server.GetPlayerFactory().GetPlayers())) + " Player" + s + ") " + utils.BrightGreen + "-----\n"
		for name, player := range server.GetPlayerFactory().GetPlayers() {
			playerList += utils.BrightGreen + name + ": " + utils.Yellow + utils.Bold + strconv.Itoa(int(player.GetPing())) + "ms" + utils.Reset + "\n"
		}
		sender.SendMessage(playerList)
	})
	list.ExemptFromPermissionCheck(true)
	return list
}
