package bedrock

import (
	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/data"
	data2 "github.com/irmine/worlds/entities/data"
)

type MovePlayerPacket struct {
	*packets.Packet
	RuntimeId            uint64
	Position             r3.Vector
	Rotation             data2.Rotation
	Mode                 byte
	OnGround             bool
	RidingRuntimeId      uint64
	TeleportCause, TeleportItem int32
}

func NewMovePlayerPacket() *MovePlayerPacket {
	return &MovePlayerPacket{Packet: packets.NewPacket(info.PacketIds[info.MovePlayerPacket]), Position: r3.Vector{}, Rotation: data2.Rotation{}}
}

func (pk *MovePlayerPacket) Encode() {
	pk.PutEntityRuntimeId(pk.RuntimeId)
	pk.PutVector(pk.Position)
	pk.PutPlayerRotation(pk.Rotation)
	pk.PutByte(pk.Mode)
	pk.PutBool(pk.OnGround)
	pk.PutEntityRuntimeId(pk.RidingRuntimeId)
	if pk.Mode == data.MoveTeleport {
		pk.PutLittleInt(pk.TeleportCause)
		pk.PutLittleInt(pk.TeleportItem)
	}
}

func (pk *MovePlayerPacket) Decode() {
	pk.RuntimeId = pk.GetEntityRuntimeId()
	pk.Position = pk.GetVector()
	pk.Rotation = pk.GetPlayerRotation()
	pk.Mode = pk.GetByte()
	pk.OnGround = pk.GetBool()
	pk.RidingRuntimeId = pk.GetEntityRuntimeId()
	if pk.Mode == data.MoveTeleport {
		pk.TeleportCause = pk.GetLittleInt()
		pk.TeleportItem = pk.GetLittleInt()
	}
}
