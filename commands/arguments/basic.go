package arguments

import (
	"strconv"
	"strings"
)

// NewFloatArg returns a new Float argument with the given name and optional value.
func NewFloatArg(name string, optional bool) *Argument {
	return &Argument{name, optional, 1, float64(0), func(value string) bool {
		return IsFloat(value)
	}, func(value string) interface{} {
		var float, _ = strconv.ParseFloat(value, 64)
		return float
	}, false}
}

// NewIntArg returns a new Int argument with the given name and optional value.
func NewIntArg(name string, optional bool) *Argument {
	return &Argument{name, optional, 1, 0, func(value string) bool {
		return IsInt(value)
	}, func(value string) interface{} {
		var i, _ = strconv.ParseInt(value, 10, 64)
		return i
	}, false}
}

// NewStringArg returns a new String argument with the given name and optional value.
func NewStringArg(name string, optional bool) *Argument {
	var arg = &Argument{name, optional, 1, "", func(value string) bool {
		return true
	}, func(value string) interface{} {
		return value
	}, true}
	return arg
}

// NewStringEnum returns a new String argument with the given name and optional value.
func NewStringEnum(name string, optional bool, options []string) *Argument {
	var arg = &Argument{name, optional, 1, "", func(value string) bool {
		for _, option := range options {
			if strings.ToLower(option) == strings.ToLower(value) {
				return true
			}
		}
		return false
	}, func(value string) interface{} {
		return strings.ToLower(value)
	}, true}
	return arg
}
