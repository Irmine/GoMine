package net

import (
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/utils"
)

// Packet handlers can be registered to listen on certain packet IDs.
// Handlers can be registered on unhandled packets in order to handle them from a plugin.
// Every packet handler has a handling function that handles the incoming packet.
type PacketHandler struct {
	function func(packet packets.IPacket, logger *utils.Logger, session *MinecraftSession) bool
	priority int
}

// NewPacketHandler returns a new packet handler with the given ID.
// NewPacketHandler will by default use a priority of 5.
func NewPacketHandler(function func(packet packets.IPacket, logger *utils.Logger, session *MinecraftSession) bool) *PacketHandler {
	return &PacketHandler{function, 5}
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
