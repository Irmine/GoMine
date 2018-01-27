package gonbt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// A List contains an array with tags of the same type.
// It is prefixed with the list length, and contains only payloads of tags inside of it.
type List struct {
	*NamedTag
	tags []INamedTag
	tagType byte
}

func NewList(name string, tagType byte, tags []INamedTag) *List {
	return &List{NewNamedTag(name, TAG_List, nil), tags, tagType}
}

// Read reads all payloads into their right tags and puts them into the list.
func (list *List) Read(reader *NBTReader) {
	list.tagType = reader.GetByte()
	var length = reader.GetInt()

	for i := int32(0); i < length && !reader.Feof(); i++ {
		var tag = GetTagById(list.tagType, "")
		tag.Read(reader)
		list.tags = append(list.tags, tag)
	}
}


// Write writes the payload of all tags to the writer.
func (list *List) Write(writer *NBTWriter) {
	writer.PutByte(list.GetTagType())
	writer.PutInt(int32(len(list.tags)))

	for _, tag := range list.tags {
		tag.Write(writer)
	}
}


// GetTags returns all tags in this list.
func (list *List) GetTags() []INamedTag {
	return list.tags
}


// GetTagType returns the tag type of this list.
func (list *List) GetTagType() byte {
	return list.tagType
}


// GetTag returns a tag at the given offset in the list.
func (list *List) GetTag(offset int) INamedTag {
	return list.tags[offset]
}


// AddTag Adds a tag to the list.
// Returns an error if a tag with an invalid type was given.
func (list *List) AddTag(tag INamedTag) error {
	if tag.GetType() != list.GetTagType() {
		return errors.New("invalid tag for list")
	}
	list.tags = append(list.tags, tag)
	return nil
}


// Pop pushes the last tag off the list.
func (list *List) Pop() INamedTag {
	var tag = list.tags[len(list.tags) - 1]
	list.tags = list.tags[:len(list.tags) - 2]
	return tag
}


// Shift pushes the first tag off the list.
func (list *List) Shift() INamedTag {
	var tag = list.tags[0]
	list.tags = list.tags[1:]
	return tag
}


// DeleteAtOffset deletes a tag at the given offset and rearranges the list.
func (list *List) DeleteAtOffset(offset int) {
	if offset > len(list.tags) - 1 || offset < 0 {
		return
	}

	list.tags = append(list.tags[:offset], list.tags[offset + 1:]...)
}


// ToString converts the entire list to a readable string. Nesting level is used to indicate indentation.
func (list *List) toString(nestingLevel int) string {
	var str = strings.Repeat(" ", nestingLevel * 2)
	var entries = " entries"
	if len(list.tags) == 1 {
		entries = " entry"
	}

	str += "TAG_List('" + list.GetName() + " (" + GetTagName(list.tagType) + ")'): " + strconv.Itoa(len(list.tags)) + entries + "\n"
	str += strings.Repeat(" ", nestingLevel * 2) + "{\n"

	for _, tag := range list.tags {
		if list, ok := tag.(*List); ok {
			str += list.toString(nestingLevel + 1)
		} else {
			if compound, ok := tag.(*Compound); ok {
				str += compound.toString(nestingLevel + 1, true)
			} else {
				str += strings.Repeat(" ", (nestingLevel + 1) * 2)
				str += GetTagName(tag.GetType()) + "(None): " + fmt.Sprint(tag.Interface()) + "\n"
			}
		}
	}
	str += strings.Repeat(" ", nestingLevel * 2) + "}\n"
	return str
}