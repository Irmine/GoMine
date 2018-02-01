package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type CraftingDataPacket struct {
	*packets.Packet
}

func NewCraftingDataPacket() *CraftingDataPacket {
	return &CraftingDataPacket{packets.NewPacket(info.CraftingDataPacket)}
}

func (pk *CraftingDataPacket) Encode() {
	pk.PutUnsignedVarInt(0)
	pk.PutBool(true)
}

func (pk *CraftingDataPacket) Decode() {

}