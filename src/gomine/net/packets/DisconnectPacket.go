package packets

import (
	"gomine/net/info"
)

type DisconnectPacket struct {
	*Packet
	HideDisconnectionScreen bool
	Message string
}

func NewDisconnectPacket() DisconnectPacket {
	return DisconnectPacket{NewPacket(info.DisconnectPacket), true, ""}
}

func (pk DisconnectPacket) Encode()  {
	pk.PutBool(pk.HideDisconnectionScreen)
	pk.PutString(pk.Message)
}

func (pk DisconnectPacket) Decode()  {
	pk.HideDisconnectionScreen = pk.GetBool()
	pk.Message = pk.GetString()
}