package packets

import "gomine/net"

type ClientHandshakePacket struct {
	DataPacket
	NetId byte
}

func Packet() DataPacket {
	pk := ClientHandshakePacket{}
	pk.NetId = net.ClientHandshake
	return pk
}

func (pk *ClientHandshakePacket) Encode()  {

}

func (pk *ClientHandshakePacket) Decode()  {

}