package players

import (
	"gomine/interfaces"
	"goraklib/server"
)

type IPacketHandler interface {
	GetId() int
	Handle(interfaces.IPacket, interfaces.IPlayer, *server.Session, interfaces.IServer) bool
}

/**
 * Packet handlers can be registered to listen on certain packet IDs.
 * Handlers can be registered on unhandled packets in order to handle them from a plugin.
 */
type PacketHandler struct {
	id int
}

/**
 * Returns a new packet handler with the given ID.
 */
func NewPacketHandler(id int) *PacketHandler {
	return &PacketHandler{id}
}

/**
* Returns the ID the handler listens on.
*/
func (handler *PacketHandler) GetId() int {
	return handler.id
}
