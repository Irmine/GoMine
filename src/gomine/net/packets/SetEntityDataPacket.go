package packets

import (
	"gomine/net/info"
)

type SetEntityDataPacket struct {
	*Packet
	RuntimeId uint64
	EntityData map[uint32][]interface{}
}

func NewSetEntityDataPacket() *SetEntityDataPacket {
	return &SetEntityDataPacket{NewPacket(info.SetEntityDataPacket), 0, make(map[uint32][]interface{})}
}

func (pk *SetEntityDataPacket) Encode() {
	pk.PutRuntimeId(pk.RuntimeId)
	pk.PutEntityData(pk.EntityData)
}

func (pk *SetEntityDataPacket) Decode() {
	pk.RuntimeId = pk.GetRuntimeId()
	pk.EntityData = pk.GetEntityData()
}