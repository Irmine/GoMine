package handlers

import (
	"gomine/interfaces"
	"goraklib/server"
	"crypto/x509"
	"crypto/ecdsa"
	"math/big"
	"crypto/sha512"
	"time"
	"gomine/utils"
	"encoding/base64"
	"gomine/net/packets/types"
	data2 "gomine/net/packets/data"
	"gomine/net/packets/p200"
)

type LoginHandler struct {
	*PacketHandler
}

func NewLoginHandler() LoginHandler {
	return LoginHandler{NewPacketHandler()}
}

/**
 * Handles the main login process.
 */
func (handler LoginHandler) Handle(packet interfaces.IPacket, player interfaces.IPlayer, session *server.Session, server interfaces.IServer) bool {

	if loginPacket, ok := packet.(*p200.LoginPacket); ok {
		_, err := server.GetPlayerFactory().GetPlayerByName(loginPacket.Username)
		if err == nil {
			return false
		}

		var successful, authenticated, pubKey = handler.VerifyLoginRequest(loginPacket.Chains, server)
		if !successful {
			server.GetLogger().Debug(loginPacket.Username, "has joined with invalid login data.")
			return true
		}

		if authenticated {
			server.GetLogger().Debug(loginPacket.Username, "has joined while being logged into XBOX Live.")
		} else {
			if server.GetConfiguration().XBOXLiveAuth {
				server.GetLogger().Debug(loginPacket.Username, "has tried to join while not being logged into XBOX Live.")
				return true
			}
			server.GetLogger().Debug(loginPacket.Username, "has joined while not being logged into XBOX Live.")
		}

		var s = player.NewMinecraftSession(server, session, types.SessionData{
			loginPacket.ClientUUID,
			loginPacket.ClientXUID,
			loginPacket.ClientId,
			loginPacket.Protocol,
			loginPacket.ClientData.GameVersion,
			loginPacket.Language,
			loginPacket.ClientData.DeviceOS,
		})
		s.GetEncryptionHandler().Data = &utils.EncryptionData{
			ClientPublicKey: pubKey,
			ServerPrivateKey: server.GetPrivateKey(),
			ServerToken: server.GetServerToken(),
		}

		player.SetMinecraftSession(s)

		player.SetDisplayName(loginPacket.Username)
		player.SetDisplayName(loginPacket.ClientData.ThirdPartyName)

		player.SetLanguage(loginPacket.Language)
		player.SetSkinId(loginPacket.SkinId)
		player.SetSkinData(loginPacket.SkinData)
		player.SetCapeData(loginPacket.CapeData)
		player.SetGeometryName(loginPacket.GeometryName)
		player.SetGeometryData(loginPacket.GeometryData)

		player.SetXBOXLiveAuthenticated(authenticated)

		if server.GetConfiguration().UseEncryption {
			var jwt = utils.ConstructEncryptionJwt(server.GetPrivateKey(), server.GetServerToken())
			player.SendServerHandshake(jwt)

			player.EnableEncryption()
		} else {
			player.SendPlayStatus(data2.StatusLoginSuccess)

			player.SendResourcePackInfo(server.GetConfiguration().ForceResourcePacks, server.GetPackHandler().GetResourceStack().GetPacks(), server.GetPackHandler().GetBehaviorStack().GetPacks())
		}

		return true
	}

	return false
}

func (handler LoginHandler) VerifyLoginRequest(chains []types.Chain, server interfaces.IServer) (successful bool, authenticated bool, clientPublicKey *ecdsa.PublicKey) {
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

		if publicKeyRaw == data2.MojangPublicKey {
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