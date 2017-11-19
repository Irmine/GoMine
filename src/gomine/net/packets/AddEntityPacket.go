package packets

import (
	"gomine/vectorMath"
	"gomine/entities"
	"gomine/net/info"
)

type AddEntityPacket struct {
	*Packet
	EntityId uint64
	EntityType uint32
	Position vectorMath.TripleVector
	Motion vectorMath.TripleVector
	Yaw float32
	Pitch float32

	Attributes []entities.Attribute
}

func NewAddEntityPacket() AddEntityPacket {
	return AddEntityPacket{NewPacket(info.AddEntityPacket), 0, 0, vectorMath.TripleVector{}, vectorMath.TripleVector{}, 0.0, 0.0, nil}
}

func (pk AddEntityPacket) Encode() {
	pk.PutEId(pk.EntityId)
	pk.PutUnsignedVarInt(pk.EntityType)
	pk.PutTripleVectorObject(pk.Position)
	pk.PutTripleVectorObject(pk.Motion)
	pk.PutRotation(pk.Yaw)
	pk.PutRotation(pk.Pitch)
}