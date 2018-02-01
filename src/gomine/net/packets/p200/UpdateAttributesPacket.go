package p200

import (
	"gomine/net/info"
	"gomine/entities/data"
	"gomine/net/packets"
)

type UpdateAttributesPacket struct {
	*packets.Packet
	RuntimeId uint64
	Attributes *data.AttributeMap
}

func NewUpdateAttributesPacket() *UpdateAttributesPacket {
	return &UpdateAttributesPacket{packets.NewPacket(info.UpdateAttributesPacket), 0, data.NewAttributeMap()}
}

func (pk *UpdateAttributesPacket) Encode() {
	pk.PutRuntimeId(pk.RuntimeId)
	pk.PutEntityAttributeMap(pk.Attributes)
}

func (pk *UpdateAttributesPacket) Decode() {
	pk.RuntimeId = pk.GetRuntimeId()
	pk.Attributes = pk.GetEntityAttributeMap()
}