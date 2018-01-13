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
	Players map[string]interfaces.IPlayer
}

func NewPlayerListPacket() *PlayerListPacket {
	return &PlayerListPacket{NewPacket(info.PlayerListPacket), 0, map[string]interfaces.IPlayer{}}
}

func (pk *PlayerListPacket) Encode() {
	pk.PutByte(pk.ListType)
	pk.PutUnsignedVarInt(uint32(len(pk.Players)))
	for _, entry := range pk.Players {
		if pk.ListType == byte(ListTypeAdd) {
			pk.PutUUID(entry.GetUUID())
			pk.PutUniqueId(entry.GetUniqueId())

			pk.PutString(entry.GetDisplayName())
			pk.PutString(entry.GetSkinId())
			pk.PutString(string(entry.GetSkinData()))
			pk.PutString(string(entry.GetCapeData()))
			pk.PutString(entry.GetGeometryName())
			pk.PutString(entry.GetGeometryData())

			pk.PutString(entry.GetXUID())
		} else {
			pk.PutUUID(entry.GetUUID())
		}
	}
}

func (pk *PlayerListPacket) Decode() {

}