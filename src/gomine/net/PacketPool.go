package net

import (
	"gomine/net/info"
	"gomine/net/packets"
	"gomine/interfaces"
)

var registeredPackets = map[int]func() interfaces.IPacket{}

func InitPacketPool() {
	RegisterPacket(info.LoginPacket, func() interfaces.IPacket { return packets.NewLoginPacket()})
	RegisterPacket(info.PlayStatusPacket, func() interfaces.IPacket { return packets.NewPlayStatusPacket()})
	RegisterPacket(info.ClientHandshakePacket, func() interfaces.IPacket { return packets.NewClientHandshakePacket()})
	RegisterPacket(info.ServerHandshakePacket, func() interfaces.IPacket { return packets.NewServerHandshakePacket()})
	RegisterPacket(info.ResourcePackInfoPacket, func() interfaces.IPacket { return packets.NewResourcePackInfoPacket()})
	RegisterPacket(info.ResourcePackClientResponsePacket, func() interfaces.IPacket { return packets.NewResourcePackClientResponsePacket()})
	RegisterPacket(info.StartGamePacket, func() interfaces.IPacket { return packets.NewStartGamePacket()})
	RegisterPacket(info.RequestChunkRadiusPacket, func() interfaces.IPacket { return packets.NewChunkRadiusRequestPacket()})
	RegisterPacket(info.ChunkRadiusUpdatedPacket, func() interfaces.IPacket { return packets.NewChunkRadiusUpdatedPacket()})
}

func RegisterPacket(id int, function func() interfaces.IPacket) {
	registeredPackets[id] = function
}

func GetPacket(id int) interfaces.IPacket {
	return registeredPackets[id]()
}