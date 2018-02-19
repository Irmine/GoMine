package defaults

import (
	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/interfaces"
)

func NewStop(server interfaces.IServer) *commands.Command {
	return commands.NewCommand("stop", "Stops the server", "gomine.stop", []string{"shutdown"}, func() {
		server.Shutdown()
	})
}
