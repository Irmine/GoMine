package entities

import (
	"gomine/vectors"
	"gomine/interfaces"
	"gomine/players/math"
)

var RuntimeId uint64 = 0

type Entity struct {
	attributeMap *AttributeMap
	Motion vectors.TripleVector
	runtimeId uint64
	closed bool
	Health float32

	Position vectors.TripleVector
	Level interfaces.ILevel
	Dimension interfaces.IDimension
	Rotation math.Rotation

}

func NewEntity(attributeMap *AttributeMap, motion vectors.TripleVector, health float32, position vectors.TripleVector, level interfaces.ILevel, rotation math.Rotation) Entity {
	RuntimeId++
	return Entity{
		attributeMap,
		motion,
		RuntimeId,
		false,
		health,
		position,
		level,
		level.GetDefaultDimension(),
		rotation,
	}
}

/**
 * Returns the attribute map of this entity.
 */
func (entity *Entity) GetAttributeMap() *AttributeMap {
	return entity.attributeMap
}

/**
 * Returns the current position of this entity.
 */
func (entity *Entity) GetPosition() vectors.TripleVector {
	return entity.Position
}

/**
 * Sets the position of this entity
 */
func (entity *Entity) SetPosition(v vectors.TripleVector)  {
	entity.Position = v
}

/**
 * Returns the level of this entity
 */
func (entity *Entity) GetLevel() interfaces.ILevel {
	return entity.Level
}

/**
 * Sets the level of this entity
 */
func (entity *Entity) SetLevel(v interfaces.ILevel)  {
	entity.Level = v
}

/**
 * Returns the level of this entity
 */
func (entity *Entity) GetDimension() interfaces.IDimension {
	return entity.Dimension
}

/**
 * Sets the level of this entity
 */
func (entity *Entity) SetDimension(v interfaces.IDimension)  {
	entity.Dimension = v
}

/**
 * Returns the current rotation of this entity.
 */
func (entity *Entity) GetRotation() math.Rotation {
	return entity.Rotation
}

/**
 * Sets the rotation of this entity
 */
func (entity *Entity) SetRotation(v math.Rotation)  {
	entity.Rotation = v
}

/**
 * Returns the motion of this entity.
 */
func (entity *Entity) GetMotion() vectors.TripleVector {
	return entity.Motion
}

/**
 * sets the motion of this entity
 */
func (entity *Entity) SetMotion(v vectors.TripleVector)  {
	entity.Motion = v
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
	entity.Position = *vectors.NewTripleVector(0, 0, 0)
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