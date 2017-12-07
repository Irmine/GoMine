package packets

import "gomine/net/info"

type CraftingDataPacket struct {
	*Packet
}

func NewCraftingDataPacket() *CraftingDataPacket {
	return &CraftingDataPacket{NewPacket(info.CraftingDataPacket)}
}

func (pk *CraftingDataPacket) Encode() {
	pk.PutUnsignedVarInt(0)
	pk.PutBool(true)
}

func (pk *CraftingDataPacket) Decode() {

}