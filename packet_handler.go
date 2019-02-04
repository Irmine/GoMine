package gomine

import (
	"crypto/ecdsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/bedrock"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/players"
	"github.com/irmine/gomine/text"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/worlds/chunks"
	"gopkg.in/yaml.v2"
	"math/big"
	"strings"
	"time"
)

var runtimeIdsTable []byte

func NewClientHandshakeHandler(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if _, ok := packet.(*bedrock.ClientHandshakePacket); ok {
			session.SendPlayStatus(data.StatusLoginSuccess)
			session.SendResourcePackInfo(server.Config.ForceResourcePacks, server.PackManager.GetResourceStack(), server.PackManager.GetBehaviorStack())
			return true
		}
		return false
	})
}

func NewCommandRequestHandler(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if pk, ok := packet.(*bedrock.CommandRequestPacket); ok {
			var args = strings.Split(pk.CommandText, " ")
			var commandName = strings.TrimLeft(args[0], "/")
			var i = 1
			for !server.CommandManager.IsCommandRegistered(commandName) {
				if i == len(args) {
					break
				}
				commandName += " " + args[i]
				i++
			}
			if !server.CommandManager.IsCommandRegistered(commandName) {
				session.SendMessage("Command could not be found.")
				return false
			}
			args = args[i:]
			var command, _ = server.CommandManager.GetCommand(commandName)
			command.Execute(session, args)

			return true
		}

		return false
	})
}

func NewLoginHandler(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if loginPacket, ok := packet.(*bedrock.LoginPacket); ok {
			var _, ok = server.SessionManager.GetSession(loginPacket.Username)
			if ok {
				return false
			}

			if loginPacket.Protocol > info.LatestProtocol {
				session.Kick("Outdated server.", false, true)
				return false
			}

			if loginPacket.Protocol < info.LatestProtocol {
				session.Kick("Outdated client.", false, true)
				return false
			}

			var successful, authenticated, pubKey = VerifyLoginRequest(loginPacket.Chains, server)

			if !successful {
				text.DefaultLogger.Debug(loginPacket.Username, "has joined with invalid login data.")
				return true
			}

			if authenticated {
				text.DefaultLogger.Debug(loginPacket.Username, "has joined while being logged into XBOX Live.")
			} else {
				if server.Config.XBOXLiveAuth {
					text.DefaultLogger.Debug(loginPacket.Username, "has tried to join while not being logged into XBOX Live.")
					session.Kick("XBOX Live account required.", false, false)
					return true
				}
				text.DefaultLogger.Debug(loginPacket.Username, "has joined while not being logged into XBOX Live.")
			}

			session.SetData(server.PermissionManager, types.SessionData{ClientUUID: loginPacket.ClientUUID, ClientXUID: loginPacket.ClientXUID, ClientId: loginPacket.ClientId, ProtocolNumber: loginPacket.Protocol, GameVersion: loginPacket.ClientData.GameVersion, Language: loginPacket.Language, DeviceOS: loginPacket.ClientData.DeviceOS})
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

			if server.Config.UseEncryption {
				var jwt = utils.ConstructEncryptionJwt(server.GetPrivateKey(), server.GetServerToken())
				session.SendServerHandshake(jwt)
				session.EnableEncryption()
			} else {
				session.SendPlayStatus(data.StatusLoginSuccess)
				session.SendResourcePackInfo(server.Config.ForceResourcePacks, server.PackManager.GetResourceStack(), server.PackManager.GetBehaviorStack())
			}

			server.SessionManager.AddMinecraftSession(session)
			return true
		}
		return false
	})
}

func NewMovePlayerHandler(_ *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if pk, ok := packet.(*bedrock.MovePlayerPacket); ok {
			if session.GetPlayer().GetDimension() == nil {
				return false
			}
			session.SyncMove(pk.Position.X, pk.Position.Y, pk.Position.Z, pk.Rotation.Pitch, pk.Rotation.Yaw, pk.Rotation.HeadYaw, pk.OnGround)
			return true
		}
		return false
	})
}

func NewRequestChunkRadiusHandler(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if chunkRadiusPacket, ok := packet.(*bedrock.RequestChunkRadiusPacket); ok {
			session.SetViewDistance(chunkRadiusPacket.Radius)
			session.SendChunkRadiusUpdated(session.GetViewDistance())

			session.Connected = true

			var sessions = server.SessionManager.GetSessions()
			var viewers = make(map[string]protocol.PlayerListEntry)
			for name, online := range sessions {
				if online.HasSpawned() {
					viewers[name] = online.GetPlayer()
					online.SendPlayerList(data.ListTypeAdd, map[string]protocol.PlayerListEntry{session.GetName(): session.GetPlayer()})
				}
			}

			session.SendPlayerList(data.ListTypeAdd, viewers)

			for _, online := range server.SessionManager.GetSessions() {
				if session.GetUUID() != online.GetUUID() {
					online.GetPlayer().SpawnPlayerTo(session)
					online.GetPlayer().AddViewer(session)

					session.GetPlayer().SpawnPlayerTo(online)
					session.GetPlayer().AddViewer(online)

					online.SendSkin(session)
					session.SendSkin(online)
				}
			}

			session.SendSetEntityData(session.GetPlayer().GetRuntimeId(), session.GetPlayer().GetEntityData())
			session.SendUpdateAttributes(session.GetPlayer().GetRuntimeId(), session.GetPlayer().GetAttributeMap())

			server.BroadcastMessage(text.Yellow+session.GetDisplayName(), "has joined the server")
			session.SendPlayStatus(data.StatusSpawn)

			return true
		}

		return false
	})
}

func NewResourcePackChunkRequestHandler(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if request, ok := packet.(*bedrock.ResourcePackChunkRequestPacket); ok {
			if !server.PackManager.IsPackLoaded(request.PackUUID) {
				// TODO: Kick the player. We can't kick yet.
				return false
			}
			var pack = server.PackManager.GetPack(request.PackUUID)
			session.SendResourcePackChunkData(request.PackUUID, request.ChunkIndex, int64(data.ResourcePackChunkSize*request.ChunkIndex), pack.GetChunk(int(data.ResourcePackChunkSize*request.ChunkIndex), data.ResourcePackChunkSize))
			return true
		}
		return false
	})
}

func NewResourcePackClientResponseHandler(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if response, ok := packet.(*bedrock.ResourcePackClientResponsePacket); ok {
			switch response.Status {
			case data.StatusRefused:
				// TODO: Kick the player. We can't kick yet.
				return false
			case data.StatusSendPacks:
				for _, packUUID := range response.PackUUIDs {
					if !server.PackManager.IsPackLoaded(packUUID) {
						// TODO: Kick the player. We can't kick yet.
						return false
					}
					session.SendResourcePackDataInfo(server.PackManager.GetPack(packUUID))
				}
			case data.StatusHaveAllPacks:
				session.SendResourcePackStack(server.Config.ForceResourcePacks, server.PackManager.GetResourceStack(), server.PackManager.GetBehaviorStack())
			case data.StatusCompleted:
				server.LevelManager.GetDefaultLevel().GetDefaultDimension().LoadChunk(0, 0, func(chunk *chunks.Chunk) {
					server.LevelManager.GetDefaultLevel().GetDefaultDimension().AddEntity(session.GetPlayer(), r3.Vector{X: 0, Y: 7, Z: 0})
					session.SendStartGame(session.GetPlayer(), GetRuntimeIdsTable())
					session.SendCraftingData()
				})
			}
			return true
		}
		return false
	})
}

func NewTextHandler(server *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if textPacket, ok := packet.(*bedrock.TextPacket); ok {
			if textPacket.TextType != data.TextChat {
				return false
			}
			for _, receiver := range server.SessionManager.GetSessions() {
				receiver.SendText(types.Text{
					Message: "<" + session.GetDisplayName() + "> " + textPacket.Message,
					PlatformChatId: textPacket.PlatformChatId,
					SourceXUID: session.GetXUID(),
					TextType: data.TextChat,
				})
			}
			text.DefaultLogger.LogChat("<" + session.GetDisplayName() + "> " + textPacket.Message)
			return true
		}
		return false
	})
}

func NewInteractHandler(_ *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if /*interactPacket*/ _, ok := packet.(*bedrock.InteractPacket); ok {
		}
		return true
	})
}

func NewSetEntityDataHandler(_ *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if /*entityData*/ _, ok := packet.(*bedrock.SetEntityDataPacket); ok {
		}
		return true
	})
}

func NewPlayerActionHandler(_ *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		//TODO: fix sending to others
		//if playerAction, ok := packet.(*bedrock.PlayerActionPacket); ok {
		//	switch playerAction.Action {
		//	case bedrock.PlayerStartSneak:
		//		session.GetPlayer().SetEntityProperty(data2.EntityDataIdFlags, data2.EntityDataSneaking, data2.EntityDataLong, true)
		//		for _, viewer := range session.GetPlayer().GetViewers() {
		//			if viewer, ok := viewer.(*net.MinecraftSession); ok {
		//				viewer.SendSetEntityData(session.GetPlayer().GetRuntimeId(), session.GetPlayer().GetEntityData())
		//			}
		//		}
		//		break
		//	}
		//}
		//return true
		return true
	})
}

func NewAnimateHandler(_ *Server) *net.PacketHandler {
	return net.NewPacketHandler(func(packet packets.IPacket, session *net.MinecraftSession) bool {
		if animate, ok := packet.(*bedrock.AnimatePacket); ok {
			for _, viewer := range session.GetPlayer().GetViewers() {
				if viewer, ok := viewer.(*net.MinecraftSession); ok {
					viewer.SendAnimate(animate.Action, animate.RuntimeId, animate.Float)
				}
			}
		}
		return true
	})
}

func VerifyLoginRequest(chains []types.Chain, _ *Server) (successful bool, authenticated bool, clientPublicKey *ecdsa.PublicKey) {
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
		text.DefaultLogger.LogError(errB64)

		key, err := x509.ParsePKIXPublicKey(b64)
		if err != nil {
			text.DefaultLogger.LogError(err)
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
		if chain.Payload.ExpirationTime <= t && chain.Payload.ExpirationTime != 0 || chain.Payload.NotBefore > t || chain.Payload.IssuedAt > chain.Payload.ExpirationTime {
			return
		}

		publicKeyRaw = chain.Payload.IdentityPublicKey
	}

	var b64, errB64 = base64.RawStdEncoding.DecodeString(publicKeyRaw)
	text.DefaultLogger.LogError(errB64)

	key, err := x509.ParsePKIXPublicKey(b64)
	if err != nil {
		text.DefaultLogger.LogError(err)
		return
	}

	clientPublicKey = key.(*ecdsa.PublicKey)

	successful = true
	return
}

func GetRuntimeIdsTable() []byte {
	if len(runtimeIdsTable) == 0 {
		var data2 interface{}
		err := yaml.Unmarshal([]byte(utils.RuntimeIdsTable__), &data2)
		if err != nil {
			text.DefaultLogger.Error(err)
			return nil
		}
		stream := packets.NewMinecraftStream()
		stream.ResetStream()
		if data2, ok := data2.([]interface{}); ok {
			stream.PutUnsignedVarInt(uint32(len(data2)))
			for _, v := range data2 {
				if v2, ok := v.(map[interface{}]interface{}); ok {
					stream.PutString(v2["name"].(string))
					stream.PutLittleShort(int16(v2["data"].(int)))
				}
			}
		}
		runtimeIdsTable = stream.GetBuffer()
	}
	return runtimeIdsTable
}