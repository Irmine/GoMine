package interfaces

import (
	"github.com/irmine/gomine/entities/data"
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/vectors"
)

type IEntity interface {
	GetNameTag() string
	SetNameTag(string)
	IsClosed() bool
	Close()
	GetHealth() float32
	SetHealth(float32)
	Kill()
	Tick()
	GetRuntimeId() uint64
	GetPosition() *vectors.TripleVector
	SetPosition(*vectors.TripleVector)
	GetDimension() IDimension
	SetDimension(IDimension)
	GetLevel() ILevel
	SetLevel(level ILevel)
	GetRotation() *math.Rotation
	SetRotation(*math.Rotation)
	GetMotion() *vectors.TripleVector
	SetMotion(*vectors.TripleVector)
	SpawnTo(IPlayer)
	SpawnToAll()
	DespawnFrom(IPlayer)
	DespawnFromAll()
	GetViewers() map[uint64]IPlayer
	AddViewer(IPlayer)
	RemoveViewer(IPlayer)
	GetUniqueId() int64
	GetEntityId() uint32
	GetEntityData() map[uint32][]interface{}
	GetAttributeMap() *data.AttributeMap
}
