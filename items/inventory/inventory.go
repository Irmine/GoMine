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

// FullInventory gets returned in AddItem when the
// inventory does not have enough space for the item.
var FullInventory = errors.New("inventory has no space for item")

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

// AddItem adds an item to the inventory.
// FullInventory gets returned if there was
// not sufficient space to fit the item.
// Items are first attempted to be stacked onto
// previously existed stacks, and once all
// pre-existing stacks are filled new stacks
// are created.
func (inventory *Inventory) AddItem(item *items.Stack) error {
	for slot, invItem := range inventory.items {
		if item.Count == 0 {
			return nil
		}
		if invItem == nil {
			continue
		}
		item.StackOn(invItem)
		inventory.SetItem(invItem, slot)
	}
	for slot, empty := range inventory.items {
		if item.Count == 0 {
			return nil
		}
		if empty != nil {
			continue
		}
		n := *item
		n.Count = 0
		item.StackOn(&n)
		inventory.SetItem(&n, slot)
	}
	if item.Count == 0 {
		return nil
	}
	return FullInventory
}

// RemoveItem removes an item from an inventory.
// A given item gets searched in the inventory,
// removing every equal stack until the count
// of the given stack has been exhausted.
// Items may be removed from multiple stacks.
// A bool gets returned to indicate if the
// complete stack got removed from the inventory.
func (inventory *Inventory) RemoveItem(searched *items.Stack) bool {
	count := searched.Count
	for slot, item := range inventory.items {
		if item == nil {
			continue
		}
		canStack, _ := item.CanStackOn(searched)
		if canStack {
			if item.Count > count {
				item.Count -= count
				inventory.SetItem(item, slot)
				count = 0
			} else {
				inventory.ClearSlot(slot)
			}
			count -= item.Count
			if count <= 0 {
				return true
			}
		}
	}
	return false
}

// ClearSlot clears a given slot in the inventory.
// ClearSlot returns ExceedingSlot if the slot exceeds
// the inventory size, and EmptySlot if the slot was
// already empty before clearing.
func (inventory *Inventory) ClearSlot(slot int) error {
	if slot >= len(inventory.items) {
		return ExceedingSlot
	}
	item := inventory.items[slot]
	if item == nil {
		return EmptySlot
	}
	inventory.SetItem(nil, slot)
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
func (inventory *Inventory) Contains(searched *items.Stack) bool {
	count := searched.Count
	for _, item := range inventory.items {
		if item == nil {
			continue
		}
		canStack, _ := searched.CanStackOn(item)
		if canStack {
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
			m[item.GetName()] += ", " + item.String()
		}
	}
	str := ""
	for _, instances := range m {
		str += instances + "\n"
	}
	return "Inventory contents:\n" + strings.TrimRight(str, "\n")
}
