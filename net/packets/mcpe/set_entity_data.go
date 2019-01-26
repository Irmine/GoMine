package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type SetEntityDataPacket struct {
	*packets.Packet
	RuntimeId  uint64
	EntityData map[uint32][]interface{}
}

func NewSetEntityDataPacket() *SetEntityDataPacket {
	return &SetEntityDataPacket{packets.NewPacket(info.PacketIds[info.SetEntityDataPacket]), 0, make(map[uint32][]interface{})}
}

func (pk *SetEntityDataPacket) Encode() {
	pk.PutEntityRuntimeId(pk.RuntimeId)
	pk.PutEntityData(pk.EntityData)
}

func (pk *SetEntityDataPacket) Decode() {
	pk.RuntimeId = pk.GetEntityRuntimeId()
	pk.EntityData = pk.GetEntityData()
}
