package packets

import (
	"gomine/entities"
	"gomine/net/info"
)

type UpdateAttributesPacket struct {
	*Packet
	EntityId uint64
	Attributes *entities.AttributeMap
}

func NewUpdateAttributesPacket() *UpdateAttributesPacket {
	return &UpdateAttributesPacket{NewPacket(info.UpdateAttributesPacket), 0, entities.NewAttributeMap()}
}

func (pk *UpdateAttributesPacket) Encode() {
	pk.PutRuntimeId(pk.EntityId)
	pk.PutEntityAttributeMap(pk.Attributes)
}

func (pk *UpdateAttributesPacket) Decode() {
	pk.EntityId = pk.GetRuntimeId()
	pk.Attributes = pk.GetEntityAttributeMap()
}