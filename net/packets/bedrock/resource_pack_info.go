package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/types"
)

type ResourcePackInfoPacket struct {
	*packets.Packet
	MustAccept    bool
	BehaviorPacks []types.ResourcePackInfoEntry
	ResourcePacks []types.ResourcePackInfoEntry
	Bool1         bool
}

func NewResourcePackInfoPacket() *ResourcePackInfoPacket {
	return &ResourcePackInfoPacket{packets.NewPacket(info.PacketIds[info.ResourcePackInfoPacket]), false, []types.ResourcePackInfoEntry{}, []types.ResourcePackInfoEntry{}, false}
}

func (pk *ResourcePackInfoPacket) Encode() {
	pk.PutBool(pk.MustAccept)
	pk.PutBool(pk.Bool1)
	pk.PutPackInfo(pk.BehaviorPacks)
	pk.PutPackInfo(pk.ResourcePacks)
}

func (pk *ResourcePackInfoPacket) Decode() {

}
