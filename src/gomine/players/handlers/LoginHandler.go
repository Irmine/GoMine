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

		handler.VerifyLoginRequest(loginPacket.Chains, server)

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

func (handler LoginHandler) VerifyLoginRequest(chains []packets.Chain, server interfaces.IServer) bool {
	var publicKeyRaw = ""
	var publicKey *ecdsa.PublicKey
	for _, chain := range chains {
		if publicKeyRaw == "" {
			if chain.Header.X5u == "" {
				return false
			}
		}

		sig := []byte(chain.Signature)
		data := []byte(chain.Header.Raw + "." + chain.Payload.Raw)

		publicKeyRaw = chain.Header.X5u
		block, _ := pem.Decode([]byte("-----BEGIN PUBLIC KEY-----\n" + publicKeyRaw + "\n-----END PUBLIC KEY-----"))

		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			server.GetLogger().LogError(err)
			return false
		}

		hash := sha512.New384()
		hash.Write(data)

		publicKey = key.(*ecdsa.PublicKey)
		r := new(big.Int).SetBytes(sig[:len(sig) / 2])
		s := new(big.Int).SetBytes(sig[len(sig) / 2:])

		println("Signature validation:", ecdsa.Verify(publicKey, hash.Sum(nil), r, s))

		println(chain.Header.Alg)
	}
	return true
}