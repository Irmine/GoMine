package bedrock

import (
	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	data2 "github.com/irmine/worlds/entities/data"
)

type MoveEntityPacket struct {
	*packets.Packet
	RuntimeId uint64
	Position  r3.Vector
	Rotation  data2.Rotation
	Flags     byte
}

func NewMoveEntityPacket() *MoveEntityPacket {
	return &MoveEntityPacket{Packet: packets.NewPacket(info.PacketIds[info.MoveEntityPacket]), Position: r3.Vector{}, Rotation: data2.Rotation{}, Flags: 0}
}

func (pk *MoveEntityPacket) Encode() {
	pk.PutEntityRuntimeId(pk.RuntimeId)
	pk.PutByte(pk.Flags)
	pk.PutVector(pk.Position)
	pk.PutEntityRotationBytes(pk.Rotation)
}

func (pk *MoveEntityPacket) Decode() {
	pk.RuntimeId = pk.GetEntityRuntimeId()
	pk.Flags = pk.GetByte()
	pk.Position = pk.GetVector()
	pk.Rotation = pk.GetEntityRotationBytes()
}
