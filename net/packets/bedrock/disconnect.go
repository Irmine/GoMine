package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type DisconnectPacket struct {
	*packets.Packet
	HideDisconnectionScreen bool
	Message                 string
}

func NewDisconnectPacket() *DisconnectPacket {
	return &DisconnectPacket{packets.NewPacket(info.PacketIds[info.DisconnectPacket]), true, ""}
}

func (pk *DisconnectPacket) Encode() {
	pk.PutBool(pk.HideDisconnectionScreen)
	pk.PutString(pk.Message)
}

func (pk *DisconnectPacket) Decode() {
	pk.HideDisconnectionScreen = pk.GetBool()
	pk.Message = pk.GetString()
}
