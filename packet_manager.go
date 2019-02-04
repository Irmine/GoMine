package gomine

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/bedrock"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/packs"
	"github.com/irmine/gomine/permissions"
	"github.com/irmine/worlds/blocks"
	"github.com/irmine/worlds/chunks"
	data2 "github.com/irmine/worlds/entities/data"
	"math"
)

type PacketManager struct {
	*protocol.PacketManagerBase
}

func NewPacketManager(server *Server) *PacketManager {
	var ids = info.PacketIds
	var proto = &PacketManager{protocol.NewPacketManagerBase(info.PacketIds, map[int]func() packets.IPacket{
		ids[info.LoginPacket]:                      func() packets.IPacket { return bedrock.NewLoginPacket() },
		ids[info.ClientHandshakePacket]:            func() packets.IPacket { return bedrock.NewClientHandshakePacket() },
		ids[info.ResourcePackClientResponsePacket]: func() packets.IPacket { return bedrock.NewResourcePackClientResponsePacket() },
		ids[info.RequestChunkRadiusPacket]:         func() packets.IPacket { return bedrock.NewRequestChunkRadiusPacket() },
		ids[info.MovePlayerPacket]:                 func() packets.IPacket { return bedrock.NewMovePlayerPacket() },
		ids[info.CommandRequestPacket]:             func() packets.IPacket { return bedrock.NewCommandRequestPacket() },
		ids[info.ResourcePackChunkRequestPacket]:   func() packets.IPacket { return bedrock.NewResourcePackChunkRequestPacket() },
		ids[info.TextPacket]:                       func() packets.IPacket { return bedrock.NewTextPacket() },
		ids[info.PlayerListPacket]:                 func() packets.IPacket { return bedrock.NewPlayerListPacket() },
		ids[info.InteractPacket]:                   func() packets.IPacket { return bedrock.NewInteractPacket() },
		ids[info.SetEntityDataPacket]:              func() packets.IPacket { return bedrock.NewSetEntityDataPacket() },
		ids[info.PlayerActionPacket]:               func() packets.IPacket { return bedrock.NewPlayerActionPacket() },
		ids[info.AnimatePacket]:                    func() packets.IPacket { return bedrock.NewAnimatePacket() },
	}, map[int][][]protocol.Handler{})}
	proto.initHandlers(server)

	return proto
}

func (protocol *PacketManager) initHandlers(server *Server) {
	protocol.RegisterHandler(info.LoginPacket, NewLoginHandler(server))
	protocol.RegisterHandler(info.ClientHandshakePacket, NewClientHandshakeHandler(server))
	protocol.RegisterHandler(info.RequestChunkRadiusPacket, NewRequestChunkRadiusHandler(server))
	protocol.RegisterHandler(info.ResourcePackClientResponsePacket, NewResourcePackClientResponseHandler(server))
	protocol.RegisterHandler(info.MovePlayerPacket, NewMovePlayerHandler(server))
	protocol.RegisterHandler(info.CommandRequestPacket, NewCommandRequestHandler(server))
	protocol.RegisterHandler(info.ResourcePackChunkRequestPacket, NewResourcePackChunkRequestHandler(server))
	protocol.RegisterHandler(info.TextPacket, NewTextHandler(server))
	protocol.RegisterHandler(info.InteractPacket, NewInteractHandler(server))
	protocol.RegisterHandler(info.SetEntityDataPacket, NewSetEntityDataHandler(server))
	protocol.RegisterHandler(info.PlayerActionPacket, NewPlayerActionHandler(server))
	protocol.RegisterHandler(info.AnimatePacket, NewAnimateHandler(server))
}

func (protocol *PacketManager) GetAddEntity(entity protocol.AddEntityEntry) packets.IPacket {
	var pk = bedrock.NewAddEntityPacket()
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

func (protocol *PacketManager) GetAddPlayer(uuid uuid.UUID, player protocol.AddPlayerEntry) packets.IPacket {
	var pk = bedrock.NewAddPlayerPacket()
	pk.UUID = uuid
	pk.Username = player.GetName()
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.Position = player.GetPosition()
	pk.Rotation = player.GetRotation()
	pk.Motion = player.GetMotion()

	return pk
}

func (protocol *PacketManager) GetChunkRadiusUpdated(radius int32) packets.IPacket {
	var pk = bedrock.NewChunkRadiusUpdatedPacket()
	pk.Radius = radius

	return pk
}

func (protocol *PacketManager) GetCraftingData() packets.IPacket {
	var pk = bedrock.NewCraftingDataPacket()

	return pk
}

func (protocol *PacketManager) GetDisconnect(message string, hideDisconnectScreen bool) packets.IPacket {
	var pk = bedrock.NewDisconnectPacket()
	pk.HideDisconnectionScreen = hideDisconnectScreen
	pk.Message = message

	return pk
}

func (protocol *PacketManager) GetFullChunkData(chunk *chunks.Chunk) packets.IPacket {
	var pk = bedrock.NewFullChunkDataPacket()
	pk.ChunkX, pk.ChunkZ = chunk.X, chunk.Z
	pk.ChunkData = chunk.ToBinary()
	return pk
}

func (protocol *PacketManager) GetMovePlayer(runtimeId uint64, position r3.Vector, rotation data2.Rotation, mode byte, onGround bool, ridingRuntimeId uint64) packets.IPacket {
	var pk = bedrock.NewMovePlayerPacket()
	pk.RuntimeId = runtimeId
	pk.Position = position
	pk.Rotation = rotation
	pk.Mode = mode
	pk.OnGround = onGround
	pk.RidingRuntimeId = ridingRuntimeId

	return pk
}

func (protocol *PacketManager) GetPlayerList(listType byte, players map[string]protocol.PlayerListEntry) packets.IPacket {
	var pk = bedrock.NewPlayerListPacket()
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

func (protocol *PacketManager) GetPlayStatus(status int32) packets.IPacket {
	var pk = bedrock.NewPlayStatusPacket()
	pk.Status = status

	return pk
}

func (protocol *PacketManager) GetRemoveEntity(uniqueId int64) packets.IPacket {
	var pk = bedrock.NewRemoveEntityPacket()
	pk.EntityUniqueId = uniqueId

	return pk
}

func (protocol *PacketManager) GetResourcePackChunkData(packUUID string, chunkIndex int32, progress int64, data []byte) packets.IPacket {
	var pk = bedrock.NewResourcePackChunkDataPacket()
	pk.PackUUID = packUUID
	pk.ChunkIndex = chunkIndex
	pk.Progress = progress
	pk.ChunkData = data

	return pk
}

func (protocol *PacketManager) GetResourcePackDataInfo(pack packs.Pack) packets.IPacket {
	var pk = bedrock.NewResourcePackDataInfoPacket()
	pk.PackUUID = pack.GetUUID()
	pk.MaxChunkSize = data.ResourcePackChunkSize
	pk.ChunkCount = int32(math.Ceil(float64(pack.GetFileSize()) / float64(data.ResourcePackChunkSize)))
	pk.CompressedPackSize = pack.GetFileSize()
	pk.Sha256 = pack.GetSha256()

	return pk
}

func (protocol *PacketManager) GetResourcePackInfo(mustAccept bool, resourcePacks *packs.Stack, behaviorPacks *packs.Stack) packets.IPacket {
	var pk = bedrock.NewResourcePackInfoPacket()
	pk.MustAccept = mustAccept

	var resourceEntries []types.ResourcePackInfoEntry
	var behaviorEntries []types.ResourcePackInfoEntry
	for _, pack := range *resourcePacks {
		resourceEntries = append(resourceEntries, types.ResourcePackInfoEntry{
			UUID:     pack.GetUUID(),
			Version:  pack.GetVersion(),
			PackSize: pack.GetFileSize(),
		})
	}
	for _, pack := range *behaviorPacks {
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

func (protocol *PacketManager) GetResourcePackStack(mustAccept bool, resourcePacks *packs.Stack, behaviorPacks *packs.Stack) packets.IPacket {
	var pk = bedrock.NewResourcePackStackPacket()
	pk.MustAccept = mustAccept
	var resourceEntries []types.ResourcePackStackEntry
	var behaviorEntries []types.ResourcePackStackEntry
	for _, pack := range *resourcePacks {
		resourceEntries = append(resourceEntries, types.ResourcePackStackEntry{
			UUID:    pack.GetUUID(),
			Version: pack.GetVersion(),
		})
	}
	for _, pack := range *behaviorPacks {
		behaviorEntries = append(behaviorEntries, types.ResourcePackStackEntry{
			UUID:    pack.GetUUID(),
			Version: pack.GetVersion(),
		})
	}

	pk.ResourcePacks = resourceEntries
	pk.BehaviorPacks = behaviorEntries

	return pk
}

func (protocol *PacketManager) GetServerHandshake(encryptionJwt string) packets.IPacket {
	var pk = bedrock.NewServerHandshakePacket()
	pk.Jwt = encryptionJwt

	return pk
}

func (protocol *PacketManager) GetSetEntityData(runtimeId uint64, data map[uint32][]interface{}) packets.IPacket {
	var pk = bedrock.NewSetEntityDataPacket()
	pk.RuntimeId = runtimeId
	pk.EntityData = data

	return pk
}

func (protocol *PacketManager) GetStartGame(player protocol.StartGameEntry, runtimeIdsTable []byte) packets.IPacket {
	var pk = bedrock.NewStartGamePacket()
	pk.Generator = 1
	pk.LevelSeed = 312402
	pk.TrustPlayers = true
	pk.DefaultPermissionLevel = permissions.LevelMember
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.PlayerGameMode = 1
	pk.PlayerPosition = player.GetPosition()
	pk.LevelGameMode = 1
	pk.LevelSpawnPosition = blocks.NewPosition(0, 40, 0)
	pk.CommandsEnabled = true
	pk.StartMap = false

	var gameRules = player.GetDimension().GetLevel().GetGameRules()
	var gameRuleEntries = map[string]types.GameRuleEntry{}
	for name, gameRule := range gameRules {
		gameRuleEntries[string(name)] = types.GameRuleEntry{Name: string(gameRule.GetName()), Value: gameRule.GetValue()}
	}

	pk.GameRules = gameRuleEntries
	pk.LevelName = player.GetDimension().GetLevel().GetName()
	pk.CurrentTick = player.GetDimension().GetLevel().GetCurrentTick()
	pk.Time = 0
	pk.AchievementsDisabled = true
	pk.BroadcastToXbox = true
	pk.BroadcastToLan = true
	pk.RuntimeIdsTable = runtimeIdsTable

	return pk
}

func (protocol *PacketManager) GetText(text types.Text) packets.IPacket {
	var pk = bedrock.NewTextPacket()
	pk.TextType = text.TextType
	pk.Translation = text.IsTranslation
	pk.Params = text.TranslationParameters
	pk.SourceName = text.SourceName
	pk.XUID = text.SourceXUID
	pk.Message = text.Message

	return pk
}

func (protocol *PacketManager) GetTransfer(address string, port uint16) packets.IPacket {
	var pk = bedrock.NewTransferPacket()
	pk.Address = address
	pk.Port = port

	return pk
}

func (protocol *PacketManager) GetUpdateAttributes(runtimeId uint64, attributeMap data2.AttributeMap) packets.IPacket {
	var pk = bedrock.NewUpdateAttributesPacket()
	pk.RuntimeId = runtimeId
	pk.Attributes = attributeMap

	return pk
}

func (protocol *PacketManager) GetNetworkChunkPublisherUpdatePacket(position blocks.Position, radius uint32) packets.IPacket {
	var pk = bedrock.NewNetworkChunkPublisherUpdatePacket()
	pk.Position = position
	pk.Radius = radius

	return pk
}

func (protocol *PacketManager) GetMoveEntity(runtimeId uint64, position r3.Vector, rot data2.Rotation, flags byte, teleport bool) packets.IPacket {
	var pk = bedrock.NewMoveEntityPacket()

	pk.RuntimeId = runtimeId
	pk.Position = position
	pk.Rotation = rot
	pk.Flags = flags

	if teleport {
		pk.Flags |= data.MoveEntityTeleport
	}

	return pk
}

func (protocol *PacketManager) GetPlayerSkin(uuid2 uuid.UUID, skinId, geometryName, geometryData string, skinData, capeData []byte) packets.IPacket {
	var pk = bedrock.NewPlayerSkinPacket()

	pk.UUID = uuid2
	pk.SkinId = skinId
	pk.SkinData = skinData
	pk.CapeData = capeData
	pk.GeometryName = geometryName
	pk.GeometryData = geometryData

	return pk
}

func (protocol *PacketManager) GetPlayerAction(runtimeId uint64, action int32, position blocks.Position, face int32) packets.IPacket {
	var pk = bedrock.NewPlayerActionPacket()

	pk.RuntimeId = runtimeId
	pk.Action = action
	pk.Position = position
	pk.Face = face

	return pk
}

func (protocol *PacketManager) GetAnimate(action int32, runtimeId uint64, float float32) packets.IPacket {
	var pk = bedrock.NewAnimatePacket()

	pk.RuntimeId = runtimeId
	pk.Action = action
	pk.Float = float

	return pk
}