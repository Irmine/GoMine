package net

import (
	"gomine/net/info"
	"gomine/net/packets"
)

var registeredPackets = map[int]packets.IPacket{}

func InitPacketPool() {
	RegisterPacket(info.LoginPacket, packets.NewLoginPacket())
	RegisterPacket(info.PlayStatusPacket, packets.NewPlayStatusPacket())
	RegisterPacket(info.ClientHandshakePacket, packets.NewClientHandshakePacket())
	RegisterPacket(info.ServerHandshakePacket, packets.NewServerHandshakePacket())
	RegisterPacket(info.ResourcePackInfoPacket, packets.NewResourcePackInfoPacket())
	RegisterPacket(info.ResourcePackClientResponsePacket, packets.NewResourcePackClientResponsePacket())
}

func RegisterPacket(id int, packet packets.IPacket) {
	registeredPackets[id] = packet
}

func GetPacket(id int) packets.IPacket {
	return registeredPackets[id]
}