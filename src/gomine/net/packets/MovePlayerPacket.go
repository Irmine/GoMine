package packets

import (
	"gomine/net/info"
	"gomine/vectors"
	"gomine/players/math"
)

const (
	Normal = iota + 0
	Reset
	Teleport
	Pitch
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
	return &MovePlayerPacket{Packet: NewPacket(info.MovePlayerPacket), Position: *vectors.NewTripleVector(0, 0, 0), Rotation: math.NewRotation(0, 0, 0)}
}

func (pk *MovePlayerPacket) Encode() {
	pk.PutRuntimeId(pk.EntityId)
	pk.PutTripleVectorObject(*pk.Position.AsTripleVector())
	pk.PutRotationObject(pk.Rotation)
	pk.PutByte(pk.Mode)
	pk.PutBool(pk.OnGround)
	pk.PutRuntimeId(pk.RidingEid)
	if pk.Mode == Teleport {
		pk.PutLittleInt(pk.ExtraInt1)
		pk.PutLittleInt(pk.ExtraInt2)
	}
}

func (pk *MovePlayerPacket) Decode() {
	pk.EntityId = pk.GetRuntimeId()
	pk.Position.SetVector(pk.GetTripleVectorObject())
	pk.Rotation = pk.GetRotationObject()
	pk.Mode = pk.GetByte()
	pk.OnGround = pk.GetBool()
	pk.RidingEid = pk.GetRuntimeId()
	if pk.Mode == Teleport {
		pk.ExtraInt1 = pk.GetLittleInt()
		pk.ExtraInt2 = pk.GetLittleInt()
	}
}