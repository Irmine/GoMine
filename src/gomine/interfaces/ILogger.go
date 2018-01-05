package interfaces

type ILogger interface {
	Log(string, string, ...interface{})
	Debug(...interface{})
	Info(...interface{})
	Notice(...interface{})
	Alert(...interface{})
	LogError(error)
	Error(...interface{})
	Warning(...interface{})
	Critical(...interface{})
	LogChat(...interface{})
	ProcessQueue(bool)
	Terminate()
	Sync()
}