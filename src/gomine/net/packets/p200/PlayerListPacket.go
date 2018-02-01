package p200

import (
	"gomine/net/info"
	"gomine/interfaces"
	"gomine/net/packets"
)

const (
	ListTypeAdd = iota
	ListTypeRemove
)

type PlayerListPacket struct {
	*packets.Packet
	ListType byte
	Players map[string]interfaces.IPlayer
}

func NewPlayerListPacket() *PlayerListPacket {
	return &PlayerListPacket{packets.NewPacket(info.PlayerListPacket), 0, map[string]interfaces.IPlayer{}}
}

func (pk *PlayerListPacket) Encode() {
	pk.PutByte(pk.ListType)
	pk.PutUnsignedVarInt(uint32(len(pk.Players)))
	for _, entry := range pk.Players {
		if pk.ListType == byte(ListTypeAdd) {
			pk.PutUUID(entry.GetUUID())
			pk.PutUniqueId(entry.GetUniqueId())

			pk.PutString(entry.GetName())
			pk.PutString(entry.GetDisplayName())
			pk.PutVarInt(entry.GetPlatform())
			pk.PutString(entry.GetSkinId())

			pk.PutLittleInt(1)
			pk.PutLengthPrefixedBytes(entry.GetSkinData())
			if len(entry.GetCapeData()) > 0{
				pk.PutLittleInt(1)
				pk.PutLengthPrefixedBytes(entry.GetCapeData())
			} else {
				pk.PutLittleInt(0)
			}

			pk.PutString(entry.GetGeometryName())
			pk.PutString(entry.GetGeometryData())

			pk.PutString(entry.GetXUID())
			pk.PutString("")
		} else {
			pk.PutUUID(entry.GetUUID())
		}
	}
}

func (pk *PlayerListPacket) Decode() {

}