package interfaces

type ILogger interface {
	Log(string, string, string)
	Debug(string)
	Info(string)
	Notice(string)
	Alert(string)
	LogError(error)
	Error(string)
	Warning(string)
	Critical(string)
	ProcessQueue(bool)
	Terminate()
	Sync()
}