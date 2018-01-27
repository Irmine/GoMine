package gonbt

import (
	"strconv"
	"strings"
)

// A compound has a map of name => tag, and is the base for any NBT structure.
// Compounds are suffixed with a TAG_End, which indicates the end of the tag.
type Compound struct {
	*NamedTag
	tags map[string]INamedTag
}

func NewCompound(name string, tags map[string]INamedTag) *Compound {
	if tags == nil {
		tags = make(map[string]INamedTag)
	}
	return &Compound{NewNamedTag(name, TAG_Compound, nil), tags}
}


// Read reads the data of the reader into the Compound.
func (compound *Compound) Read(reader *NBTReader) {
	for {
		var tag = reader.GetTag()
		if tag == nil || tag.GetType() == TAG_End {
			return
		}
		tag.Read(reader)

		compound.tags[tag.GetName()] = tag
	}
}


// Write writes all tags of the compound into the writer.
func (compound *Compound) Write(writer *NBTWriter) {
	for _, tag := range compound.tags {
		writer.PutTag(tag)
		tag.Write(writer)
	}
	writer.PutTag(NewEnd(""))
}


// HasTag checks if the compound has a tag with the given name.
func (compound *Compound) HasTag(name string) bool {
	var _, exists = compound.tags[name]
	return exists
}


// HasTagWithType checks if the compound has a tag with the given name and type.
func (compound *Compound) HasTagWithType(name string, tagType byte) bool {
	if !compound.HasTag(name) {
		return false
	}
	var tag = compound.GetTag(name)
	return tag.IsOfType(tagType)
}


// GetTag returns a tag with the given name.
func (compound *Compound) GetTag(name string) INamedTag {
	if !compound.HasTag(name) {
		return nil
	}
	return compound.tags[name]
}


// SetTag sets a tag in the compound.
func (compound *Compound) SetTag(tag INamedTag) {
	compound.tags[tag.GetName()] = tag
}


// GetTags returns all compound tags in a name => tag map.
func (compound *Compound) GetTags() map[string]INamedTag {
	return compound.tags
}


// SetByte sets a tag with the given name to the given byte.
func (compound *Compound) SetByte(name string, value byte) {
	compound.tags[name] = NewByte(name, value)
}


// GetByte returns a byte from the tag with the given name.
// If a byte tag with the name does not exist, it returns the default value.
func (compound *Compound) GetByte(name string, defaultValue byte) byte {
	if !compound.HasTagWithType(name, TAG_Byte) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(byte)
}


// SetShort sets a tag with the given name to the given int16.
func (compound *Compound) SetShort(name string, value int16) {
	compound.tags[name] = NewShort(name, value)
}


// GetShort returns a short from the tag with the given name.
// If a short tag with the name does not exist, it returns the default value.
func (compound *Compound) GetShort(name string, defaultValue int16) int16 {
	if !compound.HasTagWithType(name, TAG_Short) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int16)
}


// SetInt sets a tag with the given name to the given int32.
func (compound *Compound) SetInt(name string, value int32) {
	compound.tags[name] = NewInt(name, value)
}


// GetInt returns an int32 in an int tag with the given name.
func (compound *Compound) GetInt(name string, defaultValue int32) int32 {
	if !compound.HasTagWithType(name, TAG_Int) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int32)
}


// SetLong sets a tag with the given name to the given int64.
func (compound *Compound) SetLong(name string, value int64) {
	compound.tags[name] = NewLong(name, value)
}


// GetLong returns an int64 in a long tag with the given name.
func (compound *Compound) GetLong(name string, defaultValue int64) int64 {
	if !compound.HasTagWithType(name, TAG_Long) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int64)
}


// SetFloat sets a tag with the given name to the given float32.
func (compound *Compound) SetFloat(name string, value float32) {
	compound.tags[name] = NewFloat(name, value)
}


// GetFloat returns a float32 in a float tag with the given name.
func (compound *Compound) GetFloat(name string, defaultValue float32) float32 {
	if !compound.HasTagWithType(name, TAG_Float) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(float32)
}


// SetDouble sets a tag with the given name to the given float64.
func (compound *Compound) SetDouble(name string, value float64) {
	compound.tags[name] = NewDouble(name, value)
}


// GetDouble returns a float64 in a float tag with the given name.
func (compound *Compound) GetDouble(name string, defaultValue float64) float64 {
	if !compound.HasTagWithType(name, TAG_Double) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(float64)
}


// SetString sets a tag with the given name to the given string.
func (compound *Compound) SetString(name string, value string) {
	compound.tags[name] = NewString(name, value)
}


// GetString returns a string in a string tag with the given name.
func (compound *Compound) GetString(name string, defaultValue string) string {
	if !compound.HasTagWithType(name, TAG_String) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(string)
}


// SetByteArray sets a tag with the given name to the given byte array.
func (compound *Compound) SetByteArray(name string, value []byte) {
	compound.tags[name] = NewByteArray(name, value)
}


// GetByteArray returns a byte array in tag with the given name, or the default if none was found.
func (compound *Compound) GetByteArray(name string, defaultValue []byte) []byte {
	if !compound.HasTagWithType(name, TAG_Byte_Array) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().([]byte)
}


// SetIntArray sets a tag with the given name to the given int32 array.
func (compound *Compound) SetIntArray(name string, value []int32) {
	compound.tags[name] = NewIntArray(name, value)
}


// GetIntArray returns a int32 array in tag with the given name, or the default if none was found.
func (compound *Compound) GetIntArray(name string, defaultValue []int32) []int32 {
	if !compound.HasTagWithType(name, TAG_Int_Array) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().([]int32)
}


// SetLongArray sets a tag with the given name to the given int64 array.
func (compound *Compound) SetLongArray(name string, value []int64) {
	compound.tags[name] = NewLongArray(name, value)
}


// GetLongArray returns a int64 array in tag with the given name, or the default if none was found.
func (compound *Compound) GetLongArray(name string, defaultValue []int64) []int64 {
	if !compound.HasTagWithType(name, TAG_Long_Array) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().([]int64)
}


// SetList sets a list with the given name.
func (compound *Compound) SetList(name string, tagType byte, value []INamedTag) {
	compound.tags[name] = NewList(name, tagType, value)
}


// GetList returns a list with the given name and tag type.
// If a list with that name and/or tag type does not exist, returns nil.
func (compound *Compound) GetList(name string, tagType byte) *List {
	if !compound.HasTagWithType(name, TAG_List) {
		return nil
	}
	var list = compound.GetTag(name).(*List)
	if list.GetTagType() != tagType {
		return nil
	}
	return list
}


// SetCompound sets a compound with the given name.
func (compound *Compound) SetCompound(name string, value map[string]INamedTag) {
	compound.tags[name] = NewCompound(name, value)
}


// GetCompound returns a compound with the given name.
// If a compound with that name doesn't exist, returns nil.
func (compound *Compound) GetCompound(name string) *Compound {
	if !compound.HasTagWithType(name, TAG_Compound) {
		return nil
	}
	return compound.tags[name].(*Compound)
}


// Interface returns the compound as an interface.
func (compound *Compound) Interface() interface{} {
	return compound.tags
}


// toString converts the entire compound to a readable string. Nesting level is used to indicate indentation.
func (compound *Compound) toString(nestingLevel int, inList bool) string {
	var str = strings.Repeat(" ", nestingLevel * 2)
	var entries = " entries"

	if len(compound.tags) == 1 {
		entries = " entry"
	}

	var name = "'" + compound.GetName() + "'"
	if inList {
		name = "None"
	}

	str += "TAG_Compound(" + name + "): " + strconv.Itoa(len(compound.tags)) + entries + "\n"
	str += strings.Repeat(" ", nestingLevel * 2) + "{\n"

	for _, tag := range compound.tags {
		if list, ok := tag.(*List); ok {
			str += list.toString(nestingLevel + 1)
		} else {
			if compound, ok := tag.(*Compound); ok {
				str += compound.toString(nestingLevel + 1, false)
			} else {
				str += strings.Repeat(" ", (nestingLevel + 1) * 2)
				str += tag.ToString()
			}
		}
	}
	str += strings.Repeat(" ", nestingLevel * 2) + "}\n"
	return str
}


// ToString converts the entire compound to an uncompressed string.

func (compound *Compound) ToString() string {
	return compound.toString(0, false)
}
