package entities

type EntityInterface interface {
	getId() int
}

type Entity struct {
	nameTag      string
	attributeMap attributeMap
}

func (entity *Entity) getAttributeMap() attributeMap {
	return entity.attributeMap
}

func (entity *Entity) getNameTag() string {
	return entity.nameTag
}

func (entity *Entity) setNameTag(name string) {
	entity.nameTag = name
}
