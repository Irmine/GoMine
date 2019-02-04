package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

const (
	RightClick = 1
	LeftClick = 2
	LeaveCehicle = 3
	MouseOver = 4
)

type InteractPacket struct {
	*packets.Packet
	Action byte
	RuntimeId uint64
}

func NewInteractPacket() *InteractPacket {
	return &InteractPacket{ packets.NewPacket(info.PacketIds[info.InteractPacket]), 0, 0}
}

func (pk *InteractPacket) Encode() {
	pk.PutByte(pk.Action)
	pk.PutUnsignedVarLong(pk.RuntimeId)
}

func (pk *InteractPacket) Decode() {
	pk.Action = pk.GetByte()
	pk.RuntimeId = pk.GetUnsignedVarLong()
}