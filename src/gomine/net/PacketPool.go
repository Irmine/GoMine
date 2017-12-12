package net

import (
	"gomine/net/info"
	"gomine/net/packets"
	"gomine/interfaces"
)

var registeredPackets = map[int]func() interfaces.IPacket{}

func init() {
	RegisterPacket(info.LoginPacket, func() interfaces.IPacket { return packets.NewLoginPacket() })
	RegisterPacket(info.PlayStatusPacket, func() interfaces.IPacket { return packets.NewPlayStatusPacket() })
	RegisterPacket(info.ClientHandshakePacket, func() interfaces.IPacket { return packets.NewClientHandshakePacket() })
	RegisterPacket(info.ServerHandshakePacket, func() interfaces.IPacket { return packets.NewServerHandshakePacket() })
	RegisterPacket(info.ResourcePackInfoPacket, func() interfaces.IPacket { return packets.NewResourcePackInfoPacket() })
	RegisterPacket(info.ResourcePackClientResponsePacket, func() interfaces.IPacket { return packets.NewResourcePackClientResponsePacket() })
	RegisterPacket(info.StartGamePacket, func() interfaces.IPacket { return packets.NewStartGamePacket() })
	RegisterPacket(info.RequestChunkRadiusPacket, func() interfaces.IPacket { return packets.NewChunkRadiusRequestPacket() })
	RegisterPacket(info.ChunkRadiusUpdatedPacket, func() interfaces.IPacket { return packets.NewChunkRadiusUpdatedPacket() })
	RegisterPacket(info.FullChunkDataPacket, func() interfaces.IPacket { return packets.NewFullChunkPacket() })
	RegisterPacket(info.CraftingDataPacket, func() interfaces.IPacket { return packets.NewCraftingDataPacket() })
	RegisterPacket(info.ResourcePackStackPacket, func() interfaces.IPacket { return packets.NewResourcePackStackPacket() })
	RegisterPacket(info.MovePlayerPacket, func() interfaces.IPacket { return packets.NewMovePlayerPacket() })
	RegisterPacket(info.CommandRequestPacket, func() interfaces.IPacket { return packets.NewCommandRequestPacket() })
	RegisterPacket(info.ResourcePackDataInfoPacket, func() interfaces.IPacket { return packets.NewResourcePackDataInfoPacket() })
	RegisterPacket(info.ResourcePackChunkRequestPacket, func() interfaces.IPacket { return packets.NewResourcePackChunkRequestPacket() })
	RegisterPacket(info.ResourcePackChunkDataPacket, func() interfaces.IPacket { return packets.NewResourcePackChunkDataPacket() })
}

/**
 * Returns if a packet with the given ID is registered.
 */
func IsPacketRegistered(id int) bool {
	if _, ok := registeredPackets[id]; ok {
		return true
	}
	return false
}

/**
 * Returns a new packet with the given ID and a function that returns that packet.
 */
func RegisterPacket(id int, function func() interfaces.IPacket) {
	registeredPackets[id] = function
}

/**
 * Returns a new packet with the given ID.
 */
func GetPacket(id int) interfaces.IPacket {
	return registeredPackets[id]()
}