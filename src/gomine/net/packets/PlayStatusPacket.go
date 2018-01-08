package packets

import (
	"gomine/net/info"
)

type PlayStatusPacket struct {
	*Packet
	Status int32
}

func NewPlayStatusPacket() *PlayStatusPacket {
	return &PlayStatusPacket{NewPacket(info.PlayStatusPacket), 0}
}

func (pk *PlayStatusPacket) Encode()  {
	pk.PutInt(pk.Status)
}

func (pk *PlayStatusPacket) Decode()  {
	pk.Status = pk.GetInt()
}