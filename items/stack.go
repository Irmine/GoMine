package items

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
	// AdditionalData is raw additional data of an item stack.
	// The AdditionalData should not be directly used by plugins,
	// but should rather be modified by encapsulating items.
	AdditionalData interface{}
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

/*
	TODO: Move getting aux value in item type, rather than stack.
// GetAuxValue returns the aux value for the item stack with item data.
// This aux value is used for writing stacks over network.
func (stack Stack) GetAuxValue(data int16) int32 {
	if stack.IsBreakable() {
		data = stack.Durability
	}
	return int32(((data & 0x7fff) << 8) | int16(stack.Count))
}*/
