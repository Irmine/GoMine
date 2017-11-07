package entities

type attributeMap struct {
	attributes map[int]attribute
}

func (attMap *attributeMap) getAttributes() map[int]attribute {
	return attMap.attributes
}
