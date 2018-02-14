package arguments

import (
	"github.com/irmine/gomine/interfaces"
)

type StringArg struct {
	*Argument
}

/**
 * Returns a new String argument with the given name and optional value.
 */
func NewStringArg(name string, optional bool) *StringArg {
	return &StringArg{&Argument{name, optional, 1, ""}}
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
	argument.output = value
}

/**
 * Returns if this argument should always merge.
 */
func (argument *StringArg) ShouldMerge() bool {
	return true
}
