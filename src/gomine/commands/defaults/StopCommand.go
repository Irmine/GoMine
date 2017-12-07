package defaults

import (
	"gomine/commands"
	"gomine/interfaces"
)

type StopCommand struct {
	*commands.Command
	server interfaces.IServer
}

func NewStop(server interfaces.IServer) StopCommand {
	return StopCommand{commands.NewCommand("stop", "Stops the server", "gomine.stop", []string{"shutdown"}), server}
}

func (command StopCommand) Execute() {
	command.server.Shutdown()
}