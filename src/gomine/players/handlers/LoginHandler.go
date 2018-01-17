package handlers

import (
	"gomine/net/info"
	"gomine/interfaces"
	"gomine/net/packets"
	"goraklib/server"
	"crypto/x509"
	"crypto/ecdsa"
	"math/big"
	"crypto/sha512"
	"time"
	"gomine/utils"
	"encoding/base64"
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
		_, err := server.GetPlayerFactory().GetPlayerByName(loginPacket.Username)
		if err == nil {
			return false
		}

		var successful, authenticated, pubKey = handler.VerifyLoginRequest(loginPacket.Chains, server)
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
		player.GetEncryptionHandler().Data = &utils.EncryptionData{
			ClientPublicKey: pubKey,
			ServerPrivateKey: server.GetPrivateKey(),
			ServerToken: server.GetServerToken(),
		}

		player.SetLanguage(loginPacket.Language)
		player.SetSkinId(loginPacket.SkinId)
		player.SetSkinData(loginPacket.SkinData)
		player.SetCapeData(loginPacket.CapeData)
		player.SetGeometryName(loginPacket.GeometryName)
		player.SetGeometryData(loginPacket.GeometryData)

		var handshake = packets.NewServerHandshakePacket()
		var jwt = utils.ConstructEncryptionJwt(server.GetPrivateKey(), server.GetServerToken())
		utils.DecodeJwt(jwt)
		handshake.Jwt = jwt
		player.SendPacket(handshake)

		/*playStatus := packets.NewPlayStatusPacket()
		playStatus.Status = 0
		player.SendPacket(playStatus)

		resourceInfo := packets.NewResourcePackInfoPacket()
		resourceInfo.MustAccept = server.GetConfiguration().ForceResourcePacks

		resourceInfo.ResourcePacks = server.GetPackHandler().GetResourceStack().GetPacks()
		resourceInfo.BehaviorPacks = server.GetPackHandler().GetBehaviorStack().GetPacks()

		player.SendPacket(resourceInfo)*/

		server.GetPlayerFactory().AddPlayer(player, session)
		player.EnableEncryption()

		return true
	}

	return false
}

func (handler LoginHandler) VerifyLoginRequest(chains []packets.Chain, server interfaces.IServer) (successful bool, authenticated bool, clientPublicKey *ecdsa.PublicKey) {
	var publicKey *ecdsa.PublicKey
	var publicKeyRaw string
	for _, chain := range chains {
		if publicKeyRaw == "" {
			if chain.Header.X5u == "" {
				return
			}
			publicKeyRaw = chain.Header.X5u
		}

		sig := []byte(chain.Signature)
		data := []byte(chain.Header.Raw + "." + chain.Payload.Raw)

		var b64, errB64 = base64.RawStdEncoding.DecodeString(publicKeyRaw)
		server.GetLogger().LogError(errB64)

		key, err := x509.ParsePKIXPublicKey(b64)
		if err != nil {
			server.GetLogger().LogError(err)
			return
		}

		hash := sha512.New384()
		hash.Write(data)

		publicKey = key.(*ecdsa.PublicKey)
		r := new(big.Int).SetBytes(sig[:len(sig) / 2])
		s := new(big.Int).SetBytes(sig[len(sig) / 2:])

		if !ecdsa.Verify(publicKey, hash.Sum(nil), r, s) {
			return
		}

		if publicKeyRaw == packets.MojangPublicKey {
			authenticated = true
		}

		t := time.Now().Unix()
		if chain.Payload.ExpirationTime <= t && chain.Payload.ExpirationTime != 0 {
			return
		}

		if chain.Payload.NotBefore > t {
			return
		}

		if chain.Payload.IssuedAt > chain.Payload.ExpirationTime {
			return
		}

		publicKeyRaw = chain.Payload.IdentityPublicKey
	}

	var b64, errB64 = base64.RawStdEncoding.DecodeString(publicKeyRaw)
	server.GetLogger().LogError(errB64)

	key, err := x509.ParsePKIXPublicKey(b64)
	if err != nil {
		server.GetLogger().LogError(err)
		return
	}

	clientPublicKey = key.(*ecdsa.PublicKey)

	successful = true
	return
}