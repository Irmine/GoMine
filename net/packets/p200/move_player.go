package p200

import (
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/golang/geo/r3"
)

type MovePlayerPacket struct {
	*packets.Packet
	RuntimeId            uint64
	Position             r3.Vector
	Rotation             math.Rotation
	Mode                 byte
	OnGround             bool
	RidingRuntimeId      uint64
	ExtraInt1, ExtraInt2 int32
}

func NewMovePlayerPacket() *MovePlayerPacket {
	return &MovePlayerPacket{Packet: packets.NewPacket(info.PacketIds200[info.MovePlayerPacket]), Position: r3.Vector{}, Rotation: *math.NewRotation(0, 0, 0)}
}

func (pk *MovePlayerPacket) Encode() {
	pk.PutRuntimeId(pk.RuntimeId)
	pk.PutVector(pk.Position)
	pk.PutRotationObject(pk.Rotation, true)
	pk.PutByte(pk.Mode)
	pk.PutBool(pk.OnGround)
	pk.PutRuntimeId(pk.RidingRuntimeId)
	if pk.Mode == data.MoveTeleport {
		pk.PutLittleInt(pk.ExtraInt1)
		pk.PutLittleInt(pk.ExtraInt2)
	}
}

func (pk *MovePlayerPacket) Decode() {
	pk.RuntimeId = pk.GetRuntimeId()
	pk.Position = pk.GetVector()
	pk.Rotation = pk.GetRotationObject(true)
	pk.Mode = pk.GetByte()
	pk.OnGround = pk.GetBool()
	pk.RidingRuntimeId = pk.GetRuntimeId()
	if pk.Mode == data.MoveTeleport {
		pk.ExtraInt1 = pk.GetLittleInt()
		pk.ExtraInt2 = pk.GetLittleInt()
	}
}
