package proto

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/info"
)

type Protocol struct {
	protocolNumber int32
	idList         info.PacketIdList
	packets        map[int]func() interfaces.IPacket
	handlers       map[int][][]interfaces.IPacketHandler
}

/**
 * Returns a new Protocol with the given protocol number and packets.
 */
func NewProtocol(protocolNumber int32, idList info.PacketIdList, packets map[int]func() interfaces.IPacket, handlers map[int][][]interfaces.IPacketHandler) *Protocol {
	return &Protocol{protocolNumber, idList, packets, handlers}
}

/**
 * Returns the packet name => Id list of this protocol.
 */
func (protocol *Protocol) GetIdList() info.PacketIdList {
	return protocol.idList
}

/**
 * Returns the protocol number of this protocol.
 */
func (protocol *Protocol) GetProtocolNumber() int32 {
	return protocol.protocolNumber
}

/**
 * Returns all handlers registered for the given packet name.
 */
func (protocol *Protocol) GetHandlers(packet info.PacketName) [][]interfaces.IPacketHandler {
	var id = protocol.idList[packet]
	return protocol.handlers[id]
}

/**
 * Returns all handlers registered on the given ID.
 */
func (protocol *Protocol) GetHandlersById(id int) [][]interfaces.IPacketHandler {
	return protocol.handlers[id]
}

/**
 * Registers a new packet handler to listen for packets with the given ID.
 * Returns a bool indicating success.
 */
func (protocol *Protocol) RegisterHandler(packet info.PacketName, handler interfaces.IPacketHandler, priority int) bool {
	if !handler.SetPriority(priority) {
		return false
	}
	var id = protocol.idList[packet]
	if protocol.handlers[id] == nil {
		protocol.handlers[id] = make([][]interfaces.IPacketHandler, 11)
	}
	protocol.handlers[id][priority] = append(protocol.handlers[id][priority], handler)
	return true
}

/**
 * Deletes all packet handlers listening for packets with the given ID, on the given priority.
 */
func (protocol *Protocol) DeregisterPacketHandlers(packet info.PacketName, priority int) {
	var id = protocol.idList[packet]
	protocol.handlers[id][priority] = []interfaces.IPacketHandler{}
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
