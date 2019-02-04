package bedrock

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/irmine/binutils"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/utils"
)

type LoginPacket struct {
	*packets.Packet
	Username          string
	Protocol          int32
	ClientUUID        uuid.UUID
	ClientId          int
	ClientXUID        string
	IdentityPublicKey string
	ServerAddress     string
	Language          string

	SkinId       string
	SkinData     []byte
	CapeData     []byte
	GeometryName string
	GeometryData string

	ClientData types.ClientDataKeys
	Chains     []types.Chain
}

func NewLoginPacket() *LoginPacket {
	pk := &LoginPacket{packets.NewPacket(info.PacketIds[info.LoginPacket]), "", 0, uuid.New(), 0, "", "", "", "", "", []byte{}, []byte{}, "", "", types.ClientDataKeys{}, []types.Chain{}}
	return pk
}

func (pk *LoginPacket) Encode() {

}

func (pk *LoginPacket) Decode() {
	pk.Protocol = pk.GetInt()

	var stream = binutils.NewStream()
	stream.Buffer = []byte(pk.GetString())

	var length = int(stream.GetLittleInt())

	var chainData = &types.ChainDataKeys{}
	json.Unmarshal(stream.Get(length), &chainData)

	for _, v := range chainData.RawChains {
		WebToken := &types.WebTokenKeys{}
		pk.Chains = append(pk.Chains, pk.BuildChain(v))

		utils.DecodeJwtPayload(v, WebToken)

		if v, ok := WebToken.ExtraData["displayName"]; ok {
			pk.Username = v.(string)
		}
		if v, ok := WebToken.ExtraData["identity"]; ok {
			pk.ClientUUID = uuid.Must(uuid.Parse(v.(string)))
		}
		if v, ok := WebToken.ExtraData["XUID"]; ok {
			pk.ClientXUID = v.(string)
		}
		if len(WebToken.IdentityPublicKey) > 0 {
			pk.IdentityPublicKey = WebToken.IdentityPublicKey
		}
	}

	var clientDataJwt = stream.Get(int(stream.GetLittleInt()))
	var clientData = &types.ClientDataKeys{}

	utils.DecodeJwtPayload(string(clientDataJwt), clientData)

	pk.ClientId = clientData.ClientRandomId
	pk.ServerAddress = clientData.ServerAddress

	pk.Language = clientData.LanguageCode
	if pk.Language == "" {
		pk.Language = "en_US"
	}

	pk.SkinId = clientData.SkinId
	pk.GeometryName = clientData.GeometryId
	pk.SkinData, _ = base64.RawStdEncoding.DecodeString(clientData.SkinData)
	pk.CapeData, _ = base64.RawStdEncoding.DecodeString(clientData.CapeData)
	var geometry, _ = base64.RawStdEncoding.DecodeString(clientData.GeometryData)
	pk.GeometryData = string(geometry)

	for len(pk.SkinData) < 16384 {
		pk.SkinData = append(pk.SkinData, 0x00)
	}

	pk.ClientData = *clientData
}

func (pk *LoginPacket) BuildChain(raw string) types.Chain {
	jwt := utils.DecodeJwt(raw)
	var base64s = strings.Split(raw, ".")

	chain := types.Chain{}
	for i, str := range jwt {
		switch i {
		case 0:
			header := types.ChainHeader{}
			json.Unmarshal([]byte(str), &header)
			header.Raw = base64s[i]
			chain.Header = header
		case 1:
			payload := types.ChainPayload{}
			json.Unmarshal([]byte(str), &payload)
			payload.Raw = base64s[i]
			chain.Payload = payload
		case 2:
			chain.Signature = str
		}
	}
	return chain
}
