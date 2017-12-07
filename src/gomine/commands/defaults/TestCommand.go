package defaults

import (
	"gomine/commands"
	"gomine/interfaces"
	"gomine/commands/arguments"
	"strconv"
)

type TestCommand struct {
	*commands.Command
	server interfaces.IServer
}

func NewTest(server interfaces.IServer) TestCommand {
	var test = TestCommand{commands.NewCommand("test", "Tests the command parser", "gomine.test", []string{"test2", "test3"}), server}
	test.ExemptFromPermissionCheck(true)

	test.AppendArgument(arguments.NewFloatArg("test", false))

	var stringArg = arguments.NewStringArg("anotherTest", true)
	stringArg.SetInputAmount(2)
	test.AppendArgument(stringArg)

	test.AppendArgument(arguments.NewStringEnum("testEnum", true, []string{"option", "test_option", "test"}))
	return test
}

func (command TestCommand) Execute(sender interfaces.ICommandSender, floatArg float64, stringArg string, enumString string) {
	sender.SendMessage("Float64: " + strconv.Itoa(int(floatArg)))
	sender.SendMessage("String: " + stringArg)
	sender.SendMessage("Enum String: " + enumString)
}