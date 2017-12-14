package packets

import (
	"gomine/net/info"
	"gomine/interfaces"
)

type ResourcePackStackPacket struct {
	*Packet
	MustAccept bool
	BehaviorPacks []interfaces.IPack
	ResourcePacks []interfaces.IPack
}

func NewResourcePackStackPacket() *ResourcePackStackPacket {
	return &ResourcePackStackPacket{NewPacket(info.ResourcePackStackPacket), false, []interfaces.IPack{}, []interfaces.IPack{}}
}

func (pk *ResourcePackStackPacket) Encode() {
	pk.PutBool(pk.MustAccept)
	pk.PutPacks(pk.BehaviorPacks, false)
	pk.PutPacks(pk.ResourcePacks, false)
}

func (pk *ResourcePackStackPacket) Decode() {

}