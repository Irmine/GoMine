package packets

import (
	"gomine/net/info"
	"gomine/interfaces"
)

type ResourcePackInfoPacket struct {
	*Packet
	MustAccept bool
	BehaviorPacks []interfaces.IPack
	ResourcePacks []interfaces.IPack
}

func NewResourcePackInfoPacket() *ResourcePackInfoPacket {
	return &ResourcePackInfoPacket{NewPacket(info.ResourcePackInfoPacket), false, []interfaces.IPack{}, []interfaces.IPack{}}
}

func (pk *ResourcePackInfoPacket) Encode()  {
	pk.PutBool(pk.MustAccept)
	pk.PutPacks(pk.ResourcePacks, true)
	pk.PutPacks(pk.ResourcePacks, true)
}

func (pk *ResourcePackInfoPacket) Decode()  {

}