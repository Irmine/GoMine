package protocol

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/packs"
	"github.com/irmine/worlds/blocks"
	"github.com/irmine/worlds/chunks"
	"github.com/irmine/worlds/entities/data"
)

type IPacketManager interface {
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
	GetAddPlayer(uuid.UUID, AddPlayerEntry) packets.IPacket
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
	GetResourcePackInfo(bool, *packs.Stack, *packs.Stack) packets.IPacket
	GetResourcePackStack(bool, *packs.Stack, *packs.Stack) packets.IPacket
	GetServerHandshake(string) packets.IPacket
	GetSetEntityData(uint64, map[uint32][]interface{}) packets.IPacket
	GetStartGame(StartGameEntry, []byte) packets.IPacket
	GetText(types.Text) packets.IPacket
	GetTransfer(string, uint16) packets.IPacket
	GetUpdateAttributes(uint64, data.AttributeMap) packets.IPacket
	GetNetworkChunkPublisherUpdatePacket(position blocks.Position, radius uint32) packets.IPacket
	GetMoveEntity(uint64, r3.Vector, data.Rotation, byte, bool) packets.IPacket
	GetPlayerSkin(uuid2 uuid.UUID, skinId, geometryName, geometryData string, skinData, capeData []byte) packets.IPacket
	GetPlayerAction(runtimeId uint64, action int32, position blocks.Position, face int32) packets.IPacket
	GetAnimate(action int32, runtimeId uint64, float float32) packets.IPacket
	GetUpdateBlock(position blocks.Position, blockRuntimeId, dataLayerId uint32) packets.IPacket
}

// PacketManagerBase is a struct providing the base for a PacketManagerBase.
// It provides utility functions for a basic PacketManagerBase implementation.
type PacketManagerBase struct {
	idList         info.PacketIdList
	packets        map[int]func() packets.IPacket
	handlers       map[int][][]Handler
}

// NewBase returns a new PacketManagerBase with the given PacketManagerBase number and packets.
func NewPacketManagerBase(idList info.PacketIdList, packets map[int]func() packets.IPacket, handlers map[int][][]Handler) *PacketManagerBase {
	return &PacketManagerBase{idList, packets, handlers}
}

// GetIdList returns the packet name => Id list of the bedrock.
func (Base *PacketManagerBase) GetIdList() info.PacketIdList {
	return Base.idList
}

// GetHandlers returns all handlers registered for the given packet name.
func (Base *PacketManagerBase) GetHandlers(packet info.PacketName) [][]Handler {
	var id = Base.idList[packet]
	return Base.handlers[id]
}

// GetHandlersById returns all handlers registered on the given ID.
func (Base *PacketManagerBase) GetHandlersById(id int) [][]Handler {
	return Base.handlers[id]
}

// RegisterHandler registers a new packet handler to listen for packets with the given ID.
// This function uses the priority of the handler.
// Returns a bool indicating success.
func (Base *PacketManagerBase) RegisterHandler(packet info.PacketName, handler Handler) bool {
	var id = Base.idList[packet]
	if Base.handlers[id] == nil {
		Base.handlers[id] = make([][]Handler, 11)
	}
	Base.handlers[id][handler.GetPriority()] = append(Base.handlers[id][handler.GetPriority()], handler)
	return true
}

// DeregisterPackHandlers deregisters all packet handlers listening for packets with the given ID, on the given priority.
func (Base *PacketManagerBase) DeregisterPacketHandlers(packet info.PacketName, priority int) {
	var id = Base.idList[packet]
	Base.handlers[id][priority] = []Handler{}
}

// GetPackets returns a packet ID => packet function map containing all registered packets.
func (Base *PacketManagerBase) GetPackets() map[int]func() packets.IPacket {
	return Base.packets
}

// RegisterPacket registers a packet function with the given packet ID.
func (Base *PacketManagerBase) RegisterPacket(packetId int, packetFunc func() packets.IPacket) {
	Base.packets[packetId] = packetFunc
}

// GetPacket returns a packet with the given packet ID.
func (Base *PacketManagerBase) GetPacket(packetId int) packets.IPacket {
	return Base.packets[packetId]()
}

// IsPacketRegistered checks if the PacketManagerBase has a packet with the given packet ID.
func (Base *PacketManagerBase) IsPacketRegistered(packetId int) bool {
	var _, ok = Base.packets[packetId]
	return ok
}
