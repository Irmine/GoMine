package arguments

import (
	"gomine/interfaces"
	"strings"
)

type StringEnum struct {
	*Argument
	options []string
}

/**
 * Returns a new String argument with the given name and optional value.
 */
func NewStringEnum(name string, optional bool, options []string) *StringEnum {
	return &StringEnum{&Argument{name, optional, 1, ""}, options}
}

/**
 * Checks if the input value is able to be parsed as a String.
 */
func (argument *StringEnum) IsValidValue(value string, server interfaces.IServer) bool {
	for _, option := range argument.options {
		if strings.ToLower(option) == strings.ToLower(value) {
			return true
		}
	}
	return false
}

/**
 * Converts the given value to a valid String.
 */
func (argument *StringEnum) ConvertValue(value string, server interfaces.IServer) interface{} {
	return strings.ToLower(value)
}

/**
 * Sets the output value of this argument.
 */
func (argument *StringEnum) SetOutput(value interface{}) {
	argument.output = value
}

/**
 * Returns if this argument should always merge.
 */
func (argument *StringEnum) ShouldMerge() bool {
	return true
}