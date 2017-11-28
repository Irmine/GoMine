package net

import (
	"gomine/net/info"
	"gomine/net/packets"
	"gomine/interfaces"
)

var registeredPackets = map[int]interfaces.IPacket{}

func InitPacketPool() {
	RegisterPacket(info.LoginPacket, packets.NewLoginPacket())
	RegisterPacket(info.PlayStatusPacket, packets.NewPlayStatusPacket())
	RegisterPacket(info.ClientHandshakePacket, packets.NewClientHandshakePacket())
	RegisterPacket(info.ServerHandshakePacket, packets.NewServerHandshakePacket())
	RegisterPacket(info.ResourcePackInfoPacket, packets.NewResourcePackInfoPacket())
	RegisterPacket(info.ResourcePackClientResponsePacket, packets.NewResourcePackClientResponsePacket())
	RegisterPacket(info.StartGamePacket, packets.NewStartGamePacket())
	RegisterPacket(info.RequestChunkRadiusPacket, packets.NewChunkRadiusRequestPacket())
	RegisterPacket(info.ChunkRadiusUpdatedPacket, packets.NewChunkRadiusUpdatedPacket())
}

func RegisterPacket(id int, packet interfaces.IPacket) {
	registeredPackets[id] = packet
}

func GetPacket(id int) interfaces.IPacket {
	var packet = registeredPackets[id]

	return packet
}