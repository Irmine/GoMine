package arguments

import (
	"strings"

	"github.com/irmine/gomine/interfaces"
)

type StringEnum struct {
	*Argument
	options []string
}

// NewStringEnum returns a new String argument with the given name and optional value.
func NewStringEnum(name string, optional bool, options []string) *StringEnum {
	return &StringEnum{&Argument{name, optional, 1, ""}, options}
}

func (argument *StringEnum) IsValidValue(value string, server interfaces.IServer) bool {
	for _, option := range argument.options {
		if strings.ToLower(option) == strings.ToLower(value) {
			return true
		}
	}
	return false
}

func (argument *StringEnum) ConvertValue(value string, server interfaces.IServer) interface{} {
	return strings.ToLower(value)
}

// Returns if this argument should always merge.

func (argument *StringEnum) ShouldMerge() bool {
	return true
}
