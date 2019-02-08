package io

import (
	"github.com/irmine/gomine/net/packets"
)

type InventoryActionIOList struct {
	List []InventoryActionIO
}

func NewInventoryActionIOList() *InventoryActionIOList{
	return &InventoryActionIOList{}
}

func (IOList *InventoryActionIOList) GetCount() int {
	return len(IOList.List)
}

func (IOList *InventoryActionIOList) PutAction(io InventoryActionIO) {
	IOList.List = append(IOList.List, io)
}

func (IOList *InventoryActionIOList) WriteToBuffer(bs *packets.MinecraftStream) {
	c := len(IOList.List)
	bs.PutUnsignedVarInt(uint32(c))
	for i := 0; i < c; i++ {
		IOList.List[i].WriteToBuffer(bs)
	}
}

func (IOList *InventoryActionIOList) ReadFromBuffer(bs *packets.MinecraftStream) *InventoryActionIOList{
	c := bs.GetUnsignedVarInt()
	for i := uint32(0); i < c; i ++{
		a := NewInventoryActionIO()
		a.ReadFromBuffer(bs)
		IOList.PutAction(a)
	}
	return IOList
}