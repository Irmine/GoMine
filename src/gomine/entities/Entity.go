package entities

import (
	"gomine/vectors"
	"gomine/interfaces"
	"gomine/entities/math"
	math2 "math"
	"sync"
	"gomine/entities/data"
)

var RuntimeId uint64 = 0

const (
	DataFlags = 0
)

const (
	AffectedByGravity = 46
)

type Entity struct {
	attributeMap *data.AttributeMap
	Motion *vectors.TripleVector
	runtimeId uint64
	closed bool

	Position *vectors.TripleVector

	Level interfaces.ILevel
	Dimension interfaces.IDimension
	Rotation *math.Rotation

	NameTag string

	SpawnedTo map[uint64]interfaces.IPlayer

	mutex sync.Mutex

	EntityData map[uint32][]interface{}
}

func NewEntity(position *vectors.TripleVector, rotation *math.Rotation, motion *vectors.TripleVector, level interfaces.ILevel, dimension interfaces.IDimension) *Entity {
	RuntimeId++
	ent := Entity{
		data.NewAttributeMap(),
		motion,
		RuntimeId,
		false,
		position,
		level,
		dimension,
		rotation,
		"",
		map[uint64]interfaces.IPlayer{},
		sync.Mutex{},
		make(map[uint32][]interface{}),
	}

	ent.InitDataFlags()
	return &ent
}

/**
 * Returns the name tag of this entity.
 */
func (entity *Entity) GetNameTag() string {
	return entity.NameTag
}

/**
 * Sets the name tag of this entity.
 */
func (entity *Entity) SetNameTag(nameTag string) {
	entity.NameTag = nameTag
}

/**
 * Returns the attribute map of this entity.
 */
func (entity *Entity) GetAttributeMap() *data.AttributeMap {
	return entity.attributeMap
}

/**
 * Sets the attribute map of this entity.
 */
func (entity *Entity) SetAttributeMap(attMap *data.AttributeMap) {
	entity.attributeMap = attMap
}

/**
 * returns the entity data
 */
func (entity *Entity) GetEntityData() map[uint32][]interface{} {
	return entity.EntityData
}

/**
 * Initiates all entity data flags
 */
func (entity *Entity) InitDataFlags() {
	entity.EntityData[DataFlags] = append(entity.EntityData[DataFlags], uint32(data.Long))
	entity.EntityData[DataFlags] = append(entity.EntityData[DataFlags], int64(0))
	entity.SetDataFlag(AffectedByGravity, true)
}

/**
 * Sets entity data flag
 */
func (entity *Entity) SetDataFlag(flagId int, value bool)  {
	v := entity.EntityData[DataFlags][1].(int64)
	if value != entity.GetDataFlag(flagId) {
		v ^= int64(1 << uint(flagId))
		entity.EntityData[DataFlags][1] = v
	}
}

/**
 * Returns entity data flag
 */
func (entity *Entity) GetDataFlag(flagId int) bool {
	return (entity.EntityData[DataFlags][1].(int64) & (1 << uint(flagId))) > 0
}

/**
 * Returns the current position of this entity.
 */
func (entity *Entity) GetPosition() *vectors.TripleVector {
	return entity.Position
}

/**
 * Sets the position of this entity
 */
func (entity *Entity) SetPosition(v *vectors.TripleVector) {
	var newChunkX = int32(math2.Floor(float64(v.X))) >> 4
	var newChunkZ = int32(math2.Floor(float64(v.Z))) >> 4

	var oldChunk = entity.GetChunk()
	var newChunk = entity.GetDimension().GetChunk(newChunkX, newChunkZ)

	entity.Position = v

	if oldChunk != newChunk {
		newChunk.AddEntity(entity)
		entity.SpawnToAll()
		oldChunk.RemoveEntity(entity)
	}
}

/**
 * Returns the chunk this entity is currently in.
 */
func (entity *Entity) GetChunk() interfaces.IChunk {
	var x = int32(math2.Floor(float64(entity.Position.X))) >> 4
	var z = int32(math2.Floor(float64(entity.Position.Z))) >> 4
	return entity.GetDimension().GetChunk(x, z)
}

/**
 * Returns all players that have the chunk loaded in which this entity is.
 */
func (entity *Entity) GetViewers() map[uint64]interfaces.IPlayer {
	return entity.SpawnedTo
}

/**
 * Adds a viewer to this entity.
 */
func (entity *Entity) AddViewer(player interfaces.IPlayer) {
	entity.mutex.Lock()
	entity.SpawnedTo[player.GetRuntimeId()] = player
	entity.mutex.Unlock()
}

/**
 * Removes a viewer from this player.
 */
func (entity *Entity) RemoveViewer(player interfaces.IPlayer) {
	entity.mutex.Lock()
	delete(entity.SpawnedTo, player.GetRuntimeId())
	entity.mutex.Unlock()
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
func (entity *Entity) GetRotation() *math.Rotation {
	return entity.Rotation
}

/**
 * Sets the rotation of this entity
 */
func (entity *Entity) SetRotation(v *math.Rotation)  {
	entity.Rotation = v
}

/**
 * Returns the motion of this entity.
 */
func (entity *Entity) GetMotion() *vectors.TripleVector {
	return entity.Motion
}

/**
 * sets the motion of this entity
 */
func (entity *Entity) SetMotion(v *vectors.TripleVector)  {
	entity.Motion = v
}

/**
 * Returns the runtime ID of this entity.
 */
func (entity *Entity) GetRuntimeId() uint64 {
	return entity.runtimeId
}

/**
 * Returns the unique ID of this entity.
 * NOTE: This is currently unimplemented, and returns the runtime ID.
 */
func (entity *Entity) GetUniqueId() int64 {
	return int64(entity.runtimeId)
}

/**
 * Returns the entity ID of this entity.
 */
func (entity *Entity) GetEntityId() uint32 {
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
	entity.Level = nil
	entity.Dimension = nil
	entity.SpawnedTo = nil
}

/**
 * Returns the health points of this entity.
 */
func (entity *Entity) GetHealth() float32 {
	return entity.attributeMap.GetAttribute(data.AttributeHealth).GetValue()
}

/**
 * Sets the health points of this entity.
 */
func (entity *Entity) SetHealth(health float32) {
	entity.attributeMap.GetAttribute(data.AttributeHealth).SetValue(health)
}

/**
 * Kills the current entity.
 */
func (entity *Entity) Kill() {
	entity.SetHealth(0)
	//todo
}

/**
 * Spawns this entity to the given player.
 */
func (entity *Entity) SpawnTo(player interfaces.IPlayer)  {
	if !player.HasSpawned() {
		return
	}
	entity.AddViewer(player)
	player.SendAddEntity(entity)
}

/**
 * Despawns this entity from the given player.
 */
func (entity *Entity) DespawnFrom(player interfaces.IPlayer) {
	if !player.HasSpawned() {
		return
	}
	entity.RemoveViewer(player)
	player.SendRemoveEntity(entity)
}

/**
 * Despawns this entity from all players.
 */
func (entity *Entity) DespawnFromAll() {
	for _, p := range entity.GetLevel().GetServer().GetPlayerFactory().GetPlayers() {
		entity.DespawnFrom(p)
	}
}

/**
 * Spawns this entity to all players.
 */
func (entity *Entity) SpawnToAll()  {
	for _, p := range entity.GetChunk().GetViewers() {
		if p.GetRuntimeId() != entity.GetRuntimeId() {
			if _, ok := entity.SpawnedTo[p.GetRuntimeId()]; !ok {
				entity.SpawnTo(p)
			}
		}
	}
}

/**
 * Ticks the entity.
 */
func (entity *Entity) Tick()  {
	for runtimeId, player := range entity.GetViewers() {
		if player.IsClosed() {
			delete(entity.SpawnedTo, runtimeId)
		}
	}
}