package packets

import "gomine/net"

type LoginPacket struct {
	*Packet
	Username string
	Protocol int
	ClientUUID string
	ClientId int
	Xuid string
	IdentityPublicKey string
	ServerAddress string
	Locale string
}

func NewLoginPacket() LoginPacket {
	pk := LoginPacket{}
	pk.packetId = net.LoginPacket
	return pk
}

func (pk *LoginPacket) Encode()  {
	//todo
}

func (pk *LoginPacket) Decode()  {
	//todo
}