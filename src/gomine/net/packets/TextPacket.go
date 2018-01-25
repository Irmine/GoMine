package packets

import "gomine/net/info"

const (
	TextRaw = iota
	TextChat
	TextTranslation
	TextPopup
	TextJukeboxPopup
	TextTip
	TextSystem
	TextWhisper
	TextAnnouncement
)

type TextPacket struct {
	*Packet
	TextType byte
	IsTranslation bool
	TranslationParameters []string
	SourceName string
	SourceDisplayName string
	SourcePlatform int32
	Message string
	XUID string
	UnknownString string
}

func NewTextPacket() *TextPacket {
	return &TextPacket{NewPacket(info.TextPacket), 0, false, []string{}, "", "", 0, "", "", ""}
}

func (pk *TextPacket) Decode() {
	pk.TextType = pk.GetByte()
	pk.IsTranslation = pk.GetBool()

	switch pk.TextType {
	case TextChat:
		fallthrough
	case TextAnnouncement:
		fallthrough
	case TextWhisper:
		pk.SourceName = pk.GetString()
		pk.SourceDisplayName = pk.GetString()
		pk.SourcePlatform = pk.GetVarInt()
		fallthrough
	case TextRaw:
		fallthrough
	case TextTip:
		fallthrough
	case TextSystem:
		pk.Message = pk.GetString()
	case TextTranslation:
		fallthrough
	case TextPopup:
		fallthrough
	case TextJukeboxPopup:
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
	case TextChat:
		fallthrough
	case TextWhisper:
		fallthrough
	case TextAnnouncement:
		pk.PutString(pk.SourceName)
		pk.PutString(pk.SourceDisplayName)
		pk.PutVarInt(pk.SourcePlatform)
		fallthrough
	case TextRaw:
		fallthrough
	case TextTip:
		fallthrough
	case TextSystem:
		pk.PutString(pk.Message)
	case TextTranslation:
		fallthrough
	case TextPopup:
		fallthrough
	case TextJukeboxPopup:
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