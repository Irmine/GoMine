package enchantments

// Type holds the data of the enchantment.
// It is an immutable type, which is used
// to identify an enchantment.
type Type struct {
	stringId string
	id       int16
}

// GetStringId returns the string ID of a type.
// This string ID may be used to identify
// enchantments by user output.
func (t Type) GetStringId() string {
	return t.stringId
}

// GetId returns the enchantment ID of a type.
// It is used mainly to identify an enchantment.
func (t Type) GetId() int16 {
	return t.id
}

// Instance is an enchantment instance.
// It holds an enchantment type,
// and contains the leftover duration of an
// enchantment, and the value of it.
type Instance struct {
	Type
	// Level is the enchantment level.
	// This value indicates the strength of the
	// enchantment.
	Level byte
}
