package arguments

import (
	"gomine/interfaces"
	"strings"
)

type StringArg struct {
	*Argument
}

/**
 * Returns a new String argument with the given name and optional value.
 */
func NewStringArg(name string, optional bool) *StringArg {
	return &StringArg{&Argument{name, optional, 1, nil}}
}

/**
 * Checks if the input value is able to be parsed as a String.
 */
func (argument *StringArg) IsValidValue(value string, server interfaces.IServer) bool {
	return true
}

/**
 * Converts the given value to a valid String.
 */
func (argument *StringArg) ConvertValue(value string, server interfaces.IServer) interface{} {
	return value
}

/**
 * Sets the output value of this argument.
 */
func (argument *StringArg) SetOutput(value interface{}) {
	if slice, ok := value.([]string); ok {
		argument.output = strings.Join(slice, " ")
	} else {
		argument.output = value
	}
}