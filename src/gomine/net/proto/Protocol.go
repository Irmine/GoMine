package proto

import (
	"gomine/interfaces"
	"gomine/net/packets"
)

type Protocol struct {
	protocolNumber int32
	packets map[int]func() interfaces.IPacket
}

/**
 * Returns a new Protocol with the given protocol number and packets.
 */
func NewProtocol(protocolNumber int32, packets map[int]func() interfaces.IPacket) *Protocol {
	return &Protocol{protocolNumber, packets}
}

/**
 * Returns the protocol number of this protocol.
 */
func (protocol *Protocol) GetProtocolNumber() int32 {
	return protocol.protocolNumber
}

/**
 * Returns a packet ID => packet function map containing all registered packets.
 */
func (protocol *Protocol) GetPackets() map[int]func() interfaces.IPacket {
	return protocol.packets
}

/**
 * Registers a packet function with the given packet ID.
 */
func (protocol *Protocol) RegisterPacket(packetId int, packetFunc func() interfaces.IPacket) {
	protocol.packets[packetId] = packetFunc
}

/**
 * Returns a packet with the given packet ID.
 */
func (protocol *Protocol) GetPacket(packetId int) interfaces.IPacket {
	return protocol.packets[packetId]()
}

/**
 * Checks if the protocol has a packet with the given packet ID.
 */
func (protocol *Protocol) IsPacketRegistered(packetId int) bool {
	var _, ok = protocol.packets[packetId]
	return ok
}