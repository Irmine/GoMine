package packets

import (
	"gomine/vectors"
	"gomine/entities"
	"gomine/net/info"
	"gomine/entities/math"
)

type AddEntityPacket struct {
	*Packet
	UniqueId int64
	RuntimeId  uint64
	EntityType uint32
	Position   vectors.TripleVector
	Motion     vectors.TripleVector
	Rotation   math.Rotation

	Attributes map[int]entities.Attribute
	EntityData map[uint32][]interface{}
}

func NewAddEntityPacket() *AddEntityPacket {
	return &AddEntityPacket{NewPacket(info.AddEntityPacket), 0, 0, 0, vectors.TripleVector{}, vectors.TripleVector{}, math.Rotation{}, nil, nil}
}

func (pk *AddEntityPacket) Encode() {
	pk.PutUniqueId(pk.UniqueId)
	pk.PutRuntimeId(pk.RuntimeId)
	pk.PutUnsignedVarInt(pk.EntityType)
	pk.PutTripleVectorObject(pk.Position)
	pk.PutTripleVectorObject(pk.Motion)
	pk.PutRotationObject(pk.Rotation, false)
	pk.PutEntityAttributes(pk.Attributes)
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
	pk.Attributes = pk.GetEntityAttributes()
	pk.EntityData = pk.GetEntityData()
}