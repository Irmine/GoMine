package gonbt

// These are all tag types in NBT.
const (
	TAG_End byte = iota
	TAG_Byte
	TAG_Short
	TAG_Int
	TAG_Long
	TAG_Float
	TAG_Double
	TAG_Byte_Array
	TAG_String
	TAG_List
	TAG_Compound
	TAG_Int_Array
	TAG_Long_Array
)

// A map of tag type => tag name to convert types to names.
var tagNames = map[byte]string{
	TAG_End: "TAG_End",
	TAG_Byte: "TAG_Byte",
	TAG_Short: "TAG_Short",
	TAG_Int: "TAG_Int",
	TAG_Long: "TAG_Long",
	TAG_Float: "TAG_Float",
	TAG_Double: "TAG_Double",
	TAG_Byte_Array: "TAG_Byte_Array",
	TAG_String: "TAG_String",
	TAG_List: "TAG_List",
	TAG_Compound: "TAG_Compound",
	TAG_Int_Array: "TAG_Int_Array",
	TAG_Long_Array: "TAG_Long_Array",
}


// GetTagById gets a tag by the given ID with the given name.
func GetTagById(tagId byte, name string) INamedTag {
	switch tagId {
	case TAG_End:
		return NewEnd(name)
	case TAG_Byte:
		return NewByte(name, 0)
	case TAG_Short:
		return NewShort(name, 0)
	case TAG_Int:
		return NewInt(name, 0)
	case TAG_Long:
		return NewLong(name, 0)
	case TAG_Float:
		return NewFloat(name, 0)
	case TAG_Double:
		return NewDouble(name, 0)
	case TAG_Byte_Array:
		return NewByteArray(name, []byte{})
	case TAG_String:
		return NewString(name, "")
	case TAG_List:
		return NewList(name, TAG_Byte, []INamedTag{})
	case TAG_Compound:
		return NewCompound(name, map[string]INamedTag{})
	case TAG_Int_Array:
		return NewIntArray(name, []int32{})
	case TAG_Long_Array:
		return NewLongArray(name, []int64{})
	}
	return nil
}


// GetTagName returns the tag name associated with the given ID.
func GetTagName(tagId byte) string {
	return tagNames[tagId]
}