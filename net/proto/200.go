package proto

import (
	math2 "math"

	data2 "github.com/irmine/gomine/entities/data"
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/p200"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/permissions"
	p200handlers "github.com/irmine/gomine/players/handlers/p200"
	"github.com/irmine/gomine/packs"
	"github.com/golang/geo/r3"
)

type Protocol200 struct {
	*Protocol
}

func NewProtocol200() *Protocol200 {
	var ids = info.PacketIds200
	var proto = &Protocol200{NewProtocol(200, info.PacketIds200, map[int]func() interfaces.IPacket{
		ids[info.LoginPacket]:                      func() interfaces.IPacket { return p200.NewLoginPacket() },
		ids[info.ClientHandshakePacket]:            func() interfaces.IPacket { return p200.NewClientHandshakePacket() },
		ids[info.ResourcePackClientResponsePacket]: func() interfaces.IPacket { return p200.NewResourcePackClientResponsePacket() },
		ids[info.RequestChunkRadiusPacket]:         func() interfaces.IPacket { return p200.NewRequestChunkRadiusPacket() },
		ids[info.MovePlayerPacket]:                 func() interfaces.IPacket { return p200.NewMovePlayerPacket() },
		ids[info.CommandRequestPacket]:             func() interfaces.IPacket { return p200.NewCommandRequestPacket() },
		ids[info.ResourcePackChunkRequestPacket]:   func() interfaces.IPacket { return p200.NewResourcePackChunkRequestPacket() },
		ids[info.TextPacket]:                       func() interfaces.IPacket { return p200.NewTextPacket() },
		ids[info.PlayerListPacket]:                 func() interfaces.IPacket { return p200.NewPlayerListPacket() },
	}, map[int][][]interfaces.IPacketHandler{})}
	proto.initHandlers()

	return proto
}

func (protocol *Protocol200) initHandlers() {
	protocol.RegisterHandler(info.LoginPacket, p200handlers.NewLoginHandler(), 8)
	protocol.RegisterHandler(info.ClientHandshakePacket, p200handlers.NewClientHandshakeHandler(), 8)
	protocol.RegisterHandler(info.RequestChunkRadiusPacket, p200handlers.NewRequestChunkRadiusHandler(), 8)
	protocol.RegisterHandler(info.ResourcePackClientResponsePacket, p200handlers.NewResourcePackClientResponseHandler(), 8)
	protocol.RegisterHandler(info.MovePlayerPacket, p200handlers.NewMovePlayerHandler(), 8)
	protocol.RegisterHandler(info.CommandRequestPacket, p200handlers.NewCommandRequestHandler(), 8)
	protocol.RegisterHandler(info.ResourcePackChunkRequestPacket, p200handlers.NewResourcePackChunkRequestHandler(), 8)
	protocol.RegisterHandler(info.TextPacket, p200handlers.NewTextHandler(), 8)
}

func (protocol *Protocol200) GetAddEntity(entity interfaces.IEntity) interfaces.IPacket {
	var pk = p200.NewAddEntityPacket()
	pk.UniqueId = entity.GetUniqueId()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.EntityType = entity.GetEntityId()
	pk.Position = entity.GetPosition()
	pk.Motion = entity.GetMotion()
	pk.Rotation = *entity.GetRotation()
	pk.Attributes = entity.GetAttributeMap()
	pk.EntityData = entity.GetEntityData()

	return pk
}

func (protocol *Protocol200) GetAddPlayer(player interfaces.IPlayer) interfaces.IPacket {
	var pk = p200.NewAddPlayerPacket()
	pk.UUID = player.GetUUID()
	pk.DisplayName = player.GetDisplayName()
	pk.Username = player.GetName()
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.Position = player.GetPosition()
	pk.Rotation = *player.GetRotation()
	pk.Platform = player.GetPlatform()

	return pk
}

func (protocol *Protocol200) GetChunkRadiusUpdated(radius int32) interfaces.IPacket {
	var pk = p200.NewChunkRadiusUpdatedPacket()
	pk.Radius = radius

	return pk
}

func (protocol *Protocol200) GetCraftingData() interfaces.IPacket {
	var pk = p200.NewCraftingDataPacket()

	return pk
}

func (protocol *Protocol200) GetDisconnect(message string, hideDisconnectScreen bool) interfaces.IPacket {
	var pk = p200.NewDisconnectPacket()
	pk.HideDisconnectionScreen = hideDisconnectScreen
	pk.Message = message

	return pk
}

func (protocol *Protocol200) GetFullChunkData(chunk interfaces.IChunk) interfaces.IPacket {
	var pk = p200.NewFullChunkDataPacket()
	pk.ChunkX = chunk.GetX()
	pk.ChunkZ = chunk.GetZ()
	pk.ChunkData = chunk.ToBinary()

	return pk
}

func (protocol *Protocol200) GetMovePlayer(runtimeId uint64, position r3.Vector, rotation math.Rotation, mode byte, onGround bool, ridingRuntimeId uint64) interfaces.IPacket {
	var pk = p200.NewMovePlayerPacket()
	pk.RuntimeId = runtimeId
	pk.Position = position
	pk.Rotation = rotation
	pk.Mode = mode
	pk.OnGround = onGround
	pk.RidingRuntimeId = ridingRuntimeId

	return pk
}

func (protocol *Protocol200) GetPlayerList(listType byte, players map[string]interfaces.IPlayer) interfaces.IPacket {
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

func (protocol *Protocol200) GetPlayStatus(status int32) interfaces.IPacket {
	var pk = p200.NewPlayStatusPacket()
	pk.Status = status

	return pk
}

func (protocol *Protocol200) GetRemoveEntity(uniqueId int64) interfaces.IPacket {
	var pk = p200.NewRemoveEntityPacket()
	pk.EntityUniqueId = uniqueId

	return pk
}

func (protocol *Protocol200) GetResourcePackChunkData(packUUID string, chunkIndex int32, progress int64, data []byte) interfaces.IPacket {
	var pk = p200.NewResourcePackChunkDataPacket()
	pk.PackUUID = packUUID
	pk.ChunkIndex = chunkIndex
	pk.Progress = progress
	pk.ChunkData = data

	return pk
}

func (protocol *Protocol200) GetResourcePackDataInfo(pack packs.Pack) interfaces.IPacket {
	var pk = p200.NewResourcePackDataInfoPacket()
	pk.PackUUID = pack.GetUUID()
	pk.MaxChunkSize = data.ResourcePackChunkSize
	pk.ChunkCount = int32(math2.Ceil(float64(pack.GetFileSize()) / float64(data.ResourcePackChunkSize)))
	pk.CompressedPackSize = pack.GetFileSize()
	pk.Sha256 = pack.GetSha256()

	return pk
}

func (protocol *Protocol200) GetResourcePackInfo(mustAccept bool, resourcePacks []packs.Pack, behaviorPacks []packs.Pack) interfaces.IPacket {
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

func (protocol *Protocol200) GetResourcePackStack(mustAccept bool, resourcePacks []packs.Pack, behaviorPacks []packs.Pack) interfaces.IPacket {
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

func (protocol *Protocol200) GetServerHandshake(encryptionJwt string) interfaces.IPacket {
	var pk = p200.NewServerHandshakePacket()
	pk.Jwt = encryptionJwt

	return pk
}

func (protocol *Protocol200) GetSetEntityData(entity interfaces.IEntity, data map[uint32][]interface{}) interfaces.IPacket {
	var pk = p200.NewSetEntityDataPacket()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.EntityData = data

	return pk
}

func (protocol *Protocol200) GetStartGame(player interfaces.IPlayer) interfaces.IPacket {
	var pk = p200.NewStartGamePacket()
	pk.Generator = 1
	pk.LevelSeed = 312402
	pk.TrustPlayers = true
	pk.DefaultPermissionLevel = permissions.LevelMember
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.PlayerGameMode = 1
	pk.PlayerPosition = r3.Vector{0, 40, 0}
	pk.LevelGameMode = 1
	pk.LevelSpawnPosition = r3.Vector{0, 40, 0}
	pk.CommandsEnabled = true

	var gameRules = player.GetServer().GetDefaultLevel().GetGameRules()
	var gameRuleEntries = map[string]types.GameRuleEntry{}
	for name, gameRule := range gameRules {
		gameRuleEntries[name] = types.GameRuleEntry{Name: gameRule.GetName(), Value: gameRule.GetValue()}
	}

	pk.GameRules = gameRuleEntries
	pk.LevelName = player.GetServer().GetDefaultLevel().GetName()
	pk.CurrentTick = player.GetServer().GetCurrentTick()
	pk.Time = 0
	pk.AchievementsDisabled = true
	pk.BroadcastToXbox = true
	pk.BroadcastToLan = true

	return pk
}

func (protocol *Protocol200) GetText(text types.Text) interfaces.IPacket {
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

func (protocol *Protocol200) GetTransfer(address string, port uint16) interfaces.IPacket {
	var pk = p200.NewTransferPacket()
	pk.Address = address
	pk.Port = port

	return pk
}

func (protocol *Protocol200) GetUpdateAttributes(entity interfaces.IEntity, attributeMap *data2.AttributeMap) interfaces.IPacket {
	var pk = p200.NewUpdateAttributesPacket()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.Attributes = attributeMap

	return pk
}
