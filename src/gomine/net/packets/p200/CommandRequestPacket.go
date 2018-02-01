package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type CommandRequestPacket struct {
	*packets.Packet
	CommandText string
	Type uint32
	RequestId string
	Internal bool
}

func NewCommandRequestPacket() *CommandRequestPacket {
	return &CommandRequestPacket{packets.NewPacket(info.CommandRequestPacket), "", 0, "", false}
}

func (pk *CommandRequestPacket) Encode() {

}

func (pk *CommandRequestPacket) Decode() {
	pk.CommandText = pk.GetString()

	pk.Type = pk.GetUnsignedVarInt()

	// UUID. TODO: Implement properly.
	pk.GetLittleInt()
	pk.GetLittleInt()
	pk.GetLittleInt()
	pk.GetLittleInt()

	pk.RequestId = pk.GetString()

	pk.Internal = pk.GetBool()
}