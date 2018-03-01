package protocol

import (
	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/packs"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/worlds/chunks"
	"github.com/irmine/worlds/entities/data"
)

// Protocol is an interface satisfied by every protocol implementation.
// A protocol provides basic packet wrapping functions.
type Protocol interface {
	GetProtocolNumber() int32
	GetIdList() info.PacketIdList
	GetHandlers(packet info.PacketName) [][]Handler
	GetHandlersById(id int) [][]Handler
	RegisterHandler(packet info.PacketName, handler Handler) bool
	DeregisterPacketHandlers(packet info.PacketName, priority int)
	GetPackets() map[int]func() packets.IPacket
	RegisterPacket(packetId int, packetFunc func() packets.IPacket)
	GetPacket(packetId int) packets.IPacket
	IsPacketRegistered(packetId int) bool

	GetAddEntity(AddEntityEntry) packets.IPacket
	GetAddPlayer(utils.UUID, int32, AddPlayerEntry) packets.IPacket
	GetChunkRadiusUpdated(int32) packets.IPacket
	GetCraftingData() packets.IPacket
	GetDisconnect(string, bool) packets.IPacket
	GetFullChunkData(*chunks.Chunk) packets.IPacket
	GetMovePlayer(uint64, r3.Vector, data.Rotation, byte, bool, uint64) packets.IPacket
	GetPlayerList(byte, map[string]PlayerListEntry) packets.IPacket
	GetPlayStatus(int32) packets.IPacket
	GetRemoveEntity(int64) packets.IPacket
	GetResourcePackChunkData(string, int32, int64, []byte) packets.IPacket
	GetResourcePackDataInfo(packs.Pack) packets.IPacket
	GetResourcePackInfo(bool, []packs.Pack, []packs.Pack) packets.IPacket
	GetResourcePackStack(bool, []packs.Pack, []packs.Pack) packets.IPacket
	GetServerHandshake(string) packets.IPacket
	GetSetEntityData(uint64, map[uint32][]interface{}) packets.IPacket
	GetStartGame(StartGameEntry) packets.IPacket
	GetText(types.Text) packets.IPacket
	GetTransfer(string, uint16) packets.IPacket
	GetUpdateAttributes(uint64, data.AttributeMap) packets.IPacket
}

// Base is a struct providing the base for a Base.
// It provides utility functions for a basic Base implementation.
type Base struct {
	ProtocolNumber int32
	idList         info.PacketIdList
	packets        map[int]func() packets.IPacket
	handlers       map[int][][]Handler
}

// NewBase returns a new Base with the given Base number and packets.
func NewBase(BaseNumber int32, idList info.PacketIdList, packets map[int]func() packets.IPacket, handlers map[int][][]Handler) *Base {
	return &Base{BaseNumber, idList, packets, handlers}
}

// GetIdList returns the packet name => Id list of the protocol.
func (Base *Base) GetIdList() info.PacketIdList {
	return Base.idList
}

// GetProtocolNumber returns the Base number of the protocol.
func (Base *Base) GetProtocolNumber() int32 {
	return Base.ProtocolNumber
}

// GetHandlers returns all handlers registered for the given packet name.
func (Base *Base) GetHandlers(packet info.PacketName) [][]Handler {
	var id = Base.idList[packet]
	return Base.handlers[id]
}

// GetHandlersById returns all handlers registered on the given ID.
func (Base *Base) GetHandlersById(id int) [][]Handler {
	return Base.handlers[id]
}

// RegisterHandler registers a new packet handler to listen for packets with the given ID.
// This function uses the priority of the handler.
// Returns a bool indicating success.
func (Base *Base) RegisterHandler(packet info.PacketName, handler Handler) bool {
	var id = Base.idList[packet]
	if Base.handlers[id] == nil {
		Base.handlers[id] = make([][]Handler, 11)
	}
	Base.handlers[id][handler.GetPriority()] = append(Base.handlers[id][handler.GetPriority()], handler)
	return true
}

// DeregisterPackHandlers deregisters all packet handlers listening for packets with the given ID, on the given priority.
func (Base *Base) DeregisterPacketHandlers(packet info.PacketName, priority int) {
	var id = Base.idList[packet]
	Base.handlers[id][priority] = []Handler{}
}

// GetPackets returns a packet ID => packet function map containing all registered packets.
func (Base *Base) GetPackets() map[int]func() packets.IPacket {
	return Base.packets
}

// RegisterPacket registers a packet function with the given packet ID.
func (Base *Base) RegisterPacket(packetId int, packetFunc func() packets.IPacket) {
	Base.packets[packetId] = packetFunc
}

// GetPacket returns a packet with the given packet ID.
func (Base *Base) GetPacket(packetId int) packets.IPacket {
	return Base.packets[packetId]()
}

// IsPacketRegistered checks if the Base has a packet with the given packet ID.
func (Base *Base) IsPacketRegistered(packetId int) bool {
	var _, ok = Base.packets[packetId]
	return ok
}
