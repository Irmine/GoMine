package items

// Type is the type that identifies an item.
// Types contain an item ID and item data,
// which can be used to construct a new item stack.
type Type struct {
	name      string
	stringId  string
	id        int16
	data      int16
	breakable bool
}

// NewType returns a new non-breakable type.
// The given data gets used as the properties,
// and are immutable in the type.
func NewType(name string, stringId string, id int16, data int16) Type {
	return Type{name, stringId, id, data, false}
}

// NewBreakable returns a new breakable type.
// The given data gets used as the properties,
// and the item data is set to true.
func NewBreakable(name string, stringId string, id int16) Type {
	return Type{name, stringId, id, 0, true}
}

// GetName returns the readable name of an item type.
// This name may contains spaces.
func (t Type) GetName() string {
	return t.name
}

// GetStringId returns the string ID of an item type.
// StringIds are a string used as an identifier,
// in order to lookup items by it.
func (t Type) GetStringId() string {
	return t.stringId
}

// GetId returns the item ID of an item type.
// The ID is a signed int16. IDs may be negative.
func (t Type) GetId() int16 {
	return t.id
}

// GetData returns the data of an item type.
func (t Type) GetData() int16 {
	return t.data
}

// IsBreakable checks if an item is breakable.
// Breakable items use data fields for durability,
// but we separate them for forward compatibility sake.
func (t Type) IsBreakable() bool {
	return t.breakable
}
