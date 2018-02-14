package p160

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/types"
)

type PlayerListPacket struct {
	*packets.Packet
	ListType byte
	Entries  map[string]types.PlayerListEntry
}

func NewPlayerListPacket() *PlayerListPacket {
	return &PlayerListPacket{packets.NewPacket(info.PacketIds200[info.PlayerListPacket]), 0, map[string]types.PlayerListEntry{}}
}

func (pk *PlayerListPacket) Encode() {
	pk.PutByte(pk.ListType)
	pk.PutUnsignedVarInt(uint32(len(pk.Entries)))
	for _, entry := range pk.Entries {
		if pk.ListType == byte(data.ListTypeAdd) {
			pk.PutUUID(entry.UUID)
			pk.PutUniqueId(entry.EntityUniqueId)

			pk.PutString(entry.DisplayName)
			pk.PutString(entry.SkinId)
			pk.PutLengthPrefixedBytes(entry.SkinData)
			pk.PutLengthPrefixedBytes(entry.CapeData)

			pk.PutString(entry.GeometryName)
			pk.PutString(entry.GeometryData)

			pk.PutString(entry.XUID)
		} else {
			pk.PutUUID(entry.UUID)
		}
	}
}

func (pk *PlayerListPacket) Decode() {

}
