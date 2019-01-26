package mcpe

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/data"
)

type TextPacket struct {
	*packets.Packet
	TextType              byte
	IsTranslation         bool
	TranslationParameters []string
	SourceName            string
	SourceDisplayName     string
	SourcePlatform        int32
	Message               string
	XUID                  string
	UnknownString         string
}

func NewTextPacket() *TextPacket {
	return &TextPacket{packets.NewPacket(info.PacketIds[info.TextPacket]), 0, false, []string{}, "", "", 0, "", "", ""}
}

func (pk *TextPacket) Decode() {
	pk.TextType = pk.GetByte()
	pk.IsTranslation = pk.GetBool()

	switch pk.TextType {
	case data.TextChat, data.TextAnnouncement, data.TextWhisper:
		pk.SourceName = pk.GetString()
		fallthrough
	case data.TextRaw, data.TextTip, data.TextSystem:
		pk.Message = pk.GetString()
	case data.TextTranslation, data.TextPopup, data.TextJukeboxPopup:
		pk.Message = pk.GetString()
		var translationParameterCount = pk.GetUnsignedVarInt()
		for i := uint32(0); i < translationParameterCount; i++ {
			pk.TranslationParameters = append(pk.TranslationParameters, pk.GetString())
		}
	}
	pk.XUID = pk.GetString()
	pk.UnknownString = pk.GetString()
}

func (pk *TextPacket) Encode() {
	pk.PutByte(pk.TextType)
	pk.PutBool(pk.IsTranslation)

	switch pk.TextType {
	case data.TextChat, data.TextWhisper, data.TextAnnouncement:
		pk.PutString(pk.SourceName)
		fallthrough
	case data.TextRaw, data.TextTip, data.TextSystem:
		pk.PutString(pk.Message)
	case data.TextTranslation, data.TextPopup, data.TextJukeboxPopup:
		pk.PutString(pk.Message)
		var count = len(pk.TranslationParameters)
		pk.PutUnsignedVarInt(uint32(count))
		for _, parameter := range pk.TranslationParameters {
			pk.PutString(parameter)
		}
	}
	pk.PutString(pk.XUID)
	pk.PutString(pk.UnknownString)
}
