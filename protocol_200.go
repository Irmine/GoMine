package gomine

import (
	"math"

	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/packs"
	"github.com/irmine/gomine/permissions"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/worlds/chunks"
	data2 "github.com/irmine/worlds/entities/data"
)

type P200 struct {
	*protocol.Base
}

func NewP200(server *Server) *P200 {
	var ids = info.PacketIds200
	var proto = &P200{protocol.NewBase(200, info.PacketIds200, map[int]func() packets.IPacket{
		ids[info.LoginPacket]:                      func() packets.IPacket { return p200.NewLoginPacket() },
		ids[info.ClientHandshakePacket]:            func() packets.IPacket { return p200.NewClientHandshakePacket() },
		ids[info.ResourcePackClientResponsePacket]: func() packets.IPacket { return p200.NewResourcePackClientResponsePacket() },
		ids[info.RequestChunkRadiusPacket]:         func() packets.IPacket { return p200.NewRequestChunkRadiusPacket() },
		ids[info.MovePlayerPacket]:                 func() packets.IPacket { return p200.NewMovePlayerPacket() },
		ids[info.CommandRequestPacket]:             func() packets.IPacket { return p200.NewCommandRequestPacket() },
		ids[info.ResourcePackChunkRequestPacket]:   func() packets.IPacket { return p200.NewResourcePackChunkRequestPacket() },
		ids[info.TextPacket]:                       func() packets.IPacket { return p200.NewTextPacket() },
		ids[info.PlayerListPacket]:                 func() packets.IPacket { return p200.NewPlayerListPacket() },
	}, map[int][][]protocol.Handler{})}
	proto.initHandlers(server)

	return proto
}

func (protocol *P200) initHandlers(server *Server) {
	protocol.RegisterHandler(info.LoginPacket, NewLoginHandler_200(server))
	protocol.RegisterHandler(info.ClientHandshakePacket, NewClientHandshakeHandler_200(server))
	protocol.RegisterHandler(info.RequestChunkRadiusPacket, NewRequestChunkRadiusHandler_200(server))
	protocol.RegisterHandler(info.ResourcePackClientResponsePacket, NewResourcePackClientResponseHandler_200(server))
	protocol.RegisterHandler(info.MovePlayerPacket, NewMovePlayerHandler_200(server))
	protocol.RegisterHandler(info.CommandRequestPacket, NewCommandRequestHandler_200(server))
	protocol.RegisterHandler(info.ResourcePackChunkRequestPacket, NewResourcePackChunkRequestHandler_200(server))
	protocol.RegisterHandler(info.TextPacket, NewTextHandler_200(server))
}

func (protocol *P200) GetAddEntity(entity protocol.AddEntityEntry) packets.IPacket {
	var pk = p200.NewAddEntityPacket()
	pk.UniqueId = entity.GetUniqueId()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.EntityType = entity.GetEntityType()
	pk.Position = entity.GetPosition()
	pk.Motion = entity.GetMotion()
	pk.Rotation = entity.GetRotation()
	pk.Attributes = entity.GetAttributeMap()
	pk.EntityData = entity.GetEntityData()

	return pk
}

func (protocol *P200) GetAddPlayer(uuid utils.UUID, platform int32, player protocol.AddPlayerEntry) packets.IPacket {
	var pk = p200.NewAddPlayerPacket()
	pk.UUID = uuid
	pk.DisplayName = player.GetDisplayName()
	pk.Username = player.GetName()
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.Position = player.GetPosition()
	pk.Rotation = player.GetRotation()
	pk.Platform = platform
	pk.Motion = player.GetMotion()

	return pk
}

func (protocol *P200) GetChunkRadiusUpdated(radius int32) packets.IPacket {
	var pk = p200.NewChunkRadiusUpdatedPacket()
	pk.Radius = radius

	return pk
}

func (protocol *P200) GetCraftingData() packets.IPacket {
	var pk = p200.NewCraftingDataPacket()

	return pk
}

func (protocol *P200) GetDisconnect(message string, hideDisconnectScreen bool) packets.IPacket {
	var pk = p200.NewDisconnectPacket()
	pk.HideDisconnectionScreen = hideDisconnectScreen
	pk.Message = message

	return pk
}

func (protocol *P200) GetFullChunkData(chunk *chunks.Chunk) packets.IPacket {
	var pk = p200.NewFullChunkDataPacket()
	pk.ChunkX, pk.ChunkZ = chunk.X, chunk.Z
	pk.ChunkData = chunk.ToBinary()

	return pk
}

func (protocol *P200) GetMovePlayer(runtimeId uint64, position r3.Vector, rotation data2.Rotation, mode byte, onGround bool, ridingRuntimeId uint64) packets.IPacket {
	var pk = p200.NewMovePlayerPacket()
	pk.RuntimeId = runtimeId
	pk.Position = position
	pk.Rotation = rotation
	pk.Mode = mode
	pk.OnGround = onGround
	pk.RidingRuntimeId = ridingRuntimeId

	return pk
}

func (protocol *P200) GetPlayerList(listType byte, players map[string]protocol.PlayerListEntry) packets.IPacket {
	var pk = p200.NewPlayerListPacket()
	pk.ListType = listType
	var entries = map[string]types.PlayerListEntry{}
	for name, player := range players {
		entries[name] = types.PlayerListEntry{
			UUID:           player.GetUUID(),
			XUID:           player.GetXUID(),
			EntityUniqueId: player.GetUniqueId(),
			Username:       player.GetName(),
			DisplayName:    player.GetDisplayName(),
			Platform:       player.GetPlatform(),
			SkinId:         player.GetSkinId(),
			SkinData:       player.GetSkinData(),
			CapeData:       player.GetCapeData(),
			GeometryName:   player.GetGeometryName(),
			GeometryData:   player.GetGeometryData(),
		}
	}
	pk.Entries = entries

	return pk
}

func (protocol *P200) GetPlayStatus(status int32) packets.IPacket {
	var pk = p200.NewPlayStatusPacket()
	pk.Status = status

	return pk
}

func (protocol *P200) GetRemoveEntity(uniqueId int64) packets.IPacket {
	var pk = p200.NewRemoveEntityPacket()
	pk.EntityUniqueId = uniqueId

	return pk
}

func (protocol *P200) GetResourcePackChunkData(packUUID string, chunkIndex int32, progress int64, data []byte) packets.IPacket {
	var pk = p200.NewResourcePackChunkDataPacket()
	pk.PackUUID = packUUID
	pk.ChunkIndex = chunkIndex
	pk.Progress = progress
	pk.ChunkData = data

	return pk
}

func (protocol *P200) GetResourcePackDataInfo(pack packs.Pack) packets.IPacket {
	var pk = p200.NewResourcePackDataInfoPacket()
	pk.PackUUID = pack.GetUUID()
	pk.MaxChunkSize = data.ResourcePackChunkSize
	pk.ChunkCount = int32(math.Ceil(float64(pack.GetFileSize()) / float64(data.ResourcePackChunkSize)))
	pk.CompressedPackSize = pack.GetFileSize()
	pk.Sha256 = pack.GetSha256()

	return pk
}

func (protocol *P200) GetResourcePackInfo(mustAccept bool, resourcePacks []packs.Pack, behaviorPacks []packs.Pack) packets.IPacket {
	var pk = p200.NewResourcePackInfoPacket()
	pk.MustAccept = mustAccept

	var resourceEntries []types.ResourcePackInfoEntry
	var behaviorEntries []types.ResourcePackInfoEntry
	for _, pack := range resourcePacks {
		resourceEntries = append(resourceEntries, types.ResourcePackInfoEntry{
			UUID:     pack.GetUUID(),
			Version:  pack.GetVersion(),
			PackSize: pack.GetFileSize(),
		})
	}
	for _, pack := range behaviorPacks {
		behaviorEntries = append(behaviorEntries, types.ResourcePackInfoEntry{
			UUID:     pack.GetUUID(),
			Version:  pack.GetVersion(),
			PackSize: pack.GetFileSize(),
		})
	}

	pk.ResourcePacks = resourceEntries
	pk.BehaviorPacks = behaviorEntries

	return pk
}

func (protocol *P200) GetResourcePackStack(mustAccept bool, resourcePacks []packs.Pack, behaviorPacks []packs.Pack) packets.IPacket {
	var pk = p200.NewResourcePackStackPacket()
	pk.MustAccept = mustAccept
	var resourceEntries []types.ResourcePackStackEntry
	var behaviorEntries []types.ResourcePackStackEntry
	for _, pack := range resourcePacks {
		resourceEntries = append(resourceEntries, types.ResourcePackStackEntry{
			UUID:    pack.GetUUID(),
			Version: pack.GetVersion(),
		})
	}
	for _, pack := range behaviorPacks {
		behaviorEntries = append(behaviorEntries, types.ResourcePackStackEntry{
			UUID:    pack.GetUUID(),
			Version: pack.GetVersion(),
		})
	}

	pk.ResourcePacks = resourceEntries
	pk.BehaviorPacks = behaviorEntries

	return pk
}

func (protocol *P200) GetServerHandshake(encryptionJwt string) packets.IPacket {
	var pk = p200.NewServerHandshakePacket()
	pk.Jwt = encryptionJwt

	return pk
}

func (protocol *P200) GetSetEntityData(runtimeId uint64, data map[uint32][]interface{}) packets.IPacket {
	var pk = p200.NewSetEntityDataPacket()
	pk.RuntimeId = runtimeId
	pk.EntityData = data

	return pk
}

func (protocol *P200) GetStartGame(player protocol.StartGameEntry) packets.IPacket {
	var pk = p200.NewStartGamePacket()
	pk.Generator = 1
	pk.LevelSeed = 312402
	pk.TrustPlayers = true
	pk.DefaultPermissionLevel = permissions.LevelMember
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.PlayerGameMode = 1
	pk.PlayerPosition = player.GetPosition()
	pk.LevelGameMode = 1
	pk.LevelSpawnPosition = r3.Vector{0, 40, 0}
	pk.CommandsEnabled = true

	var gameRules = player.GetLevel().GetGameRules()
	var gameRuleEntries = map[string]types.GameRuleEntry{}
	for name, gameRule := range gameRules {
		gameRuleEntries[string(name)] = types.GameRuleEntry{Name: string(gameRule.GetName()), Value: gameRule.GetValue()}
	}

	pk.GameRules = gameRuleEntries
	pk.LevelName = player.GetLevel().GetName()
	pk.CurrentTick = player.GetLevel().GetCurrentTick()
	pk.Time = 0
	pk.AchievementsDisabled = true
	pk.BroadcastToXbox = true
	pk.BroadcastToLan = true

	return pk
}

func (protocol *P200) GetText(text types.Text) packets.IPacket {
	var pk = p200.NewTextPacket()
	pk.TextType = text.TextType
	pk.IsTranslation = text.IsTranslation
	pk.TranslationParameters = text.TranslationParameters
	pk.SourceName = text.SourceName
	pk.SourceDisplayName = text.SourceDisplayName
	pk.SourcePlatform = text.SourcePlatform
	pk.XUID = text.SourceXUID
	pk.Message = text.Message

	return pk
}

func (protocol *P200) GetTransfer(address string, port uint16) packets.IPacket {
	var pk = p200.NewTransferPacket()
	pk.Address = address
	pk.Port = port

	return pk
}

func (protocol *P200) GetUpdateAttributes(runtimeId uint64, attributeMap data2.AttributeMap) packets.IPacket {
	var pk = p200.NewUpdateAttributesPacket()
	pk.RuntimeId = runtimeId
	pk.Attributes = attributeMap

	return pk
}
