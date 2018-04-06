package items

import (
	"fmt"
	"github.com/irmine/gomine/items/enchantments"
	"github.com/irmine/gonbt"
)

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
	// Durability is the current left durability of the stack.
	// Durability on non-breakable item types has no effect.
	Durability int16
	// DisplayName is the display name of an item.
	// If a non-empty display name has been set,
	// this name will be displayed,
	// rather than the original Type name.
	DisplayName string
	// Lore is the displayed lore of an item.
	// The lore is displayed under the item,
	// when hovering over it in the inventory.
	Lore []string
	// enchantments is a map of enchantment instances,
	// that are applied on this item.
	// The map is indexed by the enchantment IDs.
	enchantments map[string]enchantments.Instance
	// additionalData is raw additional data of an item stack.
	// The additionalData may not be directly used by plugins,
	// but should rather be modified by encapsulating items.
	additionalData interface{}
	// cachedNBT is an NBT compound which gets set when parsing NBT.
	// This cached NBT is used to ensure no NBT gets lost while parsing,
	// and forms the base for NBT that gets emitted by the type.
	cachedNBT *gonbt.Compound
}

// GetDisplayName returns the displayed name of an item.
// The custom name of the item always gets returned,
// unless the custom name is empty; Then the actual
// item type name gets returned.
func (stack Stack) GetDisplayName() string {
	if stack.DisplayName == "" {
		return stack.name
	}
	return stack.DisplayName
}

// String returns a string representation of a stack.
// It implements fmt.Stringer, and returns a string as such:
// x29 Emerald (minecraft:emerald)
func (stack Stack) String() string {
	return fmt.Sprint("x", stack.Count, stack.Type)
}

// Equals checks if two item stacks are considered equal.
// Equals checks if the item type is equal and if the count is equal.
// For more deep checks, EqualsExact should be used.
func (stack Stack) Equals(stack2 Stack) bool {
	return stack.Type.Equals(stack2.Type)
}

// EqualsExact checks if two item stacks are considered exact equal.
// EqualsExact does all the checks Equals does,
// and checks if the lore and enchantments are equal,
// as well as the custom name and durability.
func (stack Stack) EqualsExact(stack2 Stack) bool {
	if len(stack.Lore) != len(stack2.Lore) {
		return false
	}
	for key, val := range stack.Lore {
		if stack2.Lore[key] != val {
			return false
		}
	}
	if len(stack.enchantments) != len(stack2.enchantments) {
		return false
	}
	for key, val := range stack.enchantments {
		if stack2.enchantments[key] != val {
			return false
		}
	}
	return stack.Equals(stack2) && stack.DisplayName == stack2.DisplayName && stack.Durability == stack2.Durability
}
