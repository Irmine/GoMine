package net

import (
	"gomine/net/info"
	"gomine/players/handlers"
	"gomine/interfaces"
)

const (
	PriorityFirst = 0
	PriorityLast = 10
)

var registeredHandlers = map[int][][]interfaces.IPacketHandler{}

func init() {
	RegisterPacketHandler(info.LoginPacket, handlers.NewLoginHandler(), 8)
	RegisterPacketHandler(info.ClientHandshakePacket, handlers.NewClientHandshakeHandler(), 8)
	RegisterPacketHandler(info.RequestChunkRadiusPacket, handlers.NewRequestChunkRadiusHandler(), 8)
	RegisterPacketHandler(info.ResourcePackClientResponsePacket, handlers.NewResourcePackClientResponseHandler(), 8)
	RegisterPacketHandler(info.MovePlayerPacket, handlers.NewMovePlayerHandler(), 8)
	RegisterPacketHandler(info.CommandRequestPacket, handlers.NewCommandRequestHandler(), 8)
	RegisterPacketHandler(info.ResourcePackChunkRequestPacket, handlers.NewResourcePackChunkRequestHandler(), 8)
	RegisterPacketHandler(info.TextPacket, handlers.NewTextHandler(), 8)
}

/**
 * Registers a new packet handler to listen for packets with the given ID.
 * Returns a bool indicating success.
 */
func RegisterPacketHandler(id int, handler interfaces.IPacketHandler, priority int) bool {
	if !handler.SetPriority(priority) {
		return false
	}
	if registeredHandlers[id] == nil {
		registeredHandlers[id] = make([][]interfaces.IPacketHandler, 11)
	}
	registeredHandlers[id][priority] = append(registeredHandlers[id][priority], handler)
	return true
}

/**
 * Returns all packet handlers registered on the given ID.
 */
func GetPacketHandlers(id int) [][]interfaces.IPacketHandler {
	return registeredHandlers[id]
}

/**
 * Deletes all packet handlers listening for packets with the given ID, on the given priority.
 */
func DeregisterPacketHandlers(id int, priority int) {
	registeredHandlers[id][priority] = []interfaces.IPacketHandler{}
}