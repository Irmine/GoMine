package arguments

import (
	"strconv"
	"gomine/interfaces"
)

type IntArg struct {
	*Argument
}

/**
 * Returns a new Int argument with the given name and optional value.
 */
func NewIntArg(name string, optional bool) *IntArg {
	return &IntArg{&Argument{name, optional, 1, nil}}
}

/**
 * Checks if the input value is able to be parsed as an Int.
 */
func (argument *IntArg) IsValidValue(value string, server interfaces.IServer) bool {
	return argument.IsInt(value)
}

/**
 * Converts the given value to a valid Int.
 */
func (argument *IntArg) ConvertValue(value string, server interfaces.IServer) interface{} {
	var int, _ = strconv.ParseInt(value, 10, 64)
	return int
}

/**
 * Returns if this argument should always merge.
 */
func (argument *IntArg) ShouldMerge() bool {
	return false
}
