package arguments

import "strconv"

type Argument struct {
	name      string
	optional  bool
	inputArgs int
	output    interface{}
}

// GetName returns the name of the argument.
func (argument *Argument) GetName() string {
	return argument.name
}

// SetName sets the name of the argument.
func (argument *Argument) SetName(name string) {
	argument.name = name
}

// IsOptional checks if the argument is optional.
func (argument *Argument) IsOptional() bool {
	return argument.optional
}

// SetOptional sets the argument optional or non-optional.
func (argument *Argument) SetOptional(value bool) {
	argument.optional = value
}

// GetInputAmount returns the amount of arguments of input this argument requires.
func (argument *Argument) GetInputAmount() int {
	return argument.inputArgs
}

// SetInputAmount sets the amount of arguments the input of this argument requires.
func (argument *Argument) SetInputAmount(amount int) {
	argument.inputArgs = amount
}

// SetOutput sets the output value of this argument.
func (argument *Argument) SetOutput(value interface{}) {
	argument.output = value
}

// GetOutput returns the output value of this argument.
func (argument *Argument) GetOutput() interface{} {
	return argument.output
}

// IsInt checks if the input string is able to be parsed as an integer.
func (argument *Argument) IsInt(value string) bool {
	var _, err = strconv.Atoi(value)
	return err == nil
}

// IsFloat checks if the input string is able to be parsed as an integer.
func (argument *Argument) IsFloat(value string) bool {
	var _, err = strconv.ParseFloat(value, 64)
	return err == nil
}

// ShouldMerge returns whether this argument should merge all its values or not.
func (argument *Argument) ShouldMerge() bool {
	return false
}
