package p200

import (
	"github.com/irmine/gomine/entities/data"
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/golang/geo/r3"
)

type AddEntityPacket struct {
	*packets.Packet
	UniqueId   int64
	RuntimeId  uint64
	EntityType uint32
	Position   r3.Vector
	Motion     r3.Vector
	Rotation   math.Rotation

	Attributes *data.AttributeMap
	EntityData map[uint32][]interface{}
}

func NewAddEntityPacket() *AddEntityPacket {
	return &AddEntityPacket{packets.NewPacket(info.PacketIds200[info.AddEntityPacket]), 0, 0, 0, r3.Vector{}, r3.Vector{}, math.Rotation{}, data.NewAttributeMap(), nil}
}

func (pk *AddEntityPacket) Encode() {
	pk.PutUniqueId(pk.UniqueId)
	pk.PutRuntimeId(pk.RuntimeId)
	pk.PutUnsignedVarInt(pk.EntityType)
	pk.PutVector(pk.Position)
	pk.PutVector(pk.Motion)
	pk.PutRotationObject(pk.Rotation, false)
	pk.PutEntityAttributeMap(pk.Attributes)
	pk.PutEntityData(pk.EntityData)
	pk.PutUnsignedVarInt(0)
}

func (pk *AddEntityPacket) Decode() {
	pk.UniqueId = pk.GetUniqueId()
	pk.RuntimeId = pk.GetRuntimeId()
	pk.EntityType = pk.GetUnsignedVarInt()
	pk.Position = pk.GetVector()
	pk.Motion = pk.GetVector()
	pk.Rotation = pk.GetRotationObject(false)
	pk.Attributes = pk.GetEntityAttributeMap()
	pk.EntityData = pk.GetEntityData()
}
