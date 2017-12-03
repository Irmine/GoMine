package packets

import (
	"encoding/json"
	"gomine/utils"
	"gomine/net/info"
	"encoding/base64"
)

type LoginPacket struct {
	*Packet
	Username string
	Protocol int32
	ClientUUID string
	ClientId int
	ClientXUID string
	IdentityPublicKey string
	ServerAddress string
	Language string

	SkinId string
	SkinData []byte
	CapeData []byte
	GeometryName string
	GeometryData []byte
}

type ChainDataKeys struct {
	Chain []string `json:"chain"`
}

type WebTokenKeys struct {
	ExtraData map[string]interface{} `json:"extraData"`
	IdentityPublicKey string `json:"identityPublicKey"`
}

type ClientDataKeys struct {
	ClientRandomId int `json:"ClientRandomId"`
	ServerAddress string `json:"ServerAddress"`
	LanguageCode string `json:"LanguageCode"`
	SkinId string `json:"SkinId"`
	SkinData string `json:"SkinData"`
	CapeData string `json:"CapeData"`
	GeometryId string `json:"SkinGeometryName"`
	GeometryData string `json:"SkinGeometry"`
}

func NewLoginPacket() *LoginPacket {
	pk := &LoginPacket{NewPacket(info.LoginPacket), "", 0, "", 0, "", "", "", "", "", []byte{}, []byte{}, "", []byte{}}
	return pk
}

func (pk *LoginPacket) Encode()  {

}

func (pk *LoginPacket) Decode()  {
	pk.Protocol = pk.GetInt()

	if pk.Protocol != info.LatestProtocol {
		if pk.Protocol > 0xffff {
			pk.Offset -= 6
			pk.Protocol = pk.GetInt()
		}
	}

	var stream = utils.NewStream()
	stream.Buffer = []byte(pk.GetString())


	var length = stream.GetLittleInt()

	var chainData = &ChainDataKeys{}
	json.Unmarshal(stream.Get(int(length)), &chainData)

	for _, v := range chainData.Chain {
		WebToken := &WebTokenKeys{}

		utils.DecodeJwt(v, WebToken)

		if v, ok := WebToken.ExtraData["displayName"]; ok {
			pk.Username = v.(string)
		}
		if v, ok := WebToken.ExtraData["identity"]; ok {
			pk.ClientUUID = v.(string)
		}
		if v, ok := WebToken.ExtraData["XUID"]; ok {
			pk.ClientXUID = v.(string)
		}
		if len(WebToken.IdentityPublicKey) > 0 {
			pk.IdentityPublicKey = WebToken.IdentityPublicKey
		}
	}

	var clientDataJwt = stream.Get(int(stream.GetLittleInt()))
	var clientData = &ClientDataKeys{}

	utils.DecodeJwt(string(clientDataJwt), clientData)

	pk.ClientId = clientData.ClientRandomId
	pk.ServerAddress = clientData.ServerAddress

	pk.Language = clientData.LanguageCode
	if pk.Language == "" {
		pk.Language = "en_US"
	}

	pk.SkinId = clientData.SkinId
	pk.SkinData, _ = base64.RawStdEncoding.DecodeString(clientData.SkinData)
	pk.CapeData, _ = base64.RawStdEncoding.DecodeString(clientData.CapeData)
	pk.GeometryData, _ = base64.RawStdEncoding.DecodeString(clientData.GeometryData)
}