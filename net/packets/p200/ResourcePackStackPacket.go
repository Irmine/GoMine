package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
	"gomine/net/packets/types"
)

type ResourcePackStackPacket struct {
	*packets.Packet
	MustAccept bool
	BehaviorPacks []types.ResourcePackStackEntry
	ResourcePacks []types.ResourcePackStackEntry
}

func NewResourcePackStackPacket() *ResourcePackStackPacket {
	return &ResourcePackStackPacket{packets.NewPacket(info.PacketIds200[info.ResourcePackStackPacket]), false, []types.ResourcePackStackEntry{}, []types.ResourcePackStackEntry{}}
}

func (pk *ResourcePackStackPacket) Encode() {
	pk.PutBool(pk.MustAccept)
	pk.PutPackStack(pk.BehaviorPacks)
	pk.PutPackStack(pk.ResourcePacks)
}

func (pk *ResourcePackStackPacket) Decode() {

}