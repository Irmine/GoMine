package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type ServerHandshakePacket struct {
	*packets.Packet
	Jwt string
}

func NewServerHandshakePacket() *ServerHandshakePacket {
	return &ServerHandshakePacket{packets.NewPacket(info.PacketIds[info.ServerHandshakePacket]), ""}
}

func (pk *ServerHandshakePacket) Encode() {
	pk.PutString(pk.Jwt)
}

func (pk *ServerHandshakePacket) Decode() {

}
