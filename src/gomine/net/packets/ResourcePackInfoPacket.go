package packets

import "gomine/net/info"

type ResourcePackInfoPacket struct {
	*Packet
	MustAccept bool
	// Behavior pack entries
	// Resource pack entries
}

func NewResourcePackInfoPacket() *ResourcePackInfoPacket {
	return &ResourcePackInfoPacket{NewPacket(info.ResourcePackInfoPacket), false}
}

func (pk *ResourcePackInfoPacket) Encode()  {
	pk.PutBool(pk.MustAccept)
	pk.PutLittleShort(0)
	pk.PutLittleShort(0)
}

func (pk *ResourcePackInfoPacket) Decode()  {

}