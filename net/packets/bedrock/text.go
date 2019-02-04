package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/data"
)

type TextPacket struct {
	*packets.Packet
	TextType byte
	Translation bool
	SourceName string
	Message string
	XUID string
	PlatformChatId string

	Params []string
}

func NewTextPacket() *TextPacket {
	return &TextPacket{
		packets.NewPacket(info.PacketIds[info.TextPacket]),
		data.TextRaw,
		false,
		"",
		"",
		"",
		"",
		[]string{},
	}
}

func (pk *TextPacket) Encode()  {
	pk.PutByte(pk.TextType)
	pk.PutBool(pk.Translation)

	switch pk.TextType {
	case data.TextRaw, data.TextTip, data.TextSystem:
		pk.PutString(pk.Message)
		break
	case data.TextChat, data.TextWhisper, data.TextAnnouncement:
		pk.PutString(pk.SourceName)
		pk.PutString(pk.Message)
		break
	case data.TextTranslation, data.TextPopup, data.TextJukeboxPopup:
		pk.PutString(pk.Message)
		pk.PutUnsignedVarInt(uint32(len(pk.Params)))
		for _, v := range pk.Params {
			pk.PutString(v)
		}
		break
	}

	pk.PutString(pk.XUID)
	pk.PutString(pk.PlatformChatId)
}

func (pk *TextPacket) Decode() {
	pk.TextType = pk.GetByte()
	pk.Translation = pk.GetBool()

	switch pk.TextType {
	case data.TextRaw, data.TextTip, data.TextSystem:
		pk.Message = pk.GetString()
		break
	case data.TextChat, data.TextWhisper, data.TextAnnouncement:
		pk.SourceName = pk.GetString()
		pk.Message = pk.GetString()
		break
	case data.TextTranslation, data.TextPopup, data.TextJukeboxPopup:
		pk.Message = pk.GetString()
		c := pk.GetUnsignedVarInt()
		for i := uint32(0); i < c; i++ {
			pk.Params = append(pk.Params, pk.GetString())
		}
		break
	}

	pk.XUID = pk.GetString()
	pk.PlatformChatId = pk.GetString()
}