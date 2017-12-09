package handlers

/**
 * Packet handlers can be registered to listen on certain packet IDs.
 * Handlers can be registered on unhandled packets in order to handle them from a plugin.
 */
type PacketHandler struct {
	id int
	priority int
}

/**
 * Returns a new packet handler with the given ID.
 */
func NewPacketHandler(id int) *PacketHandler {
	return &PacketHandler{id, 0}
}

/**
* Returns the ID the handler listens on.
*/
func (handler *PacketHandler) GetId() int {
	return handler.id
}

/**
 * Sets the priority of this handler in an integer 0 - 10.
 * 0 is executed first, 10 is executed last.
 */
func (handler *PacketHandler) SetPriority(priority int) bool {
	if priority > 10 || priority < 0 {
		return false
	}
	handler.priority = priority
	return true
}

/**
 * Returns the priority of this handler in an integer 0 - 10.
 */
func (handler *PacketHandler) GetPriority() int {
	return handler.priority
}