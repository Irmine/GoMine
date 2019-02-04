package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type RemoveEntityPacket struct {
	*packets.Packet
	EntityUniqueId int64
}

func NewRemoveEntityPacket() *RemoveEntityPacket {
	return &RemoveEntityPacket{packets.NewPacket(info.PacketIds[info.RemoveEntityPacket]), 0}
}

func (pk *RemoveEntityPacket) Encode() {
	pk.PutEntityUniqueId(pk.EntityUniqueId)
}

func (pk *RemoveEntityPacket) Decode() {

}
