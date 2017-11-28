package packets

import (
	"encoding/json"
	"gomine/utils"
	"gomine/net/info"
	"fmt"
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
}

type ChainDataKeys struct {
	Chain []string `json:"chain"`
}

type WebTokenKeys struct {
	ExtraData map[string]interface{} `json:"extraData"`
	IdentityPublicKey string `json:"identityPublicKey"`
}

type ClientDataKeys struct {
	ClientData map[string]interface{} `json:"clientData"`
}

func NewLoginPacket() *LoginPacket {
	pk := &LoginPacket{NewPacket(info.LoginPacket), "", 0, "", 0, "", "", "", ""}
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
	fmt.Println(clientData)

	if v, ok := clientData.ClientData["ClientRandomId"]; ok {
		pk.ClientId = v.(int)
	}else{
		pk.ClientId = 0
	}

	if v, ok := clientData.ClientData["ServerAddress"]; ok {
		pk.ServerAddress = v.(string)
	}else{
		pk.ServerAddress = ""
	}

	if v, ok := clientData.ClientData["LanguageCode"]; ok {
		pk.Language = v.(string)
		fmt.Println(pk.Language)
	}else{
		pk.Language = "en_US"
	}
}