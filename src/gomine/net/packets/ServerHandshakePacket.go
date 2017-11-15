package packets

import "gomine/net"

type ServerHandshakePacket struct {
	*Packet
}

func NewServerHandshakePacket() ServerHandshakePacket {
	return ServerHandshakePacket{NewPacket(net.ServerHandshake)}
}

func (pk *ServerHandshakePacket) Encode()  {

}

func (pk *ServerHandshakePacket) Decode()  {

}