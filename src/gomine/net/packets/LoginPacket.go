package packets

import (
	"gomine/net"
	"encoding/json"
	"gomine/utils"
)

type LoginPacket struct {
	*Packet
	Username string
	Protocol int32
	ClientUUID string
	ClientId int
	Xuid string
	IdentityPublicKey string
	ServerAddress string
	Language string
}

type ChainDataKeys struct {
	chain map[string]string
}

type WebTokenKeys struct {
	extraData map[string]interface{}
	identityPublicKey string
}

type ClientDataKeys struct {
	clientData map[string]interface{}
}

var ChainData = ChainDataKeys{}
var ClientDataJwt []byte
var ClientData = ClientDataKeys{}

func NewLoginPacket() LoginPacket {
	pk := LoginPacket{NewPacket(net.Login), "", 0, "", 0, "", "", "", ""}
	return pk
}

func (pk *LoginPacket) Encode()  {
	//todo
}

func (pk *LoginPacket) Decode()  {

	pk.Protocol = pk.GetInt()

	if pk.Protocol != net.LatestProtocol {
		if pk.Protocol > 0xffff {
			pk.Offset -= 6
			pk.Protocol = pk.GetInt()
		}
	}

	json.Unmarshal(pk.Get(int(pk.GetLittleInt())), ChainData)
	for _, v := range ChainData.chain {
		WebToken := WebTokenKeys{}
		utils.DecodeJwt(v, WebToken)
		if v, ok := WebToken.extraData["username"]; ok {
			pk.Username = v.(string)
		}
		if v, ok := WebToken.extraData["indentity"]; ok {
			pk.ClientUUID = v.(string)
		}
		if v, ok := WebToken.extraData["XUID"]; ok {
			pk.Xuid = v.(string)
		}
		if len(WebToken.identityPublicKey) > 0 {
			pk.IdentityPublicKey = WebToken.identityPublicKey
		}
	}
	ClientDataJwt = pk.Get(int(pk.GetLittleInt()))
	utils.DecodeJwt(string(ClientDataJwt), ClientData)
	if v, ok := ClientData.clientData["ClientRandomId"]; ok {
		pk.ClientId = v.(int)
	}else{
		pk.ClientId = 0
	}
	if v, ok := ClientData.clientData["ServerAddress"]; ok {
		pk.ServerAddress = v.(string)
	}else{
		pk.ServerAddress = ""
	}
	if v, ok := ClientData.clientData["LanguageCode"]; ok {
		pk.Language = v.(string)
	}else{
		pk.Language = v.(string)
	}
}