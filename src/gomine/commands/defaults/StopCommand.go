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
	var stop = StopCommand{commands.NewCommand("stop", "gomine.stop", []string{"shutdown"}), server}
	return stop
}

func (command StopCommand) Execute(sender interfaces.ICommandSender, arguments []interfaces.ICommandArgument) bool {
	command.server.Shutdown()
	return true
}
