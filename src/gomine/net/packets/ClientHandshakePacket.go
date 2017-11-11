package packets

type ClientHandshakePacket struct {
	DataPacket
	NetId byte
}

func Packet() DataPacket {
	pk := ClientHandshakePacket{
		NetId: 0x04,
	}
	return pk
}

func (pk *ClientHandshakePacket) Encode()  {

}

func (pk *ClientHandshakePacket) Decode()  {

}