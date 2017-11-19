package packets

import (
	"gomine/net/info"
)

var packets = map[int]IPacket{}

func InitPacketPool() {
	RegisterPacket(info.LoginPacket, NewLoginPacket())
	RegisterPacket(info.PlayStatusPacket, NewPlayStatusPacket())
	RegisterPacket(info.ClientHandshakePacket, NewClientHandshakePacket())
	RegisterPacket(info.ServerHandshakePacket, NewServerHandshakePacket())
}

func RegisterPacket(id int, packet IPacket) {
	packets[id] = packet
}

func GetPacket(id int) IPacket {
	return packets[id]
}