package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type CraftingDataPacket struct {
	*packets.Packet
}

func NewCraftingDataPacket() *CraftingDataPacket {
	return &CraftingDataPacket{packets.NewPacket(info.PacketIds[info.CraftingDataPacket])}
}

func (pk *CraftingDataPacket) Encode() {
	pk.PutUnsignedVarInt(0)
	pk.PutBool(true)
}

func (pk *CraftingDataPacket) Decode() {

}
