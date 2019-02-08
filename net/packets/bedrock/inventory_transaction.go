package bedrock

import (
	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/items"
	"github.com/irmine/gomine/items/inventory/io"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/worlds/blocks"
)

// Transaction Types
const (
	Normal = iota + 0
	Mismatch
	UseItem
	UseItemOnEntity
	ReleaseItem
)

// Action Types
const (
	ItemClickBlock = iota + 0
	ItemClickAir
	ItemBreakBlock

	//CONSUMABLE ITEMS
	ItemRelease = iota + 0
	ItemConsume
)

// Entity Action types
const (
	ItemOnEntityInteract = iota + 0
	ItemOnEntityAttack
)

type InventoryTransactionPacket struct {
	*packets.Packet
	ActionList *io.InventoryActionIOList
	TransactionType, ActionType uint32
	Face, HotbarSlot int32
	ItemSlot *items.Stack
	BlockPosition blocks.Position
	PlayerPosition, ClickPosition, HeadPosition r3.Vector
	RuntimeId uint64
}

func NewInventoryTransactionPacket() *InventoryTransactionPacket {
	pk := &InventoryTransactionPacket{Packet: packets.NewPacket(info.PacketIds[info.InventoryTransactionPacket]),
		ActionList: io.NewInventoryActionIOList(),
		TransactionType: 0,
		ActionType: 0,
		Face: 0,
		HotbarSlot: 0,
		ItemSlot: &items.Stack{},
		PlayerPosition: r3.Vector{},
		ClickPosition: r3.Vector{},
		HeadPosition: r3.Vector{},
		RuntimeId: 0,
	}
	return pk
}

func (pk *InventoryTransactionPacket) Encode()  {
	pk.PutUnsignedVarInt(pk.TransactionType)
	pk.ActionList.WriteToBuffer(pk.MinecraftStream)

	switch pk.TransactionType {
	case Normal, Mismatch:
		break
	case UseItem:
		pk.PutUnsignedVarInt(pk.ActionType)
		pk.PutBlockPosition(pk.BlockPosition)
		pk.PutVarInt(pk.Face)
		pk.PutVarInt(pk.HotbarSlot)
		pk.PutItem(pk.ItemSlot)
		pk.PutVector(pk.PlayerPosition)
		pk.PutVector(pk.ClickPosition)
		break
	case UseItemOnEntity:
		pk.PutUnsignedVarLong(pk.RuntimeId)
		pk.PutUnsignedVarInt(pk.ActionType)
		pk.PutVarInt(pk.HotbarSlot)
		pk.PutItem(pk.ItemSlot)
		pk.PutVector(pk.PlayerPosition)
		pk.PutVector(pk.ClickPosition)
		break
	case ReleaseItem:
		pk.PutUnsignedVarInt(pk.ActionType)
		pk.PutVarInt(pk.HotbarSlot)
		pk.PutItem(pk.ItemSlot)
		pk.PutVector(pk.HeadPosition)
		break
	default:
		panic("Unknown transaction type passed: " + string(pk.TransactionType))
	}
}

func (pk *InventoryTransactionPacket) Decode() {
	pk.TransactionType = pk.GetUnsignedVarInt()
	pk.ActionList.ReadFromBuffer(pk.MinecraftStream)

	switch pk.TransactionType{
	case Normal, Mismatch:
		break
	case UseItem:
		pk.ActionType = pk.GetUnsignedVarInt()
		pk.BlockPosition = pk.GetBlockPosition()
		pk.Face = pk.GetVarInt()
		pk.HotbarSlot = pk.GetVarInt()
		pk.ItemSlot = pk.GetItem()
		pk.PlayerPosition = pk.GetVector()
		pk.ClickPosition = pk.GetVector()
		break
	case UseItemOnEntity:
		pk.RuntimeId = pk.GetUnsignedVarLong()
		pk.ActionType = pk.GetUnsignedVarInt()
		pk.HotbarSlot = pk.GetVarInt()
		pk.ItemSlot = pk.GetItem()
		pk.PlayerPosition = pk.GetVector()
		pk.ClickPosition = pk.GetVector()
		break
	case ReleaseItem:
		pk.ActionType = pk.GetUnsignedVarInt()
		pk.HotbarSlot = pk.GetVarInt()
		pk.ItemSlot = pk.GetItem()
		pk.HeadPosition = pk.GetVector()
		break
	default:
		panic("Error: Unknown transaction type received: " + string(pk.TransactionType))
	}
}
