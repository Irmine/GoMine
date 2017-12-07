package packets

import (
	"gomine/vectors"
	"gomine/entities"
	"gomine/net/info"
	"gomine/worlds/locations"
)

type AddEntityPacket struct {
	*Packet
	EntityId   uint64
	EntityType uint32
	Position   locations.EntityPosition
	Motion     vectors.TripleVector
	Yaw        float32
	Pitch      float32

	Attributes map[int]entities.Attribute
	EntityData map[uint32][]interface{}
}

func NewAddEntityPacket() *AddEntityPacket {
	return &AddEntityPacket{NewPacket(info.AddEntityPacket), 0, 0, locations.EntityPosition{}, vectors.TripleVector{}, 0.0, 0.0, nil, nil}
}

func (pk *AddEntityPacket) Encode() {
	pk.PutRuntimeId(pk.EntityId)
	pk.PutUnsignedVarInt(pk.EntityType)
	pk.PutTripleVectorObject(*pk.Position.AsTripleVector())
	pk.PutTripleVectorObject(pk.Motion)
	pk.PutRotationObject(pk.Position.Rotation)
	pk.PutEntityAttributes(pk.Attributes)
	pk.PutEntityData(pk.EntityData)
}

func (pk *AddEntityPacket) Decode() {
	pk.EntityId = pk.GetRuntimeId()
	pk.EntityType = pk.GetUnsignedVarInt()
	pk.Position.SetVector(*pk.GetTripleVectorObject())
	pk.Motion = *pk.GetTripleVectorObject()
	pk.Position.Rotation = pk.GetRotationObject()
	pk.Attributes = pk.GetEntityAttributes()
	pk.EntityData = pk.GetEntityData()
}