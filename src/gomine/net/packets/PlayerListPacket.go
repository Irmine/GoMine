package packets

import (
	"gomine/net/info"
	"gomine/interfaces"
)

const (
	ListTypeAdd = iota
	ListTypeRemove
)

type PlayerListPacket struct {
	*Packet
	ListType byte
	Players []interfaces.IPlayer
}

func NewPlayerListPacket() *PlayerListPacket {
	return &PlayerListPacket{NewPacket(info.PlayerListPacket), 0, []interfaces.IPlayer{}}
}

func (pk *PlayerListPacket) Encode() {
	pk.PutByte(pk.ListType)
	pk.PutUnsignedVarInt(uint32(len(pk.Players)))
	for _, entry := range pk.Players {
		if pk.ListType == byte(ListTypeAdd) {
			pk.PutLittleInt(0)
			pk.PutLittleInt(0)
			pk.PutLittleInt(0)
			pk.PutLittleInt(0)

			pk.PutVarLong(0)

			pk.PutString(entry.GetDisplayName())
			pk.PutString(entry.GetSkinId())
			pk.PutString(string(entry.GetSkinData()))
			pk.PutString(string(entry.GetCapeData()))
			pk.PutString(entry.GetGeometryName())
			pk.PutString(entry.GetGeometryData())

			pk.PutString(entry.GetXUID())
		}
	}
}

func (pk *PlayerListPacket) Decode() {

}
