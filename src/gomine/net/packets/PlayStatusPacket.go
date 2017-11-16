package packets

import "gomine/net"

type PlayStatusPacket struct {
	*Packet
	status int32
	protocol int
}

func NewPlayStatusPacket() PlayStatusPacket {
	return PlayStatusPacket{NewPacket(net.PlayStatus), 0, net.LatestProtocol}
}

func (pk *PlayStatusPacket) EncodeHeader() {
	if pk.protocol < 130 {
		pk.PutByte(byte(pk.packetId))
	} else {
		pk.EncodeHeader()
	}
}

func (pk *PlayStatusPacket) Encode()  {
	pk.PutInt(pk.status)
}

func (pk *PlayStatusPacket) Decode()  {
	pk.status = pk.GetInt()
}