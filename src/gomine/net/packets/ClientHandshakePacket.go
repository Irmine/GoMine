package packets

import "gomine/net"

type ClientHandshakePacket struct {
	*Packet
}

func NewClientHandshakePacket() ClientHandshakePacket {
	return ClientHandshakePacket{NewPacket(net.ClientHandshake)}
}

func (pk *ClientHandshakePacket) Encode()  {

}

func (pk *ClientHandshakePacket) Decode()  {

}