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

var registeredHandlers = map[int]map[int][]interfaces.IPacketHandler{}

func init() {
	RegisterPacketHandler(info.LoginPacket, handlers.NewLoginHandler(), PriorityLast)
	RegisterPacketHandler(info.RequestChunkRadiusPacket, handlers.NewChunkRadiusRequestHandler(), PriorityLast)
	RegisterPacketHandler(info.ResourcePackClientResponsePacket, handlers.NewResourcePackClientResponseHandler(), PriorityLast)
	//RegisterPacketHandler(info.MovePlayerPacket, handlers.NewMovePlayerHandler(), PriorityLast)
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
		registeredHandlers[id] = make(map[int][]interfaces.IPacketHandler)
	}
	registeredHandlers[id][priority] = append(registeredHandlers[id][priority], handler)
	return true
}

/**
 * Returns all packet handlers registered on the given ID.
 */
func GetPacketHandlers(id int) map[int][]interfaces.IPacketHandler {
	return registeredHandlers[id]
}