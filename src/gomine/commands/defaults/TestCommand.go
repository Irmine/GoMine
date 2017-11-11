package defaults

import (
	"gomine/commands"
	"gomine/interfaces"
	"gomine/commands/arguments"
	"fmt"
)

type TestCommand struct {
	*commands.Command
	server interfaces.IServer
}

func NewTest(server interfaces.IServer) TestCommand {
	var test = TestCommand{commands.NewCommand("test", "gomine.stop", []string{"test"}), server}

	var intArg = arguments.NewFloatArg("test", false)
	intArg.SetInputAmount(2)
	test.AppendArgument(intArg)

	var stringArg = arguments.NewStringArg("anotherTest", true)
	stringArg.SetInputAmount(2)
	test.AppendArgument(stringArg)
	return test
}

func (command TestCommand) Execute(arguments []interfaces.ICommandArgument) bool {
	fmt.Println(arguments[0].GetOutput())
	fmt.Println(arguments[1].GetOutput())
	return true
}

