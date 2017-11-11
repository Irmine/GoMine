package interfaces

type ICommandArgument interface {
	GetName() string
	SetName(string)
	IsOptional() bool
	SetOptional(bool)
	GetInputAmount() int
	SetOutput(interface{})
	GetOutput() interface{}
	IsValidValue(string, IServer) bool
	ConvertValue(string, IServer) interface{}
}
