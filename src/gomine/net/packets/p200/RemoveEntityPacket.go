package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type RemoveEntityPacket struct {
	*packets.Packet
	EntityUniqueId int64
}

func NewRemoveEntityPacket() *RemoveEntityPacket {
	return &RemoveEntityPacket{packets.NewPacket(info.RemoveEntityPacket), 0}
}

func (pk *RemoveEntityPacket) Encode() {
	pk.PutUniqueId(pk.EntityUniqueId)
}

func (pk *RemoveEntityPacket) Decode() {

}