package packets

import "gomine/net"

type ServerHandshakePacket struct {
	*Packet
	Jwt string
}

func NewServerHandshakePacket() ServerHandshakePacket {
	return ServerHandshakePacket{NewPacket(net.ServerHandshakePacket), ""}
}

func (pk *ServerHandshakePacket) Encode()  {
	pk.PutString(pk.Jwt)
}

func (pk *ServerHandshakePacket) Decode()  {
	pk.Jwt = pk.GetString()
}