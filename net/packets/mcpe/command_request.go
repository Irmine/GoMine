package mcpe

import (
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type CommandRequestPacket struct {
	*packets.Packet
	CommandText string
	Type        uint32
	UUID        uuid.UUID
	RequestId   string
	Internal    bool
}

func NewCommandRequestPacket() *CommandRequestPacket {
	return &CommandRequestPacket{packets.NewPacket(info.PacketIds[info.CommandRequestPacket]), "", 0, uuid.New(), "", false}
}

func (pk *CommandRequestPacket) Encode() {

}

func (pk *CommandRequestPacket) Decode() {
	pk.CommandText = pk.GetString()
	pk.Type = pk.GetUnsignedVarInt()
	pk.UUID = pk.GetUUID()
	pk.RequestId = pk.GetString()
	pk.Internal = pk.GetBool()
}
