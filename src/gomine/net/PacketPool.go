package net

import (
	"gomine/net/info"
	"gomine/net/packets"
	"gomine/interfaces"
)

var registeredPackets = map[int]func() interfaces.IPacket{}

func init() {
	RegisterPacket(info.LoginPacket, func() interfaces.IPacket { return packets.NewLoginPacket() })
	RegisterPacket(info.ClientHandshakePacket, func() interfaces.IPacket { return packets.NewClientHandshakePacket() })
	RegisterPacket(info.ResourcePackClientResponsePacket, func() interfaces.IPacket { return packets.NewResourcePackClientResponsePacket() })
	RegisterPacket(info.RequestChunkRadiusPacket, func() interfaces.IPacket { return packets.NewRequestChunkRadiusPacket() })
	RegisterPacket(info.MovePlayerPacket, func() interfaces.IPacket { return packets.NewMovePlayerPacket() })
	RegisterPacket(info.CommandRequestPacket, func() interfaces.IPacket { return packets.NewCommandRequestPacket() })
	RegisterPacket(info.ResourcePackChunkRequestPacket, func() interfaces.IPacket { return packets.NewResourcePackChunkRequestPacket() })
	RegisterPacket(info.TextPacket, func() interfaces.IPacket { return packets.NewTextPacket() })
	RegisterPacket(info.PlayerSkinPacket, func() interfaces.IPacket { return packets.NewPlayerListPacket() })
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

/**
 * Deletes a registered packet with the given ID.
 */
func DeregisterPacket(id int) {
	delete(registeredPackets, id)
}