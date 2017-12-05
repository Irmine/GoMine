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
	var test = TestCommand{commands.NewCommand("test", "Tests the command parser", "gomine.stop", []string{"test"}), server}
	test.ExemptFromPermissionCheck(true)

	test.AppendArgument(arguments.NewFloatArg("test", false, 0))

	var stringArg = arguments.NewStringArg("anotherTest", true, "test")
	stringArg.SetInputAmount(2)
	test.AppendArgument(stringArg)

	test.AppendArgument(arguments.NewStringEnum("testEnum", true, []string{"option", "test_option", "test"}, "test"))
	return test
}

func (command TestCommand) Execute(sender interfaces.ICommandSender, floatArg float64, stringArg string, enumString string) {
	fmt.Println("Float64:", floatArg)
	fmt.Println("String:", stringArg)
	fmt.Println("Enum String:", enumString)
}