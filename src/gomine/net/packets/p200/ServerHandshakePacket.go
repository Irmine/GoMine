package p200

import (
	"gomine/net/info"
	"gomine/net/packets"
)

type ServerHandshakePacket struct {
	*packets.Packet
	Jwt string
}

func NewServerHandshakePacket() *ServerHandshakePacket {
	return &ServerHandshakePacket{packets.NewPacket(info.ServerHandshakePacket), ""}
}

func (pk *ServerHandshakePacket) Encode()  {
	pk.PutString(pk.Jwt)
}

func (pk *ServerHandshakePacket) Decode()  {

}