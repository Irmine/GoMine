package packets

import (
	"gomine/net/info"
)

type SetEntityDataPacket struct {
	*Packet
	EntityId uint64
	EntityData map[uint32][]interface{}
}

func NewSetEntityDataPacket() *SetEntityDataPacket {
	return &SetEntityDataPacket{NewPacket(info.SetEntityDataPacket), 0, make(map[uint32][]interface{})}
}

func (pk *SetEntityDataPacket) Encode() {
	pk.PutRuntimeId(pk.EntityId)
	pk.PutEntityData(pk.EntityData)
}

func (pk *SetEntityDataPacket) Decode() {
	pk.EntityId = pk.GetRuntimeId()
	pk.EntityData = pk.GetEntityData()
}