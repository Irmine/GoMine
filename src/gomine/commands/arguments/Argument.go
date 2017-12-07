package arguments

import "strconv"

type Argument struct {
	name string
	optional bool
	inputArgs int
	output interface{}
}

/**
 * Returns the name of the argument.
 */
func (argument *Argument) GetName() string {
	return argument.name
}

/**
 * Sets the name of the argument.
 */
func (argument *Argument) SetName(name string) {
	argument.name = name
}

/**
 * Returns if the argument is optional.
 */
func (argument *Argument) IsOptional() bool {
	return argument.optional
}

/**
 * Sets the argument optional or non-optional.
 */
func (argument *Argument) SetOptional(value bool) {
	argument.optional = value
}

/**
 * Returns the amount of arguments of input this argument requires.
 */
func (argument *Argument) GetInputAmount() int {
	return argument.inputArgs
}

/**
 * Sets the amount of arguments the input of this argument requires.
 */
func (argument *Argument) SetInputAmount(amount int) {
	argument.inputArgs = amount
}

/**
 * Sets the output value of this argument.
 */
func (argument *Argument) SetOutput(value interface{}) {
	argument.output = value
}

/**
 * Returns the output value of this argument.
 */
func (argument *Argument) GetOutput() interface{} {
	return argument.output
}

/**
 * Checks if the input string is able to be parsed as an integer.
 */
func (argument *Argument) IsInt(value string) bool {
	var _, err = strconv.Atoi(value)
	return err == nil
}

/**
 * Checks if the input string is able to be parsed as an integer.
 */
func (argument *Argument) IsFloat(value string) bool {
	var _, err = strconv.ParseFloat(value, 64)
	return err == nil
}

/**
 * Returns if this argument should always merge.
 */
func (argument *Argument) ShouldMerge() bool {
	return false
}