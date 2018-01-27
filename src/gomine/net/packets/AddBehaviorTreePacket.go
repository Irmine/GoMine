package packets

import (
	"gomine/net/info"
)

type AddBehaviorTreePacket struct {
	*Packet
	Unknown string
}

func NewAddBehaviorTreePacket() * AddBehaviorTreePacket{
	return &AddBehaviorTreePacket{NewPacket(info.AddBehaviorTreePacket), ""}
}

func (pk *AddBehaviorTreePacket) Decode() {
}

func (pk *AddBehaviorTreePacket) Encode() {
	pk.PutString(pk.Unknown)
}
