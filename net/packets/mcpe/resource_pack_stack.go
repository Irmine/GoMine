package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/types"
)

type ResourcePackStackPacket struct {
	*packets.Packet
	MustAccept    bool
	BehaviorPacks []types.ResourcePackStackEntry
	ResourcePacks []types.ResourcePackStackEntry
}

func NewResourcePackStackPacket() *ResourcePackStackPacket {
	return &ResourcePackStackPacket{packets.NewPacket(info.PacketIds[info.ResourcePackStackPacket]), false, []types.ResourcePackStackEntry{}, []types.ResourcePackStackEntry{}}
}

func (pk *ResourcePackStackPacket) Encode() {
	pk.PutBool(pk.MustAccept)
	pk.PutPackStack(pk.BehaviorPacks)
	pk.PutPackStack(pk.ResourcePacks)
}

func (pk *ResourcePackStackPacket) Decode() {

}
