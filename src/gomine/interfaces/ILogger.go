package interfaces

type ILogger interface {
	Log(string, string)
	Debug(string)
	Info(string)
	Alert(string)
	Warning(string)
	Critical(string)
}
