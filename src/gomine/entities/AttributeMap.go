package entities

type AttributeMap struct {
	attributes map[int]Attribute
}

func (attMap *AttributeMap) getAttributes() map[int]Attribute {
	return attMap.attributes
}
