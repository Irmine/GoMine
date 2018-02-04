package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type PlayStatusPacket struct {
	*packets.Packet
	Status int32
}

func NewPlayStatusPacket() *PlayStatusPacket {
	return &PlayStatusPacket{packets.NewPacket(info.PacketIds200[info.PlayStatusPacket]), 0}
}

func (pk *PlayStatusPacket) Encode()  {
	pk.PutInt(pk.Status)
}

func (pk *PlayStatusPacket) Decode()  {
	pk.Status = pk.GetInt()
}