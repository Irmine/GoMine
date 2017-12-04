package net

import (
	"gomine/net/info"
	"gomine/players/handlers"
	"gomine/interfaces"
)

var registeredHandlers = map[int][]interfaces.IPacketHandler{}

func InitHandlerPool() {
	RegisterPacketHandler(info.LoginPacket, handlers.NewLoginHandler())
	RegisterPacketHandler(info.RequestChunkRadiusPacket, handlers.NewChunkRadiusRequestHandler())
	RegisterPacketHandler(info.ResourcePackClientResponsePacket, handlers.NewResourcePackClientResponseHandler())
}

/**
 * Registers a new packet handler to listen for packets with the given ID.
 */
func RegisterPacketHandler(id int, handler interfaces.IPacketHandler) {
	registeredHandlers[id] = append(registeredHandlers[id], handler)
}

/**
 * Returns all packet handlers registered on the given ID.
 */
func GetPacketHandlers(id int) []interfaces.IPacketHandler {
	return registeredHandlers[id]
}