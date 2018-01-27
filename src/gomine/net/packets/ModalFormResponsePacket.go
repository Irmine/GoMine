package packets

import (
	"gomine/net/info"
)

type ModalFormResponsePacket struct {
	*Packet
	FormId int32
	FormData string
}

func NewModalFormResponsePacket() * ModalFormResponsePacket{
	return &ModalFormResponsePacket{NewPacket(info.ModalFormResponsePacket), -012453, ""}
}

func (pk *ModalFormResponsePacket) Decode() {
	pk.FormId = pk.GetVarInt()
	pk.FormData = pk.GetString()
}

func (pk *ModalFormResponsePacket) Encode() {

}
