package gomine

import (
	"crypto/ecdsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/players"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/worlds/chunks"
	"math/big"
	"strings"
	"time"
)

func NewClientHandshakeHandler_200(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, logger *utils.Logger, session *net.MinecraftSession) bool {
		if _, ok := packet.(*p200.ClientHandshakePacket); ok {
			session.SendPlayStatus(data.StatusLoginSuccess)
			session.SendResourcePackInfo(server.GetConfiguration().ForceResourcePacks, server.GetPackManager().GetResourceStack(), server.GetPackManager().GetBehaviorStack())
			return true
		}
		return false
	})
}

func NewCommandRequestHandler_200(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, logger *utils.Logger, session *net.MinecraftSession) bool {
		if pk, ok := packet.(*p200.CommandRequestPacket); ok {
			var args = strings.Split(pk.CommandText, " ")
			var commandName = strings.TrimLeft(args[0], "/")
			var i = 1
			for !server.GetCommandManager().IsCommandRegistered(commandName) {
				if i == len(args) {
					break
				}
				commandName += " " + args[i]
				i++
			}
			if !server.GetCommandManager().IsCommandRegistered(commandName) {
				session.SendMessage("Command could not be found.")
				return false
			}
			args = args[i:]
			var command, _ = server.GetCommandManager().GetCommand(commandName)
			command.Execute(session, args)

			return true
		}

		return false
	})
}

func NewLoginHandler_200(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, logger *utils.Logger, session *net.MinecraftSession) bool {
		if loginPacket, ok := packet.(*p200.LoginPacket); ok {
			var _, ok = server.GetSessionManager().GetSession(loginPacket.Username)
			if ok {
				return false
			}
			if !server.GetNetworkAdapter().GetProtocolManager().IsProtocolRegistered(loginPacket.Protocol) {
				server.GetLogger().Debug(loginPacket.Username, "tried joining with unsupported protocol:", loginPacket.Protocol)
				return false
			}

			var successful, authenticated, pubKey = VerifyLoginRequest(loginPacket.Chains, server)
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

			session.SetData(server.GetPermissionManager(), types.SessionData{loginPacket.ClientUUID, loginPacket.ClientXUID, loginPacket.ClientId, loginPacket.Protocol, loginPacket.ClientData.GameVersion, loginPacket.Language, loginPacket.ClientData.DeviceOS})
			session.SetPlayer(players.NewPlayer(loginPacket.ClientUUID, loginPacket.ClientXUID, int32(loginPacket.ClientData.DeviceOS), loginPacket.Username))

			session.GetEncryptionHandler().Data = &utils.EncryptionData{
				ClientPublicKey:  pubKey,
				ServerPrivateKey: server.GetPrivateKey(),
				ServerToken:      server.GetServerToken(),
			}

			session.GetPlayer().SetName(loginPacket.Username)
			session.GetPlayer().SetDisplayName(loginPacket.Username)
			session.GetPlayer().SetSkinId(loginPacket.SkinId)
			session.GetPlayer().SetSkinData(loginPacket.SkinData)
			session.GetPlayer().SetCapeData(loginPacket.CapeData)
			session.GetPlayer().SetGeometryName(loginPacket.GeometryName)
			session.GetPlayer().SetGeometryData(loginPacket.GeometryData)
			session.SetXBOXLiveAuthenticated(authenticated)

			if server.GetConfiguration().UseEncryption {
				var jwt = utils.ConstructEncryptionJwt(server.GetPrivateKey(), server.GetServerToken())
				session.SendServerHandshake(jwt)

				session.EnableEncryption()
			} else {
				session.SendPlayStatus(data.StatusLoginSuccess)

				session.SendResourcePackInfo(server.GetConfiguration().ForceResourcePacks, server.GetPackManager().GetResourceStack(), server.GetPackManager().GetBehaviorStack())
			}
			server.GetSessionManager().AddMinecraftSession(session)
			return true
		}
		return false
	})
}

func NewMovePlayerHandler_200(_ *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, logger *utils.Logger, session *net.MinecraftSession) bool {
		if pk, ok := packet.(*p200.MovePlayerPacket); ok {
			if session.GetPlayer().GetDimension() == nil {
				return false
			}

			session.SyncMove(pk.Position.X, pk.Position.Y, pk.Position.Z, pk.Rotation.Pitch, pk.Rotation.Yaw, pk.Rotation.HeadYaw, pk.OnGround)

			for _, viewer := range session.GetPlayer().GetViewers() {
				viewer.SendPacket(pk)
			}

			return true
		}

		return false
	})
}

func NewRequestChunkRadiusHandler_200(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, logger *utils.Logger, session *net.MinecraftSession) bool {
		if chunkRadiusPacket, ok := packet.(*p200.RequestChunkRadiusPacket); ok {
			session.SetViewDistance(chunkRadiusPacket.Radius)
			session.SendChunkRadiusUpdated(session.GetViewDistance())

			hasChunks := session.NeedsChunks
			session.NeedsChunks = true
			println(hasChunks)

			if !hasChunks {
				var sessions = server.GetSessionManager().GetSessions()
				var viewers = make(map[string]protocol.PlayerListEntry)
				for name, online := range sessions {
					if online.HasSpawned() {
						viewers[name] = online.GetPlayer()
						online.SendPlayerList(data.ListTypeAdd, map[string]protocol.PlayerListEntry{session.GetName(): session.GetPlayer()})
					}
				}
				session.SendPlayerList(data.ListTypeAdd, viewers)

				for _, online := range server.GetSessionManager().GetSessions() {
					if session.GetUUID() != online.GetUUID() {
						online.GetPlayer().SpawnTo(session)
						online.GetPlayer().SpawnPlayerTo(session)
					}
				}

				session.GetPlayer().SpawnPlayerToAll()
				session.GetPlayer().SpawnToAll()

				session.SendUpdateAttributes(session.GetPlayer().GetRuntimeId(), session.GetPlayer().GetAttributeMap())
				server.BroadcastMessage(utils.Yellow+session.GetDisplayName(), "has joined the server")

				session.SendPlayStatus(data.StatusSpawn)
			}

			return true
		}

		return false
	})
}

func NewResourcePackChunkRequestHandler_200(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, logger *utils.Logger, session *net.MinecraftSession) bool {
		if request, ok := packet.(*p200.ResourcePackChunkRequestPacket); ok {
			if !server.GetPackManager().IsPackLoaded(request.PackUUID) {
				// TODO: Kick the player. We can't kick yet.
				return false
			}
			var pack = server.GetPackManager().GetPack(request.PackUUID)
			session.SendResourcePackChunkData(request.PackUUID, request.ChunkIndex, int64(data.ResourcePackChunkSize*request.ChunkIndex), pack.GetChunk(int(data.ResourcePackChunkSize*request.ChunkIndex), data.ResourcePackChunkSize))
			return true
		}
		return false
	})
}

func NewResourcePackClientResponseHandler_200(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, logger *utils.Logger, session *net.MinecraftSession) bool {
		if response, ok := packet.(*p200.ResourcePackClientResponsePacket); ok {
			switch response.Status {
			case data.StatusRefused:
				// TODO: Kick the player. We can't kick yet.
				return false
			case data.StatusSendPacks:
				for _, packUUID := range response.PackUUIDs {
					if !server.GetPackManager().IsPackLoaded(packUUID) {
						// TODO: Kick the player. We can't kick yet.
						return false
					}
					session.SendResourcePackDataInfo(server.GetPackManager().GetPack(packUUID))
				}
			case data.StatusHaveAllPacks:
				session.SendResourcePackStack(server.GetConfiguration().ForceResourcePacks, server.GetPackManager().GetResourceStack(), server.GetPackManager().GetBehaviorStack())
			case data.StatusCompleted:
				server.GetLevelManager().GetDefaultLevel().GetDefaultDimension().LoadChunk(0, 0, func(chunk *chunks.Chunk) {
					server.GetLevelManager().GetDefaultLevel().GetDefaultDimension().AddEntity(session.GetPlayer(), r3.Vector{X: 0, Y: 40, Z: 0})
					session.SendStartGame(session.GetPlayer())
					session.SendCraftingData()
				})
			}
			return true
		}
		return false
	})
}

func NewTextHandler_200(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, logger *utils.Logger, session *net.MinecraftSession) bool {
		if textPacket, ok := packet.(*p200.TextPacket); ok {
			if textPacket.TextType != data.TextChat {
				return false
			}
			for _, receiver := range server.GetSessionManager().GetSessions() {
				receiver.SendText(types.Text{Message: textPacket.Message, SourceName: textPacket.SourceName, SourceDisplayName: textPacket.SourceDisplayName, SourcePlatform: textPacket.SourcePlatform, SourceXUID: session.GetXUID(), TextType: data.TextChat})
			}
			server.GetLogger().LogChat("<" + session.GetDisplayName() + "> " + textPacket.Message)
			return true
		}
		return false
	})
}

func VerifyLoginRequest(chains []types.Chain, server *Server) (successful bool, authenticated bool, clientPublicKey *ecdsa.PublicKey) {
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
		d := []byte(chain.Header.Raw + "." + chain.Payload.Raw)

		var b64, errB64 = base64.RawStdEncoding.DecodeString(publicKeyRaw)
		server.GetLogger().LogError(errB64)

		key, err := x509.ParsePKIXPublicKey(b64)
		if err != nil {
			server.GetLogger().LogError(err)
			return
		}

		hash := sha512.New384()
		hash.Write(d)

		publicKey = key.(*ecdsa.PublicKey)
		r := new(big.Int).SetBytes(sig[:len(sig)/2])
		s := new(big.Int).SetBytes(sig[len(sig)/2:])

		if !ecdsa.Verify(publicKey, hash.Sum(nil), r, s) {
			return
		}

		if publicKeyRaw == data.MojangPublicKey {
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
