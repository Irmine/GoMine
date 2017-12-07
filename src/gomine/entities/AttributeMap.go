package entities

type AttributeMap struct {
	attributes map[int]Attribute
}

func NewAttributeMap() *AttributeMap {
	return &AttributeMap{}
}

func (attMap *AttributeMap) getAttributes() map[int]Attribute {
	return attMap.attributes
}