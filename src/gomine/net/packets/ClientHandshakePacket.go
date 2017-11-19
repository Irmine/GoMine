package packets

import (
	"gomine/net/info"
)

type ClientHandshakePacket struct {
	*Packet
}

func NewClientHandshakePacket() ClientHandshakePacket {
	return ClientHandshakePacket{NewPacket(info.ClientHandshakePacket)}
}

func (pk ClientHandshakePacket) Encode()  {

}

func (pk ClientHandshakePacket) Decode()  {

}