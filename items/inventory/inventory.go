package inventory

import (
	"errors"
	"github.com/irmine/gomine/items"
	"strings"
)

// Inventory is a container of item stacks.
// Every inventory has a fixed amount of
// max slots, and the item stack count will
// never exceed these slots.
type Inventory struct {
	// items is a slice of item stacks.
	// The length of this slice will remain
	// fixed for the lifetime of an inventory.
	items []*items.Stack
}

// ExceedingSlot gets returned when an slot
// gets given that exceeds the inventory size.
// This may be for GetItem, or SetItem as example.
var ExceedingSlot = errors.New("slot given exceeds the inventory")

// EmptySlot gets returned in GetItem when a slot
// gets given and no item is available in that slot.
var EmptySlot = errors.New("slot given contains no item")

// NewInventory returns a new inventory with size.
// An item slice gets made with the size,
// which's length will never grow or shrink.
func NewInventory(size int) *Inventory {
	return &Inventory{make([]*items.Stack, size)}
}

// IsEmpty checks if a slot in the inventory is empty.
// True gets returned if no item was in the slot.
// True is also returned when the slot exceeds the
// maximum size of the inventory.
func (inventory *Inventory) IsEmpty(slot int) bool {
	if slot >= len(inventory.items) {
		return true
	}
	item := inventory.items[slot]
	return item == nil
}

// GetItem returns an item in a slot in an inventory.
// If the slot exceeds the max inventory size,
// a nil item gets returned with ExceedingSlot error.
// If there is no item available at that slot,
// a nil item gets returned with EmptySlot.
// If the item was retrieved successfully,
// the item gets returned with no error.
func (inventory *Inventory) GetItem(slot int) (*items.Stack, error) {
	if slot >= len(inventory.items) {
		return nil, ExceedingSlot
	}
	item := inventory.items[slot]
	if item == nil {
		return nil, EmptySlot
	}
	return item, nil
}

// SetItem sets an item in a slot in an inventory.
// If the slot exceeds the max inventory size,
// a nil item gets returned with ExceedingSlot error,
// otherwise returns nil.
func (inventory *Inventory) SetItem(stack *items.Stack, slot int) error {
	if slot >= len(inventory.items) {
		return ExceedingSlot
	}
	inventory.items[slot] = stack
	return nil
}

// GetAll returns a copied slice of all item stacks,
// that are currently contained within the inventory.
// Operating on this slice will not operate directly
// on the content of this inventory.
func (inventory *Inventory) GetAll() []*items.Stack {
	slice := make([]*items.Stack, len(inventory.items))
	copy(slice, inventory.items)
	return slice
}

// SetAll sets all items in the inventory.
// This function merely copies the items from
// slice to slice, and does not implement any
// other behaviour. Use SetItem where possible.
func (inventory *Inventory) SetAll(items []*items.Stack) {
	copy(inventory.items, items)
}

// Contains checks if the inventory contains an item.
// This function checks through the whole inventory,
// to try and find out the total count of items with
// the same type of the item stack.
// The checked item stack may therefore be split out
// over multiple stacks in the inventory.
// Contains only checks for the count of the item,
// and the right block type. Use ContainsExact for
// a deeper equality check.
func (inventory *Inventory) Contains(searched *items.Stack) bool {
	count := searched.Count
	for _, item := range inventory.items {
		if item == nil {
			continue
		}
		if item.Equals(*searched) {
			count -= item.Count
			if count <= 0 {
				return true
			}
		}
	}
	return false
}

// ContainsExact checks if the inventory contains an exact item.
// This function checks through the whole inventory,
// to try and find out the total count of items with
// the same type of the item stack.
// The checked item stack may therefore be split out
// over multiple stacks in the inventory.
// Exact items are checked against properties such as lore,
// enchantments, custom name and such.
func (inventory *Inventory) ContainsExact(searched *items.Stack) bool {
	count := searched.Count
	for _, item := range inventory.items {
		if item == nil {
			continue
		}
		if item.EqualsExact(*searched) {
			count -= item.Count
			if count <= 0 {
				return true
			}
		}
	}
	return false
}

// String returns a string representation of an inventory.
// String implements the fmt.Stringer interface.
func (inventory *Inventory) String() string {
	m := make(map[string]string)
	for _, item := range inventory.items {
		if item == nil {
			continue
		}
		if _, ok := m[item.GetName()]; !ok {
			m[item.GetName()] = "- " + item.String()
		} else {
			m[item.GetName()] += " " + item.String()
		}
	}
	str := ""
	for _, instances := range m {
		str += instances + "\n"
	}
	return strings.TrimRight(str, "\n")
}
