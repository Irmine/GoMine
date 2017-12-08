package entities

import (
	"gomine/vectors"
	"gomine/interfaces"
	"gomine/worlds/locations"
)

var RuntimeId uint64 = 0

type Entity struct {
	nameTag      string
	attributeMap *AttributeMap
	yaw, pitch float64
	position *locations.Position
	motion *vectors.TripleVector
	runtimeId uint64
	closed bool
	Health float32
}

func NewEntity(nameTag string, attributeMap *AttributeMap, yaw, pitch float64, position *locations.Position, motion *vectors.TripleVector, health float32) Entity {
	RuntimeId++
	return Entity{
		nameTag,
		attributeMap,
		yaw,
		pitch,
		position,
		motion,
		RuntimeId,
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
func (entity *Entity) GetPosition() *locations.Position {
	return entity.position
}

/**
 * Returns the motion of this entity.
 */
func (entity *Entity) GetMotion() *vectors.TripleVector {
	return entity.motion
}

/**
 * Returns the runtime ID of this entity.
 */
func (entity *Entity) GetRuntimeId() uint64 {
	return entity.runtimeId
}

/**
 * Returns the entity ID of this entity.
 */
func (entity *Entity) GetEntityId() int32 {
	return 0
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

	entity.position = &locations.Position{}
}

/**
 * Returns the health points of this entity.
 */
func (entity *Entity) GetHealth() float32 {
	return entity.Health
}

/**
 * Sets the health points of this entity.
 */
func (entity *Entity) SetHealth(health float32) {
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