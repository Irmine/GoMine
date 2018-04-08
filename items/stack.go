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
	Count int
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

// CanStackWith checks if two stacks can stack with each other.
// A bool is returned which indicates if the two can stack,
// and an integer is returned which specifies the count of
// of the item that can still be stacked on this stack.
// The returned integer may be 0, if the stack is already
// at the max stack size.
func (stack Stack) CanStackOn(stack2 *Stack) (bool, int) {
	if !stack.Type.Equals(stack2.Type) || stack.DisplayName != stack2.DisplayName || !stack.EqualsEnchantments(stack2) || !stack.EqualsLore(stack2) {
		return false, 0
	}
	count := stack2.maxStackSize - stack2.Count
	countLeft := stack.Count
	if countLeft < count {
		count = countLeft
	}
	return true, count
}

// StackOn attempts to stack a stack on another stack.
// A first bool is returned which indicates if the two stacked
// successfully. A second bool is returned which is true as long as
// the item stack is not at count 0.
// An integer is returned to specify the count of items that got stacked
// on the other stack. The integer returned may be 0, which happens if the
// other stack is already at max stack size.
func (stack *Stack) StackOn(stack2 *Stack) (success bool, notZero bool, stackCount int) {
	canStack, count := stack.CanStackOn(stack2)
	countLeft := stack.Count != 0
	if !canStack {
		return false, countLeft, count
	}
	stack2.Count += count
	stack.Count -= count
	return true, countLeft, count
}

// Equals checks if two item stacks are considered equal.
// Equals checks if the item type is equal and if the count is equal.
// For more deep checks, EqualsExact should be used.
func (stack Stack) Equals(stack2 *Stack) bool {
	return stack.Type.Equals(stack2.Type) && stack2.Count == stack.Count && stack.Durability == stack2.Durability && stack.DisplayName == stack2.DisplayName
}

// EqualsExact checks if two item stacks are considered exact equal.
// EqualsExact does all the checks Equals does,
// and checks if the lore and enchantments are equal.
func (stack Stack) EqualsExact(stack2 *Stack) bool {
	return stack.Equals(stack2) && stack.EqualsLore(stack2) && stack.EqualsEnchantments(stack2)
}

// EqualsLore checks if the lore of two item
// stacks are equal to each other.
func (stack Stack) EqualsLore(stack2 *Stack) bool {
	if len(stack.Lore) != len(stack2.Lore) {
		return false
	}
	for key, val := range stack.Lore {
		if stack2.Lore[key] != val {
			return false
		}
	}
	return true
}

// EqualsEnchantments checks if enchantments of two
// item stacks are equal to each other.
func (stack Stack) EqualsEnchantments(stack2 *Stack) bool {
	if len(stack.enchantments) != len(stack2.enchantments) {
		return false
	}
	for key, val := range stack.enchantments {
		if stack2.enchantments[key] != val {
			return false
		}
	}
	return true
}
