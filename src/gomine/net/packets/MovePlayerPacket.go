package packets

import (
	"gomine/net/info"
	"gomine/vectors"
	"gomine/entities/math"
)

const (
	MoveNormal = iota
	MoveReset
	MoveTeleport
	MovePitch
)

type MovePlayerPacket struct {
	*Packet
	EntityId uint64
	Position vectors.TripleVector
	Rotation math.Rotation
	Mode byte
	OnGround bool
	RidingEid uint64
	ExtraInt1, ExtraInt2 int32
}

func NewMovePlayerPacket() *MovePlayerPacket {
	return &MovePlayerPacket{Packet: NewPacket(info.MovePlayerPacket), Position: *vectors.NewTripleVector(0, 0, 0), Rotation: *math.NewRotation(0, 0, 0)}
}

func (pk *MovePlayerPacket) Encode() {
	pk.PutRuntimeId(pk.EntityId)
	pk.PutTripleVectorObject(*pk.Position.AsTripleVector())
	pk.PutRotationObject(pk.Rotation, true)
	pk.PutByte(pk.Mode)
	pk.PutBool(pk.OnGround)
	pk.PutRuntimeId(pk.RidingEid)
	if pk.Mode == MoveTeleport {
		pk.PutLittleInt(pk.ExtraInt1)
		pk.PutLittleInt(pk.ExtraInt2)
	}
}

func (pk *MovePlayerPacket) Decode() {
	pk.EntityId = pk.GetRuntimeId()
	pk.Position.SetVector(pk.GetTripleVectorObject())
	pk.Rotation = pk.GetRotationObject(true)
	pk.Mode = pk.GetByte()
	pk.OnGround = pk.GetBool()
	pk.RidingEid = pk.GetRuntimeId()
	if pk.Mode == MoveTeleport {
		pk.ExtraInt1 = pk.GetLittleInt()
		pk.ExtraInt2 = pk.GetLittleInt()
	}
}