package packets

import "gomine/net/info"

type ResourcePackStackPacket struct {
	*Packet
	MustAccept bool
}

func NewResourcePackStackPacket() *ResourcePackStackPacket {
	return &ResourcePackStackPacket{NewPacket(info.ResourcePackStackPacket), false}
}

func (pk *ResourcePackStackPacket) Encode() {

}

func (pk *ResourcePackStackPacket) Decode() {
	pk.PutBool(pk.MustAccept)
	pk.PutUnsignedVarInt(0)
	pk.PutUnsignedVarInt(0)
}
