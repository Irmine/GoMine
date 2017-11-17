package packets

import "gomine/net"

type ClientHandshakePacket struct {
	*Packet
}

func NewClientHandshakePacket() ClientHandshakePacket {
	return ClientHandshakePacket{NewPacket(net.ClientHandshakePacket)}
}

func (pk *ClientHandshakePacket) Encode()  {

}

func (pk *ClientHandshakePacket) Decode()  {

}