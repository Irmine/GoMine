package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type PlayStatusPacket struct {
	*packets.Packet
	Status int32
}

func NewPlayStatusPacket() *PlayStatusPacket {
	return &PlayStatusPacket{packets.NewPacket(info.PacketIds[info.PlayStatusPacket]), 0}
}

func (pk *PlayStatusPacket) Encode() {
	pk.PutInt(pk.Status)
}

func (pk *PlayStatusPacket) Decode() {
	pk.Status = pk.GetInt()
}
