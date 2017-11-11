package interfaces

type ICommandSender interface {
	HasPermission(string) bool
	SendMessage(string)
}
