package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

const (
	SwingArm = 1
	StopSleeping = 3
	CriticalHit = 4
)

type AnimatePacket struct {
	*packets.Packet
	Action int32
	RuntimeId uint64
	Float float32
}

func NewAnimatePacket() *AnimatePacket {
	return &AnimatePacket{ Packet: packets.NewPacket(info.PacketIds[info.AnimatePacket])}
}

func (pk *AnimatePacket) Encode() {
	pk.PutVarInt(pk.Action)
	pk.PutUnsignedVarLong(pk.RuntimeId)
	if uint(pk.Action) & 0x80 == 1 {
		pk.PutLittleFloat(pk.Float)
	}
}

func (pk *AnimatePacket) Decode() {
	pk.Action = pk.GetVarInt()
	pk.RuntimeId = pk.GetUnsignedVarLong()
	if uint(pk.Action) & 0x80 == 1 {
		pk.Float = pk.GetLittleFloat()
	}
}