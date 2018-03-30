package items

import "github.com/irmine/gonbt"

// Stack is an instance of a given amount of items.
// A stack may also be referred to as an item instance.
// A stack holds additional information about an item,
// that could differ on an every item base.
type Stack struct {
	// Stack embeds Type. Therefore functions
	// in the Type struct may also be used in Stack.
	Type
	// Count is the current count of an item.
	// The count of an item is usually 16/64.
	Count byte
	// CustomName is the custom name of an item.
	// If a non-empty custom name has been set,
	// this name will be displayed,
	// rather than the original Type name.
	CustomName string
	// Durability is the current left durability of the stack.
	// Durability on non-breakable item types has no effect.
	Durability int16
}

func (stack *Stack) ParseNBT(compound *gonbt.Compound) {

}
