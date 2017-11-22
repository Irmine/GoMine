package entities

import (
	"gomine/vectorMath"
	"gomine/interfaces"
)

var EId uint64 = 0

type EntityInterface interface {
	getId() int
}

type Entity struct {
	nameTag      string
	attributeMap AttributeMap
	yaw, pitch float64
	position, motion vectorMath.TripleVector
	eId uint64
	closed bool
	Health int
}

func NewEntity(nameTag string, attributeMap AttributeMap, yaw float64, pitch float64, position vectorMath.TripleVector, motion vectorMath.TripleVector, health int) Entity {
	EId++
	return Entity{
		nameTag,
		attributeMap,
		yaw,
		pitch,
		position,
		motion,
		EId,
		false,
		health,
	}
}

func (entity *Entity) GetAttributeMap() AttributeMap {
	return entity.attributeMap
}

func (entity *Entity) GetNameTag() string {
	return entity.nameTag
}

func (entity *Entity) SetNameTag(name string) {
	entity.nameTag = name
}

func (entity *Entity) GetPosition() vectorMath.TripleVector {
	return entity.position
}

func (entity *Entity) GetMotion() vectorMath.TripleVector {
	return entity.motion
}

func (entity *Entity) GetId() uint64 {
	return entity.eId
}

func (entity *Entity) IsClosed() bool {
	return entity.closed
}

func (entity *Entity) Close() {
	entity.closed = true
	//todo
}

func (entity *Entity) GetHealth() int {
	return entity.Health
}

func (entity *Entity) SetHealth(health int) {
	entity.Health = health
}

func (entity *Entity) Kill() {
	entity.Health = 0
	//todo
}

func (entity *Entity) SpawnTo(player interfaces.IPlayer)  {
	//todo
}

func (entity *Entity) SpawnToAll()  {
	//todo
}

func (entity *Entity) SpawnPacket(player interfaces.IPlayer)  {
	//todo
}

func (entity *Entity) Tick()  {

}