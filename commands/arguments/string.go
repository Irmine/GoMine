package arguments

import (
	"github.com/irmine/gomine/interfaces"
)

type StringArg struct {
	*Argument
}

// NewStringArg returns a new String argument with the given name and optional value.
func NewStringArg(name string, optional bool) *StringArg {
	return &StringArg{&Argument{name, optional, 1, ""}}
}

func (argument *StringArg) IsValidValue(value string, server interfaces.IServer) bool {
	return true
}

func (argument *StringArg) ConvertValue(value string, server interfaces.IServer) interface{} {
	return value
}

func (argument *StringArg) ShouldMerge() bool {
	return true
}
