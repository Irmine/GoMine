package bedrock

import (
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type PlayerSkinPacket struct {
	*packets.Packet

	UUID        uuid.UUID
	SkinId      string
	NewSkinName string
	OldSkinName string

	SkinData     []byte
	CapeData     []byte
	GeometryName string
	GeometryData string

	PremiumSkin  bool
}

func NewPlayerSkinPacket() *PlayerSkinPacket {
	return &PlayerSkinPacket{packets.NewPacket(info.PacketIds[info.PlayerSkinPacket]), uuid.New(), "", "", "", []byte{}, []byte{}, "", "", false}
}

func (pk *PlayerSkinPacket) Encode() {
	pk.PutUUID(pk.UUID)
	pk.PutString(pk.SkinId)
	pk.PutString(pk.NewSkinName)
	pk.PutString(pk.OldSkinName)
	pk.PutLengthPrefixedBytes(pk.SkinData)
	pk.PutLengthPrefixedBytes(pk.CapeData)
	pk.PutString(pk.GeometryName)
	pk.PutString(pk.GeometryData)
	pk.PutBool(pk.PremiumSkin)
}

func (pk *PlayerSkinPacket) Decode() {
	pk.UUID = pk.GetUUID()
	pk.SkinId = pk.GetString()
	pk.NewSkinName = pk.GetString()
	pk.OldSkinName = pk.GetString()
	pk.SkinData = pk.GetLengthPrefixedBytes()
	pk.CapeData = pk.GetLengthPrefixedBytes()
	pk.GeometryName = pk.GetString()
	pk.GeometryData = pk.GetString()
	pk.PremiumSkin = pk.GetBool()
}
