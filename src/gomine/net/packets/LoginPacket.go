package packets

import (
	"encoding/json"
	"gomine/utils"
	"gomine/net/info"
	"encoding/base64"
	"strings"
)

const (
	MojangPublicKey = "MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE8ELkixyLcwlZryUQcu1TvPOmI2B7vX83ndnWRUaXm74wFfa5f/lwQNTfrLVHa2PmenpGI6JhIMUJaWZrjmMj90NoKNFSNBuKdm8rYiXsfaz3K36x/1U26HpG0ZxK/V1V"
)

type LoginPacket struct {
	*Packet
	Username string
	Protocol int32
	ClientUUID utils.UUID
	ClientId int
	ClientXUID string
	IdentityPublicKey string
	ServerAddress string
	Language string

	SkinId string
	SkinData []byte
	CapeData []byte
	GeometryName string
	GeometryData string

	ClientData ClientDataKeys
	Chains []Chain
}

type ChainDataKeys struct {
	RawChains []string `json:"chain"`
	Chains []Chain
}


type Chain struct {
	Header ChainHeader
	Payload ChainPayload
	Signature string
}

type ChainHeader struct {
	X5u string `json:"x5u"`
	Alg string `json:"alg"`

	Raw string
}

type ChainPayload struct {
	CertificateAuthority bool `json:"certificateAuthority"`
	ExpirationTime int64 `json:"exp"`
	IdentityPublicKey string `json:"identityPublicKey"`
	NotBefore int64 `json:"nbf"`
	RandomNonce int `json:"randomNonce"`
	Issuer string `json:"iss"`
	IssuedAt int64 `json:"iat"`

	Raw string
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
	CurrentInputMode string `json:"CurrentInputMode"`
	DefaultInputMode string `json:"DefaultInputMode"`
	DeviceModel string `json:"DeviceModel"`
	DeviceOS int `json:"DeviceOS"`
	GameVersion string `json:"GameVersion"`
	GuiScale int `json:"GuiScale"`
	UIProfile int `json:"UIProfile"`
}

func NewLoginPacket() *LoginPacket {
	pk := &LoginPacket{NewPacket(info.LoginPacket), "", 0, utils.UUID{}, 0, "", "", "", "", "", []byte{}, []byte{}, "", "", ClientDataKeys{}, []Chain{}}
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

	var length = int(stream.GetLittleInt())

	var chainData = &ChainDataKeys{}
	json.Unmarshal(stream.Get(length), &chainData)

	for _, v := range chainData.RawChains {
		WebToken := &WebTokenKeys{}
		pk.Chains = append(pk.Chains, pk.BuildChain(v))

		utils.DecodeJwtPayload(v, WebToken)

		if v, ok := WebToken.ExtraData["displayName"]; ok {
			pk.Username = v.(string)
		}
		if v, ok := WebToken.ExtraData["identity"]; ok {
			pk.ClientUUID = utils.UUIDFromString(v.(string))
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

	utils.DecodeJwtPayload(string(clientDataJwt), clientData)

	pk.ClientId = clientData.ClientRandomId
	pk.ServerAddress = clientData.ServerAddress

	pk.Language = clientData.LanguageCode
	if pk.Language == "" {
		pk.Language = "en_US"
	}

	pk.SkinId = clientData.SkinId
	pk.SkinData, _ = base64.RawURLEncoding.DecodeString(clientData.SkinData)
	pk.CapeData, _ = base64.RawURLEncoding.DecodeString(clientData.CapeData)
	var geometry, _ = base64.RawURLEncoding.DecodeString(clientData.GeometryData)
	pk.GeometryData = string(geometry)

	pk.ClientData = *clientData
}

func (pk *LoginPacket) BuildChain(raw string) Chain {
	jwt := utils.DecodeJwt(raw)
	var base64s = strings.Split(raw, ".")

	chain := Chain{}
	for i, str := range jwt {
		switch i {
		case 0:
			header := ChainHeader{}
			json.Unmarshal([]byte(str), &header)
			header.Raw = base64s[i]
			chain.Header = header
		case 1:
			payload := ChainPayload{}
			json.Unmarshal([]byte(str), &payload)
			payload.Raw = base64s[i]
			chain.Payload = payload
		case 2:
			chain.Signature = str
		}
	}
	return chain
}