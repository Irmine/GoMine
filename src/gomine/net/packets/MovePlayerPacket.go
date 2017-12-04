package packets

import (
	"gomine/vectors"
	"gomine/net/info"
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
	Pitch, Yaw, HeadYaw float32
	Mode byte
	OnGround bool
	RidingEid uint64
	ExtraInt1, ExtraInt2 int32
}

func NewMovePlayerPacket() MovePlayerPacket {
	return MovePlayerPacket{Packet: NewPacket(info.MovePlayerPacket)}
}

func (pk *MovePlayerPacket) Encode() {
	pk.PutRuntimeId(pk.EntityId)
	pk.PutTripleVectorObject(pk.Position)
	pk.PutLittleFloat(pk.Pitch)
	pk.PutLittleFloat(pk.Yaw)
	pk.PutLittleFloat(pk.HeadYaw)
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
	pk.Position = *pk.GetTripleVectorObject()
	pk.Pitch = pk.GetLittleFloat()
	pk.Yaw = pk.GetLittleFloat()
	pk.HeadYaw = pk.GetLittleFloat()
	pk.Mode = pk.GetByte()
	pk.OnGround = pk.GetBool()
	pk.RidingEid = pk.GetRuntimeId()
	if pk.Mode == Teleport {
		pk.ExtraInt1 = pk.GetLittleInt()
		pk.ExtraInt2 = pk.GetLittleInt()
	}
}