package packets

import (
	"gomine/net/info"
	"gomine/worlds/locations"
	"fmt"
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
	Position locations.EntityPosition
	Mode byte
	OnGround bool
	RidingEid uint64
	ExtraInt1, ExtraInt2 int32
}

func NewMovePlayerPacket() *MovePlayerPacket {
	return &MovePlayerPacket{Packet: NewPacket(info.MovePlayerPacket)}
}

func (pk *MovePlayerPacket) Encode() {
	pk.PutRuntimeId(pk.EntityId)
	pk.PutTripleVectorObject(*pk.Position.AsTripleVector())
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
	fmt.Println("vec : ", *pk.GetTripleVectorObject())
	pk.Position.SetVector(*pk.GetTripleVectorObject())
	pk.Position.SetRotation(pk.GetRotationObject())
	pk.Mode = pk.GetByte()
	pk.OnGround = pk.GetBool()
	pk.RidingEid = pk.GetRuntimeId()
	if pk.Mode == Teleport {
		pk.ExtraInt1 = pk.GetLittleInt()
		pk.ExtraInt2 = pk.GetLittleInt()
	}
}