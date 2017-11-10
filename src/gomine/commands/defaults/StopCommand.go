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
	var stop = StopCommand{commands.NewCommand(StopCommand{}, "stop", "gomine.stop", []string{"kill", "shutdown"}), server}
	return stop
}

func (command StopCommand) Execute(commandText string) bool {
	command.server.Shutdown()
	return true
}
