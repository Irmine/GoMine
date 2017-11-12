package packets

import (
	"gomine/utils"
	"gomine/net"
)

type ServerHandshakePacket struct {
	DataPacket
	Buffer []byte
	Offset int
	NetId  byte
	Jwt    string
}

func NewServerHandshakePacket() DataPacket {
	pk := ServerHandshakePacket{}
	pk.Offset = 0
	pk.NetId = net.ServerHandshake
	return pk
}

func (pk *ServerHandshakePacket) Encode()  {
	utils.WriteString(&pk.Buffer, pk.Jwt)
}

func (pk *ServerHandshakePacket) Decode()  {
	pk.Jwt = utils.ReadString(&pk.Buffer, &pk.Offset)
}