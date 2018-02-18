package handlers

// Packet handlers can be registered to listen on certain packet IDs.
// Handlers can be registered on unhandled packets in order to handle them from a plugin.
type PacketHandler struct {
	priority int
}

// NewPacketHandler returns a new packet handler with the given ID.
func NewPacketHandler() *PacketHandler {
	return &PacketHandler{0}
}

// SetPriority sets the priority of this handler in an integer 0 - 10.
// 0 is executed first, 10 is executed last.
func (handler *PacketHandler) SetPriority(priority int) bool {
	if priority > 10 || priority < 0 {
		return false
	}
	handler.priority = priority
	return true
}

// GetPriority returns the priority of this handler in an integer 0 - 10.
func (handler *PacketHandler) GetPriority() int {
	return handler.priority
}
