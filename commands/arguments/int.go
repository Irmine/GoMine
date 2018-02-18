package arguments

import (
	"strconv"

	"github.com/irmine/gomine/interfaces"
)

type IntArg struct {
	*Argument
}

// NewIntArg returns a new Int argument with the given name and optional value.
func NewIntArg(name string, optional bool) *IntArg {
	return &IntArg{&Argument{name, optional, 1, 0}}
}

func (argument *IntArg) IsValidValue(value string, server interfaces.IServer) bool {
	return argument.IsInt(value)
}

func (argument *IntArg) ConvertValue(value string, server interfaces.IServer) interface{} {
	var i, _ = strconv.ParseInt(value, 10, 64)
	return i
}
