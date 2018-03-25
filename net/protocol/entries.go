package protocol

import (
	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/worlds"
	"github.com/irmine/worlds/entities/data"
)

type AddEntityEntry interface {
	GetUniqueId() int64
	GetRuntimeId() uint64
	GetEntityType() uint32
	GetPosition() r3.Vector
	GetMotion() r3.Vector
	GetRotation() data.Rotation
	GetAttributeMap() data.AttributeMap
	GetEntityData() map[uint32][]interface{}
}

type AddPlayerEntry interface {
	AddEntityEntry
	GetDisplayName() string
	GetName() string
}

type PlayerListEntry interface {
	AddPlayerEntry
	GetXUID() string
	GetUUID() utils.UUID
	GetSkinId() string
	GetSkinData() []byte
	GetCapeData() []byte
	GetGeometryName() string
	GetGeometryData() string
	GetPlatform() int32
}

type StartGameEntry interface {
	GetRuntimeId() uint64
	GetUniqueId() int64
	GetPosition() r3.Vector
	GetDimension() *worlds.Dimension
}
