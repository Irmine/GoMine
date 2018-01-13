package packets

import (
	"gomine/net/info"
)

type RemoveEntityPacket struct {
	*Packet
	EntityUniqueId int64
}

func NewRemoveEntityPacket() *RemoveEntityPacket {
	return &RemoveEntityPacket{NewPacket(info.RemoveEntityPacket), 0}
}

func (pk *RemoveEntityPacket) Encode() {
	pk.PutUniqueId(pk.EntityUniqueId)
}

func (pk *RemoveEntityPacket) Decode() {

}