package packets

import "gomine/utils"

type ServerHandshakePacket struct {
	DataPacket
	Buffer []byte
	Offset int
	NetId  byte
	Jtw    string
}

func NewServerHandshakePacket() DataPacket {
	pk := ServerHandshakePacket{
		NetId: 0x04,
	}
	pk.Offset = len(pk.Buffer)
	return pk
}

func (pk *ServerHandshakePacket) Encode()  {
	utils.WriteString(&pk.Buffer, pk.Jtw)
}

func (pk *ServerHandshakePacket) Decode()  {
	pk.Jtw = utils.ReadString(&pk.Buffer, &pk.Offset)
}