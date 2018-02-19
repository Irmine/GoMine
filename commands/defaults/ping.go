package defaults

import (
	"strconv"

	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/utils"
)

func NewPing() *commands.Command {
	var ping = commands.NewCommand("ping", "Returns your latency", "gomine.ping", []string{}, func(sender commands.Sender) {
		if player, ok := sender.(interfaces.IPlayer); ok {
			player.SendMessage(utils.Yellow + "Your current latency/ping is: " + strconv.Itoa(int(player.GetPing())))
		}
	})
	ping.ExemptFromPermissionCheck(true)

	return ping
}
