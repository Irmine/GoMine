package interfaces

type ILogger interface {
	Log(message string, logLevel string)
	Debug(message string)
	Info(message string)
	Alert(message string)
	Warning(message string)
	Critical(message string)
}
