package commands

type Sender interface {
	HasPermission(string) bool
	SendMessage(...interface{})
}
