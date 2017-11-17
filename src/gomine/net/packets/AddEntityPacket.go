package packets

import (
	"gomine/vectorMath"
	"gomine/net"
	"gomine/entities"
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
	return AddEntityPacket{NewPacket(net.AddEntityPacket), 0, 0, nil, nil, 0.0, 0.0, nil}
}

func (pk *AddEntityPacket) Encode() {
	pk.PutEId(pk.EntityId)
	pk.PutUnsignedVarInt(pk.EntityType)
	pk.PutTripleVectorObject(pk.Position)
	pk.PutTripleVectorObject(pk.Motion)
	pk.PutRotation(pk.Yaw)
	pk.PutRotation(pk.Pitch)
}