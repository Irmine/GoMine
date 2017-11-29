package entities

import (
	"gomine/vectorMath"
	"gomine/interfaces"
)

var runtimeId uint64 = 0

type EntityInterface interface {
	getId() int
}

type Entity struct {
	nameTag      string
	attributeMap *AttributeMap
	yaw, pitch float64
	position, motion *vectorMath.TripleVector
	runtimeId uint64
	closed bool
	Health int
}

func NewEntity(nameTag string, attributeMap *AttributeMap, yaw, pitch float64, position, motion *vectorMath.TripleVector, health int) Entity {
	runtimeId++
	return Entity{
		nameTag,
		attributeMap,
		yaw,
		pitch,
		position,
		motion,
		runtimeId,
		false,
		health,
	}
}

/**
 * Returns the attribute map of this entity.
 */
func (entity *Entity) GetAttributeMap() *AttributeMap {
	return entity.attributeMap
}

/**
 * Returns the name tag of this entity.
 */
func (entity *Entity) GetNameTag() string {
	return entity.nameTag
}

/**
 * Sets the name tag of this entity.
 */
func (entity *Entity) SetNameTag(name string) {
	entity.nameTag = name
}

/**
 * Returns the current position of this entity.
 */
func (entity *Entity) GetPosition() *vectorMath.TripleVector {
	return entity.position
}

/**
 * Returns the motion of this entity.
 */
func (entity *Entity) GetMotion() *vectorMath.TripleVector {
	return entity.motion
}

/**
 * Returns the runtime ID of this entity.
 */
func (entity *Entity) GetRuntimeId() uint64 {
	return entity.runtimeId
}

/**
 * Checks if the entity is closed and not to be used anymore.
 */
func (entity *Entity) IsClosed() bool {
	return entity.closed
}

/**
 * Closes the entity making it unable to be used.
 */
func (entity *Entity) Close() {
	entity.closed = true
	//todo
}

/**
 * Returns the health points of this entity.
 */
func (entity *Entity) GetHealth() int {
	return entity.Health
}

/**
 * Sets the health points of this entity.
 */
func (entity *Entity) SetHealth(health int) {
	entity.Health = health
}

/**
 * Kills the current entity.
 */
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