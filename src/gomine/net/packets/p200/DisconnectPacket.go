package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type DisconnectPacket struct {
	*packets.Packet
	HideDisconnectionScreen bool
	Message string
}

func NewDisconnectPacket() *DisconnectPacket {
	return &DisconnectPacket{packets.NewPacket(info.PacketIds200[info.DisconnectPacket]), true, ""}
}

func (pk *DisconnectPacket) Encode() {
	pk.PutBool(pk.HideDisconnectionScreen)
	pk.PutString(pk.Message)
}

func (pk *DisconnectPacket) Decode() {
	pk.HideDisconnectionScreen = pk.GetBool()
	pk.Message = pk.GetString()
}