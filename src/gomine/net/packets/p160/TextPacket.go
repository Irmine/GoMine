package p160

import (
	"gomine/net/info"
	"gomine/net/packets"
	"gomine/net/packets/data"
)

type TextPacket struct {
	*packets.Packet
	TextType byte
	IsTranslation bool
	TranslationParameters []string
	SourceName string
	Message string
	XUID string
}

func NewTextPacket() *TextPacket {
	return &TextPacket{packets.NewPacket(info.PacketIds200[info.TextPacket]), 0, false, []string{}, "", "", ""}
}

func (pk *TextPacket) Decode() {
	pk.TextType = pk.GetByte()
	pk.IsTranslation = pk.GetBool()

	switch pk.TextType {
	case data.TextChat:
		fallthrough
	case data.TextAnnouncement:
		fallthrough
	case data.TextWhisper:
		pk.SourceName = pk.GetString()
		fallthrough
	case data.TextRaw:
		fallthrough
	case data.TextTip:
		fallthrough
	case data.TextSystem:
		pk.Message = pk.GetString()
	case data.TextTranslation:
		fallthrough
	case data.TextPopup:
		fallthrough
	case data.TextJukeboxPopup:
		pk.Message = pk.GetString()
		var translationParameterCount = pk.GetUnsignedVarInt()
		for i := uint32(0); i < translationParameterCount; i++ {
			pk.TranslationParameters = append(pk.TranslationParameters, pk.GetString())
		}
	}
	pk.XUID = pk.GetString()
}

func (pk *TextPacket) Encode() {
	pk.PutByte(pk.TextType)
	pk.PutBool(pk.IsTranslation)

	switch pk.TextType {
	case data.TextChat:
		fallthrough
	case data.TextWhisper:
		fallthrough
	case data.TextAnnouncement:
		pk.PutString(pk.SourceName)
		fallthrough
	case data.TextRaw:
		fallthrough
	case data.TextTip:
		fallthrough
	case data.TextSystem:
		pk.PutString(pk.Message)
	case data.TextTranslation:
		fallthrough
	case data.TextPopup:
		fallthrough
	case data.TextJukeboxPopup:
		pk.PutString(pk.Message)
		var count = len(pk.TranslationParameters)
		pk.PutUnsignedVarInt(uint32(count))
		for _, parameter := range pk.TranslationParameters {
			pk.PutString(parameter)
		}
	}
	pk.PutString(pk.XUID)
}
