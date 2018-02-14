package defaults

import (
	"strconv"

	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/utils"
)

type PingCommand struct {
	*commands.Command
}

func NewPing() PingCommand {
	var ping = PingCommand{commands.NewCommand("ping", "Returns your latency", "gomine.ping", []string{})}
	ping.ExemptFromPermissionCheck(true)

	return ping
}

func (command PingCommand) Execute(sender interfaces.ICommandSender) {
	if player, ok := sender.(interfaces.IPlayer); ok {
		player.SendMessage(utils.Yellow + "Your current latency/ping is: " + strconv.Itoa(int(player.GetPing())))
	}
}
