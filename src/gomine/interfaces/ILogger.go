package interfaces

type ILogger interface {
	Log(string, string, string)
	Debug(string)
	Info(string)
	Notice(string)
	Alert(string)
	Warning(string)
	Critical(string)
}
