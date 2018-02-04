package p200

import (
	"gomine/vectors"
	"gomine/net/info"
	"gomine/entities/math"
	"gomine/entities/data"
	"gomine/net/packets"
)

type AddEntityPacket struct {
	*packets.Packet
	UniqueId int64
	RuntimeId  uint64
	EntityType uint32
	Position   vectors.TripleVector
	Motion     vectors.TripleVector
	Rotation   math.Rotation

	Attributes *data.AttributeMap
	EntityData map[uint32][]interface{}
}

func NewAddEntityPacket() *AddEntityPacket {
	return &AddEntityPacket{packets.NewPacket(info.PacketIds200[info.AddEntityPacket]), 0, 0, 0, vectors.TripleVector{}, vectors.TripleVector{}, math.Rotation{}, data.NewAttributeMap(), nil}
}

func (pk *AddEntityPacket) Encode() {
	pk.PutUniqueId(pk.UniqueId)
	pk.PutRuntimeId(pk.RuntimeId)
	pk.PutUnsignedVarInt(pk.EntityType)
	pk.PutTripleVectorObject(pk.Position)
	pk.PutTripleVectorObject(pk.Motion)
	pk.PutRotationObject(pk.Rotation, false)
	pk.PutEntityAttributeMap(pk.Attributes)
	pk.PutEntityData(pk.EntityData)
	pk.PutUnsignedVarInt(0)
}

func (pk *AddEntityPacket) Decode() {
	pk.UniqueId = pk.GetUniqueId()
	pk.RuntimeId = pk.GetRuntimeId()
	pk.EntityType = pk.GetUnsignedVarInt()
	pk.Position.SetVector(pk.GetTripleVectorObject())
	pk.Motion = *pk.GetTripleVectorObject()
	pk.Rotation = pk.GetRotationObject(false)
	pk.Attributes = pk.GetEntityAttributeMap()
	pk.EntityData = pk.GetEntityData()
}