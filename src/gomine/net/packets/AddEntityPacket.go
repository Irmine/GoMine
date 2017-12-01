package packets

import (
	"gomine/vectors"
	"gomine/entities"
	"gomine/net/info"
)

type AddEntityPacket struct {
	*Packet
	EntityId   uint64
	EntityType uint32
	Position   vectors.TripleVector
	Motion     vectors.TripleVector
	Yaw        float32
	Pitch      float32

	Attributes map[int]entities.Attribute
	EntityData map[uint32][]interface{}
}

func NewAddEntityPacket() *AddEntityPacket {
	return &AddEntityPacket{NewPacket(info.AddEntityPacket), 0, 0, vectors.TripleVector{}, vectors.TripleVector{}, 0.0, 0.0, nil, nil}
}

func (pk *AddEntityPacket) Encode() {
	pk.PutRuntimeId(pk.EntityId)
	pk.PutUnsignedVarInt(pk.EntityType)
	pk.PutTripleVectorObject(pk.Position)
	pk.PutTripleVectorObject(pk.Motion)
	pk.PutRotation(pk.Yaw)
	pk.PutRotation(pk.Pitch)
	pk.PutEntityAttributes(pk.Attributes)
	pk.PutEntityData(pk.EntityData)
}

func (pk *AddEntityPacket) Decode() {
	pk.EntityId = pk.GetRuntimeId()
	pk.EntityType = pk.GetUnsignedVarInt()
	pk.Position = *pk.GetTripleVectorObject()
	pk.Motion = *pk.GetTripleVectorObject()
	pk.Yaw = pk.GetRotation()
	pk.Pitch = pk.GetRotation()
	pk.Attributes = pk.GetEntityAttributes()
	pk.EntityData = pk.GetEntityData()
}