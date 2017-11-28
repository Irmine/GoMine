package packets

import (
	"gomine/net/info"
)

type ServerHandshakePacket struct {
	*Packet
	Jwt string
}

func NewServerHandshakePacket() *ServerHandshakePacket {
	return &ServerHandshakePacket{NewPacket(info.ServerHandshakePacket), ""}
}

func (pk *ServerHandshakePacket) Encode()  {
	pk.PutString(pk.Jwt)
}

func (pk *ServerHandshakePacket) Decode()  {
	pk.Jwt = pk.GetString()
}