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
	RegisterPacketHandler(info.LoginPacket, handlers.NewLoginHandler(), PriorityLast)
	RegisterPacketHandler(info.ClientHandshakePacket, handlers.NewClientHandshakeHandler(), PriorityLast)
	RegisterPacketHandler(info.RequestChunkRadiusPacket, handlers.NewChunkRadiusRequestHandler(), PriorityLast)
	RegisterPacketHandler(info.ResourcePackClientResponsePacket, handlers.NewResourcePackClientResponseHandler(), PriorityLast)
	RegisterPacketHandler(info.MovePlayerPacket, handlers.NewMovePlayerHandler(), PriorityLast)
	RegisterPacketHandler(info.CommandRequestPacket, handlers.NewCommandRequestHandler(), PriorityLast)
	RegisterPacketHandler(info.ResourcePackChunkRequestPacket, handlers.NewResourcePackChunkRequestHandler(), PriorityLast)
	RegisterPacketHandler(info.TextPacket, handlers.NewTextHandler(), PriorityLast)
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