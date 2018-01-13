package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"gomine/net/packets"
	"goraklib/server"
	"crypto/x509"
	"crypto/ecdsa"
	"encoding/pem"
	"math/big"
	"crypto/sha512"
	"time"
)

type LoginHandler struct {
	*PacketHandler
}

func NewLoginHandler() LoginHandler {
	return LoginHandler{NewPacketHandler(info.LoginPacket)}
}

/**
 * Handles the main login process.
 */
func (handler LoginHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {

	if loginPacket, ok := packet.(*packets.LoginPacket); ok {
		_, online := server.GetPlayerFactory().GetPlayerByName(loginPacket.Username)
		if online == nil {
			return false
		}

		var successful, authenticated = handler.VerifyLoginRequest(loginPacket.Chains, server)
		if !successful {
			server.GetLogger().Error(loginPacket.Username, "had an invalid chain and is denied access")
			return false
		}
		if authenticated {
			server.GetLogger().Debug(loginPacket.Username, "has joined while being logged into XBOX Live")
		} else {
			server.GetLogger().Debug(loginPacket.Username, "tried to join, but is not logged into XBOX Live")
		}

		var player = player.New(server, session, loginPacket.Username, loginPacket.ClientUUID, loginPacket.ClientXUID, loginPacket.ClientId)
		player.SetLanguage(loginPacket.Language)
		player.SetSkinId(loginPacket.SkinId)
		player.SetSkinData(loginPacket.SkinData)
		player.SetCapeData(loginPacket.CapeData)
		player.SetGeometryName(loginPacket.GeometryName)
		player.SetGeometryData(loginPacket.GeometryData)

		playStatus := packets.NewPlayStatusPacket()
		playStatus.Status = 0
		player.SendPacket(playStatus)

		resourceInfo := packets.NewResourcePackInfoPacket()
		resourceInfo.MustAccept = server.GetConfiguration().ForceResourcePacks

		resourceInfo.ResourcePacks = server.GetPackHandler().GetResourceStack().GetPacks()
		resourceInfo.BehaviorPacks = server.GetPackHandler().GetBehaviorStack().GetPacks()

		player.SendPacket(resourceInfo)

		server.GetPlayerFactory().AddPlayer(player, session)

		return true
	}

	return false
}

func (handler LoginHandler) VerifyLoginRequest(chains []packets.Chain, server interfaces.IServer) (successful bool, authenticated bool) {
	var publicKeyRaw = ""
	var publicKey *ecdsa.PublicKey
	for _, chain := range chains {
		if publicKeyRaw == "" {
			if chain.Header.X5u == "" {
				return false, authenticated
			}
			publicKeyRaw = chain.Header.X5u
		}

		sig := []byte(chain.Signature)
		data := []byte(chain.Header.Raw + "." + chain.Payload.Raw)

		block, _ := pem.Decode([]byte("-----BEGIN PUBLIC KEY-----\n" + publicKeyRaw + "\n-----END PUBLIC KEY-----"))

		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			server.GetLogger().LogError(err)
			return false, authenticated
		}

		hash := sha512.New384()
		hash.Write(data)

		publicKey = key.(*ecdsa.PublicKey)
		r := new(big.Int).SetBytes(sig[:len(sig) / 2])
		s := new(big.Int).SetBytes(sig[len(sig) / 2:])

		if !ecdsa.Verify(publicKey, hash.Sum(nil), r, s) {
			return false, authenticated
		}

		if publicKeyRaw == packets.MojangPublicKey {
			authenticated = true
		}

		t := time.Now().Unix()
		if chain.Payload.ExpirationTime <= t {
			return false, authenticated
		}

		if chain.Payload.NotBefore > t {
			return false, authenticated
		}

		if chain.Payload.IssuedAt > chain.Payload.ExpirationTime {
			return false, authenticated
		}

		publicKeyRaw = chain.Payload.IdentityPublicKey
	}
	return true, authenticated
}