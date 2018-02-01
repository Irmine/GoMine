package p200

import (
	"gomine/utils"
	"gomine/net/info"
	"gomine/net/packets"
)

type PlayerSkinPacket struct {
	*packets.Packet

	UUID utils.UUID
	SkinId string
	NewSkinName string
	OldSkinName string

	SkinData []byte
	CapeData []byte
	GeometryName string
	GeometryData string
}

func NewPlayerSkinPacket() *PlayerSkinPacket {
	return &PlayerSkinPacket{packets.NewPacket(info.PlayerSkinPacket), utils.UUID{}, "", "", "", []byte{}, []byte{}, "", ""}
}

func (pk *PlayerSkinPacket) Encode() {
	pk.UUID = pk.GetUUID()
	pk.SkinId = pk.GetString()
	pk.NewSkinName = pk.GetString()
	pk.OldSkinName = pk.GetString()
	pk.SkinData = pk.GetLengthPrefixedBytes()
	pk.CapeData = pk.GetLengthPrefixedBytes()
	pk.GeometryName = pk.GetString()
	pk.GeometryData = pk.GetString()
}

func (pk *PlayerSkinPacket) Decode() {
	pk.PutUUID(pk.UUID)
	pk.PutString(pk.SkinId)
	pk.PutString(pk.NewSkinName)
	pk.PutString(pk.OldSkinName)
	pk.PutLengthPrefixedBytes(pk.SkinData)
	pk.PutLengthPrefixedBytes(pk.CapeData)
	pk.PutString(pk.GeometryName)
	pk.PutString(pk.GeometryData)
}