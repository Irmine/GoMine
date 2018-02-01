package p200

import (
	"gomine/net/info"
	"gomine/interfaces"
	"gomine/net/packets"
)

type ResourcePackInfoPacket struct {
	*packets.Packet
	MustAccept bool
	BehaviorPacks []interfaces.IPack
	ResourcePacks []interfaces.IPack
}

func NewResourcePackInfoPacket() *ResourcePackInfoPacket {
	return &ResourcePackInfoPacket{packets.NewPacket(info.ResourcePackInfoPacket), false, []interfaces.IPack{}, []interfaces.IPack{}}
}

func (pk *ResourcePackInfoPacket) Encode()  {
	pk.PutBool(pk.MustAccept)
	pk.PutPacks(pk.BehaviorPacks, true)
	pk.PutPacks(pk.ResourcePacks, true)
}

func (pk *ResourcePackInfoPacket) Decode()  {

}