package entities

import (
	math2 "math"
	"sync"

	"github.com/irmine/gomine/entities/data"
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/interfaces"
	"github.com/golang/geo/r3"
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
	Motion       r3.Vector
	runtimeId    uint64
	closed       bool
	Position     r3.Vector
	Level        interfaces.ILevel
	Dimension    interfaces.IDimension
	Rotation     *math.Rotation
	NameTag      string
	SpawnedTo    map[uint64]interfaces.IPlayer
	mutex        sync.Mutex
	EntityData   map[uint32][]interface{}
}

func NewEntity(position r3.Vector, rotation *math.Rotation, motion r3.Vector, level interfaces.ILevel, dimension interfaces.IDimension) *Entity {
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
		make(map[uint64]interfaces.IPlayer),
		sync.Mutex{},
		make(map[uint32][]interface{}),
	}

	ent.InitDataFlags()
	return &ent
}

// GetNameTag returns the name tag of this entity.
func (entity *Entity) GetNameTag() string {
	return entity.NameTag
}

// SetNameTag sets the name tag of this entity.
func (entity *Entity) SetNameTag(nameTag string) {
	entity.NameTag = nameTag
}

// GetAttributeMap returns the attribute map of this entity.
func (entity *Entity) GetAttributeMap() *data.AttributeMap {
	return entity.attributeMap
}

// SetAttributeMap sets the attribute map of this entity.
func (entity *Entity) SetAttributeMap(attMap *data.AttributeMap) {
	entity.attributeMap = attMap
}

// GetEntityData returns the entity data map.
func (entity *Entity) GetEntityData() map[uint32][]interface{} {
	return entity.EntityData
}

// InitDataFlags initializes all default data flags.
func (entity *Entity) InitDataFlags() {
	entity.EntityData[DataFlags] = append(entity.EntityData[DataFlags], uint32(data.Long))
	entity.EntityData[DataFlags] = append(entity.EntityData[DataFlags], int64(0))
	entity.SetDataFlag(AffectedByGravity, true)
}

// SetDataFlag sets the given data flag to the given value.
func (entity *Entity) SetDataFlag(flagId int, value bool) {
	v := entity.EntityData[DataFlags][1].(int64)
	if value != entity.GetDataFlag(flagId) {
		v ^= int64(1 << uint(flagId))
		entity.EntityData[DataFlags][1] = v
	}
}

// GetDataFlag returns teh value of the given flag Id.
func (entity *Entity) GetDataFlag(flagId int) bool {
	return (entity.EntityData[DataFlags][1].(int64) & (1 << uint(flagId))) > 0
}

// GetPosition returns the current position of this entity.
func (entity *Entity) GetPosition() r3.Vector {
	return entity.Position
}

// SetPosition sets the position of this entity
func (entity *Entity) SetPosition(v r3.Vector) {
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

// GetChunk returns the chunk this entity is currently in.
func (entity *Entity) GetChunk() interfaces.IChunk {
	var x = int32(math2.Floor(float64(entity.Position.X))) >> 4
	var z = int32(math2.Floor(float64(entity.Position.Z))) >> 4
	return entity.GetDimension().GetChunk(x, z)
}

// GetViewers returns all players that have the chunk loaded in which this entity is.
func (entity *Entity) GetViewers() map[uint64]interfaces.IPlayer {
	return entity.SpawnedTo
}

// AddViewer adds a viewer to this entity.
func (entity *Entity) AddViewer(player interfaces.IPlayer) {
	entity.mutex.Lock()
	entity.SpawnedTo[player.GetRuntimeId()] = player
	entity.mutex.Unlock()
}

// RemoveViewer removes a viewer from this entity.
func (entity *Entity) RemoveViewer(player interfaces.IPlayer) {
	entity.mutex.Lock()
	delete(entity.SpawnedTo, player.GetRuntimeId())
	entity.mutex.Unlock()
}

// GetLevel returns the level of this entity.
func (entity *Entity) GetLevel() interfaces.ILevel {
	return entity.Level
}

// SetLevel sets the level of this entity.
func (entity *Entity) SetLevel(v interfaces.ILevel) {
	entity.Level = v
}

// GetDimension returns the dimension of this entity.
func (entity *Entity) GetDimension() interfaces.IDimension {
	return entity.Dimension
}

// SetDimension sets the dimension of the entity.
func (entity *Entity) SetDimension(v interfaces.IDimension) {
	entity.Dimension = v
}

// GetRotation returns the current rotation of this entity.
func (entity *Entity) GetRotation() *math.Rotation {
	return entity.Rotation
}

// SetRotation sets the rotation of this entity.
func (entity *Entity) SetRotation(v *math.Rotation) {
	entity.Rotation = v
}

// GetMotion returns the motion of this entity.
func (entity *Entity) GetMotion() r3.Vector {
	return entity.Motion
}

// SetMotion sets the motion of this entity.
func (entity *Entity) SetMotion(v r3.Vector) {
	entity.Motion = v
}

// GetRuntimeId returns the runtime ID of this entity.
func (entity *Entity) GetRuntimeId() uint64 {
	return entity.runtimeId
}

// GetUniqueId returns the unique ID of this entity.
// NOTE: This is currently unimplemented, and returns the runtime ID.
func (entity *Entity) GetUniqueId() int64 {
	return int64(entity.runtimeId)
}

// GetEntityId returns the entity ID of this entity.
func (entity *Entity) GetEntityId() uint32 {
	return 0
}

// IsClosed checks if the entity is closed and not to be used anymore.
func (entity *Entity) IsClosed() bool {
	return entity.closed
}

// Close closes the entity making it unable to be used.
func (entity *Entity) Close() {
	entity.closed = true
	entity.Level = nil
	entity.Dimension = nil
	entity.SpawnedTo = nil
}

// GetHealth returns the health points of this entity.
func (entity *Entity) GetHealth() float32 {
	return entity.attributeMap.GetAttribute(data.AttributeHealth).GetValue()
}

// SetHealth sets the health points of this entity.
func (entity *Entity) SetHealth(health float32) {
	entity.attributeMap.GetAttribute(data.AttributeHealth).SetValue(health)
}

// Kill kills the entity.
func (entity *Entity) Kill() {
	entity.SetHealth(0)
	//todo
}

// SpawnTo spawns this entity to the given player.
func (entity *Entity) SpawnTo(player interfaces.IPlayer) {
	if !player.HasSpawned() {
		return
	}
	if entity.GetRuntimeId() == player.GetRuntimeId() {
		return
	}
	entity.AddViewer(player)
	player.SendAddEntity(entity)
}

// DespawnFrom despawns this entity from the given player.
func (entity *Entity) DespawnFrom(player interfaces.IPlayer) {
	if !player.HasSpawned() {
		return
	}
	entity.RemoveViewer(player)
	player.SendRemoveEntity(entity)
}

// DespawnFromAll despawns this entity from all players.
func (entity *Entity) DespawnFromAll() {
	for _, p := range entity.GetLevel().GetServer().GetPlayerFactory().GetPlayers() {
		entity.DespawnFrom(p)
	}
}

// SpawnToAll spawns this entity to all players.
func (entity *Entity) SpawnToAll() {
	for _, p := range entity.GetChunk().GetViewers() {
		if p.GetRuntimeId() != entity.GetRuntimeId() {
			if _, ok := entity.SpawnedTo[p.GetRuntimeId()]; !ok {
				entity.SpawnTo(p)
			}
		}
	}
}

// Tick ticks the entity.
func (entity *Entity) Tick() {
	for runtimeId, player := range entity.GetViewers() {
		if player.IsClosed() {
			delete(entity.SpawnedTo, runtimeId)
		}
	}
}
