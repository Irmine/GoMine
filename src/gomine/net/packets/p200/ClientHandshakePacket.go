package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type ClientHandshakePacket struct {
	*packets.Packet
}

func NewClientHandshakePacket() *ClientHandshakePacket {
	return &ClientHandshakePacket{packets.NewPacket(info.PacketIds200[info.ClientHandshakePacket])}
}

func (pk *ClientHandshakePacket) Encode()  {

}

func (pk *ClientHandshakePacket) Decode()  {

}