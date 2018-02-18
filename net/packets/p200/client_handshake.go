package p200

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type ClientHandshakePacket struct {
	*packets.Packet
}

func NewClientHandshakePacket() *ClientHandshakePacket {
	return &ClientHandshakePacket{packets.NewPacket(info.PacketIds200[info.ClientHandshakePacket])}
}

func (pk *ClientHandshakePacket) Encode() {

}

func (pk *ClientHandshakePacket) Decode() {

}
