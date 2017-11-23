package arguments

import (
	"strconv"
	"gomine/interfaces"
)

type FloatArg struct {
	*Argument
}

/**
 * Returns a new Float argument with the given name and optional value.
 */
func NewFloatArg(name string, optional bool) *FloatArg {
	return &FloatArg{&Argument{name, optional, 1, nil}}
}

/**
 * Checks if the input value is able to be parsed as an Float.
 */
func (argument *FloatArg) IsValidValue(value string, server interfaces.IServer) bool {
	return argument.IsFloat(value)
}

/**
 * Converts the given value to a valid Float.
 */
func (argument *FloatArg) ConvertValue(value string, server interfaces.IServer) interface{} {
	var float, _ = strconv.ParseFloat(value, 64)
	return float
}

/**
 * Returns if this argument should always merge.
 */
func (argument *FloatArg) ShouldMerge() bool {
	return false
}