package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
	"gomine/net/packets/types"
)

type ResourcePackInfoPacket struct {
	*packets.Packet
	MustAccept bool
	BehaviorPacks []types.ResourcePackInfoEntry
	ResourcePacks []types.ResourcePackInfoEntry
}

func NewResourcePackInfoPacket() *ResourcePackInfoPacket {
	return &ResourcePackInfoPacket{packets.NewPacket(info.PacketIds200[info.ResourcePackInfoPacket]), false, []types.ResourcePackInfoEntry{}, []types.ResourcePackInfoEntry{}}
}

func (pk *ResourcePackInfoPacket) Encode()  {
	pk.PutBool(pk.MustAccept)
	pk.PutPackInfo(pk.BehaviorPacks)
	pk.PutPackInfo(pk.ResourcePacks)
}

func (pk *ResourcePackInfoPacket) Decode()  {

}