package io

import (
	"github.com/irmine/gomine/items"
	"github.com/irmine/gomine/net/packets"
)

const (
	ContainerSource = iota + 0
	WorldSource = 2
	//CreativeSource = 3
)

type InventoryActionIO struct {
	Source uint32
	WindowId int32
	SourceFlags uint32
	InventorySlot uint32
	OldItem *items.Stack
	NewItem *items.Stack
}

func NewInventoryActionIO() InventoryActionIO{
	return InventoryActionIO{}
}

func (IO *InventoryActionIO) WriteToBuffer(bs *packets.MinecraftStream) {
	bs.PutUnsignedVarInt(IO.Source)

	switch IO.Source {
	case ContainerSource:
		bs.PutVarInt(IO.WindowId)
		break
	case WorldSource:
		bs.PutUnsignedVarInt(IO.SourceFlags)
		break
	}

	bs.PutUnsignedVarInt(IO.InventorySlot)
	bs.PutItem(IO.OldItem)
	bs.PutItem(IO.NewItem)
}

func (IO *InventoryActionIO) ReadFromBuffer(bs *packets.MinecraftStream) InventoryActionIO {
	IO.Source = bs.GetUnsignedVarInt()

	switch IO.Source {
	case ContainerSource:
		IO.WindowId = bs.GetVarInt()
		break
	case WorldSource:
		IO.SourceFlags = bs.GetUnsignedVarInt()
		break
	}

	IO.InventorySlot = bs.GetUnsignedVarInt()
	IO.OldItem = bs.GetItem()
	IO.NewItem = bs.GetItem()

	return *IO
}