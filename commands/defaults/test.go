package defaults

import (
	"strconv"

	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/commands/arguments"
)

func NewTest() *commands.Command {
	var test = commands.NewCommand("test space command", "Tests the command parser", "gomine.test", []string{"test2", "test3"}, func(sender commands.Sender, floatArg float64, stringArg string, enumString string) {
		sender.SendMessage("Float64: " + strconv.Itoa(int(floatArg)))
		sender.SendMessage("String: " + stringArg)
		sender.SendMessage("Enum String: " + enumString)
	})
	test.ExemptFromPermissionCheck(true)

	test.AppendArgument(arguments.NewFloat("test", false))

	var stringArg = arguments.NewString("anotherTest", true)
	stringArg.SetInputAmount(2)
	test.AppendArgument(stringArg)

	test.AppendArgument(arguments.NewStringEnum("testEnum", true, []string{"option", "test_option", "test"}))
	return test
}
