package arguments

import (
	"strconv"

	"github.com/irmine/gomine/interfaces"
)

type FloatArg struct {
	*Argument
}

// NewFloatArg returns a new Float argument with the given name and optional value.
func NewFloatArg(name string, optional bool) *FloatArg {
	return &FloatArg{&Argument{name, optional, 1, float64(0)}}
}

func (argument *FloatArg) IsValidValue(value string, server interfaces.IServer) bool {
	return argument.IsFloat(value)
}

func (argument *FloatArg) ConvertValue(value string, server interfaces.IServer) interface{} {
	var float, _ = strconv.ParseFloat(value, 64)
	return float
}
