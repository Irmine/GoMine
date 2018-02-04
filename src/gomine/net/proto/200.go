package proto

import (
	"gomine/interfaces"
	"gomine/net/info"
	"gomine/net/packets/p200"
	"gomine/vectors"
	"gomine/entities/math"
	"gomine/players/handlers"
	math2 "math"
	"gomine/net/packets/data"
	"gomine/permissions"
	data2 "gomine/entities/data"
	"gomine/net/packets/types"
)

type Protocol200 struct {
	*Protocol
}

func NewProtocol200() *Protocol200 {
	var ids = info.PacketIds200
	var proto = &Protocol200{NewProtocol(200, info.PacketIds200, map[int]func() interfaces.IPacket {
		ids[info.LoginPacket]:				 		func() interfaces.IPacket { return p200.NewLoginPacket() },
		ids[info.ClientHandshakePacket]:			func() interfaces.IPacket { return p200.NewClientHandshakePacket() },
		ids[info.ResourcePackClientResponsePacket]:	func() interfaces.IPacket { return p200.NewResourcePackClientResponsePacket() },
		ids[info.RequestChunkRadiusPacket]:			func() interfaces.IPacket { return p200.NewRequestChunkRadiusPacket() },
		ids[info.MovePlayerPacket]:					func() interfaces.IPacket { return p200.NewMovePlayerPacket() },
		ids[info.CommandRequestPacket]:				func() interfaces.IPacket { return p200.NewCommandRequestPacket() },
		ids[info.ResourcePackChunkRequestPacket]:	func() interfaces.IPacket { return p200.NewResourcePackChunkRequestPacket() },
		ids[info.TextPacket]:						func() interfaces.IPacket { return p200.NewTextPacket() },
		ids[info.PlayerListPacket]:					func() interfaces.IPacket { return p200.NewPlayerListPacket() },
	}, map[int][][]interfaces.IPacketHandler{})}
	proto.initHandlers()

	return proto
}

func (protocol *Protocol200) initHandlers() {
	protocol.RegisterHandler(info.LoginPacket, handlers.NewLoginHandler(), 8)
	protocol.RegisterHandler(info.ClientHandshakePacket, handlers.NewClientHandshakeHandler(), 8)
	protocol.RegisterHandler(info.RequestChunkRadiusPacket, handlers.NewRequestChunkRadiusHandler(), 8)
	protocol.RegisterHandler(info.ResourcePackClientResponsePacket, handlers.NewResourcePackClientResponseHandler(), 8)
	protocol.RegisterHandler(info.MovePlayerPacket, handlers.NewMovePlayerHandler(), 8)
	protocol.RegisterHandler(info.CommandRequestPacket, handlers.NewCommandRequestHandler(), 8)
	protocol.RegisterHandler(info.ResourcePackChunkRequestPacket, handlers.NewResourcePackChunkRequestHandler(), 8)
	protocol.RegisterHandler(info.TextPacket, handlers.NewTextHandler(), 8)
}

func (protocol *Protocol200) GetAddEntity(entity interfaces.IEntity) *p200.AddEntityPacket {
	var pk = p200.NewAddEntityPacket()
	pk.UniqueId = entity.GetUniqueId()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.EntityType = entity.GetEntityId()
	pk.Position = *entity.GetPosition()
	pk.Motion = *entity.GetMotion()
	pk.Rotation = *entity.GetRotation()
	pk.Attributes = entity.GetAttributeMap()
	pk.EntityData = entity.GetEntityData()

	return pk
}

func (protocol *Protocol200) GetAddPlayer(player interfaces.IPlayer) *p200.AddPlayerPacket {
	var pk = p200.NewAddPlayerPacket()
	pk.UUID = player.GetUUID()
	pk.DisplayName = player.GetDisplayName()
	pk.Username = player.GetName()
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.Position = *player.GetPosition()
	pk.Rotation = *player.GetRotation()
	pk.Platform = player.GetPlatform()
	
	return pk
}

func (protocol *Protocol200) GetChunkRadiusUpdated(radius int32) *p200.ChunkRadiusUpdatedPacket {
	var pk = p200.NewChunkRadiusUpdatedPacket()
	pk.Radius = radius

	return pk
}

func (protocol *Protocol200) GetCraftingData() *p200.CraftingDataPacket {
	var pk = p200.NewCraftingDataPacket()

	return pk
}

func (protocol *Protocol200) GetDisconnect(message string, hideDisconnectScreen bool) *p200.DisconnectPacket {
	var pk = p200.NewDisconnectPacket()
	pk.HideDisconnectionScreen = hideDisconnectScreen
	pk.Message = message

	return pk
}

func (protocol *Protocol200) GetFullChunkData(chunk interfaces.IChunk) *p200.FullChunkDataPacket {
	var pk = p200.NewFullChunkDataPacket()
	pk.ChunkX = chunk.GetX()
	pk.ChunkZ = chunk.GetZ()
	pk.ChunkData = chunk.ToBinary()

	return pk
}

func (protocol *Protocol200) GetMovePlayer(runtimeId uint64, position vectors.TripleVector, rotation math.Rotation, mode byte, onGround bool, ridingRuntimeId uint64) *p200.MovePlayerPacket {
	var pk = p200.NewMovePlayerPacket()
	pk.RuntimeId = runtimeId
	pk.Position = position
	pk.Rotation = rotation
	pk.Mode = mode
	pk.OnGround = onGround
	pk.RidingRuntimeId = ridingRuntimeId

	return pk
}

func (protocol *Protocol200) GetPlayerList(listType byte, players map[string]interfaces.IPlayer) *p200.PlayerListPacket {
	var pk = p200.NewPlayerListPacket()
	pk.ListType = listType
	var entries = map[string]types.PlayerListEntry{}
	for name, player := range players {
		entries[name] = types.PlayerListEntry{
			UUID: player.GetUUID(),
			XUID: player.GetXUID(),
			EntityUniqueId: player.GetUniqueId(),
			Username: player.GetName(),
			DisplayName: player.GetDisplayName(),
			Platform: player.GetPlatform(),
			SkinId: player.GetSkinId(),
			SkinData: player.GetSkinData(),
			CapeData: player.GetCapeData(),
			GeometryName: player.GetGeometryName(),
			GeometryData: player.GetGeometryData(),
		}
	}
	pk.Entries = entries

	return pk
}

func (protocol *Protocol200) GetPlayStatus(status int32) *p200.PlayStatusPacket {
	var pk = p200.NewPlayStatusPacket()
	pk.Status = status

	return pk
}

func (protocol *Protocol200) GetRemoveEntity(uniqueId int64) *p200.RemoveEntityPacket {
	var pk = p200.NewRemoveEntityPacket()
	pk.EntityUniqueId = uniqueId

	return pk
}

func (protocol *Protocol200) GetResourcePackChunkData(packUUID string, chunkIndex int32, progress int64, data []byte) *p200.ResourcePackChunkDataPacket {
	var pk = p200.NewResourcePackChunkDataPacket()
	pk.PackUUID = packUUID
	pk.ChunkIndex = chunkIndex
	pk.Progress = progress
	pk.ChunkData = data

	return pk
}

func (protocol *Protocol200) GetResourcePackDataInfo(pack interfaces.IPack) *p200.ResourcePackDataInfoPacket {
	var pk = p200.NewResourcePackDataInfoPacket()
	pk.PackUUID = pack.GetUUID()
	pk.MaxChunkSize = data.ResourcePackChunkSize
	pk.ChunkCount = int32(math2.Ceil(float64(pack.GetFileSize()) / float64(data.ResourcePackChunkSize)))
	pk.CompressedPackSize = pack.GetFileSize()
	pk.Sha256 = pack.GetSha256()

	return pk
}

func (protocol *Protocol200) GetResourcePackInfo(mustAccept bool, resourcePacks []interfaces.IPack, behaviorPacks []interfaces.IPack) *p200.ResourcePackInfoPacket {
	var pk = p200.NewResourcePackInfoPacket()
	pk.MustAccept = mustAccept

	var resourceEntries []types.ResourcePackInfoEntry
	var behaviorEntries []types.ResourcePackInfoEntry
	for _, pack := range resourcePacks {
		resourceEntries = append(resourceEntries, types.ResourcePackInfoEntry{
			UUID: pack.GetUUID(),
			Version: pack.GetVersion(),
			PackSize: pack.GetFileSize(),
		})
	}
	for _, pack := range behaviorPacks {
		behaviorEntries = append(behaviorEntries, types.ResourcePackInfoEntry{
			UUID: pack.GetUUID(),
			Version: pack.GetVersion(),
			PackSize: pack.GetFileSize(),
		})
	}

	pk.ResourcePacks = resourceEntries
	pk.BehaviorPacks = behaviorEntries

	return pk
}

func (protocol *Protocol200) GetResourcePackStack(mustAccept bool, resourcePacks []interfaces.IPack, behaviorPacks []interfaces.IPack) *p200.ResourcePackStackPacket {
	var pk = p200.NewResourcePackStackPacket()
	pk.MustAccept = mustAccept
	var resourceEntries []types.ResourcePackStackEntry
	var behaviorEntries []types.ResourcePackStackEntry
	for _, pack := range resourcePacks {
		resourceEntries = append(resourceEntries, types.ResourcePackStackEntry{
			UUID: pack.GetUUID(),
			Version: pack.GetVersion(),
		})
	}
	for _, pack := range behaviorPacks {
		behaviorEntries = append(behaviorEntries, types.ResourcePackStackEntry{
			UUID: pack.GetUUID(),
			Version: pack.GetVersion(),
		})
	}

	pk.ResourcePacks = resourceEntries
	pk.BehaviorPacks = behaviorEntries

	return pk
}

func (protocol *Protocol200) GetServerHandshake(encryptionJwt string) *p200.ServerHandshakePacket {
	var pk = p200.NewServerHandshakePacket()
	pk.Jwt = encryptionJwt

	return pk
}

func (protocol *Protocol200) GetSetEntityData(entity interfaces.IEntity, data map[uint32][]interface{}) *p200.SetEntityDataPacket {
	var pk = p200.NewSetEntityDataPacket()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.EntityData = data

	return pk
}

func (protocol *Protocol200) GetStartGame(player interfaces.IPlayer) *p200.StartGamePacket {
	var pk = p200.NewStartGamePacket()
	pk.Generator = 1
	pk.LevelSeed = 312402
	pk.TrustPlayers = true
	pk.DefaultPermissionLevel = permissions.LevelMember
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.PlayerGameMode = 1
	pk.PlayerPosition = *vectors.NewTripleVector(20, 20, 20)
	pk.LevelGameMode = 1
	pk.LevelSpawnPosition = *vectors.NewTripleVector(20, 20, 20)
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

func (protocol *Protocol200) GetText(text types.Text) *p200.TextPacket {
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

func (protocol *Protocol200) GetTransfer(address string, port uint16) *p200.TransferPacket {
	var pk = p200.NewTransferPacket()
	pk.Address = address
	pk.Port = port

	return pk
}

func (protocol *Protocol200) GetUpdateAttributes(entity interfaces.IEntity, attributeMap *data2.AttributeMap) *p200.UpdateAttributesPacket {
	var pk = p200.NewUpdateAttributesPacket()
	pk.RuntimeId = entity.GetRuntimeId()
	pk.Attributes = attributeMap

	return pk
}