package net

import (
	"gomine/players"
	"gomine/net/info"
	"gomine/players/handlers"
)

var registeredHandlers = map[int][]players.IPacketHandler{}

func InitHandlerPool() {
	RegisterPacketHandler(info.LoginPacket,					handlers.NewLoginHandler())
	RegisterPacketHandler(info.RequestChunkRadiusPacket,	handlers.NewChunkRadiusRequestHandler())
}

/**
 * Registers a new packet handler to listen for packets with the given ID.
 */
func RegisterPacketHandler(id int, handler players.IPacketHandler) {
	registeredHandlers[id] = append(registeredHandlers[id], handler)
}

/**
 * Returns all packet handlers registered on the given ID.
 */
func GetPacketHandlers(id int) []players.IPacketHandler {
	return registeredHandlers[id]
}