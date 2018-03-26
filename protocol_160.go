package gomine

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/p160"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/permissions"
)

type P160 struct {
	*P200
}

func NewP160(server *Server) *P160 {
	var proto = &P160{NewP200(server)}
	proto.ProtocolNumber = 160
	proto.DeregisterPacketHandlers(info.TextPacket, 8)
	proto.RegisterHandler(info.TextPacket, NewTextHandler_160(server))
	proto.RegisterPacket(info.PacketIds200[info.TextPacket], func() packets.IPacket { return p160.NewTextPacket() })

	return proto
}

func (protocol *P160) GetAddPlayer(uuid uuid.UUID, platform int32, viewer protocol.AddPlayerEntry) packets.IPacket {
	var pk = p160.NewAddPlayerPacket()
	pk.UUID = uuid
	pk.Username = viewer.GetDisplayName()
	pk.EntityRuntimeId = viewer.GetRuntimeId()
	pk.EntityUniqueId = viewer.GetUniqueId()
	pk.Position = viewer.GetPosition()
	pk.Rotation = viewer.GetRotation()

	return pk
}

func (protocol *P160) GetStartGame(player protocol.StartGameEntry) packets.IPacket {
	var pk = p160.NewStartGamePacket()
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

	return pk
}

func (protocol *P160) GetText(text types.Text) packets.IPacket {
	var pk = p160.NewTextPacket()
	pk.TextType = text.TextType
	pk.IsTranslation = text.IsTranslation
	pk.TranslationParameters = text.TranslationParameters
	pk.SourceName = text.SourceName
	pk.XUID = text.SourceXUID
	pk.Message = text.Message

	return pk
}

func (protocol *P160) GetPlayerList(listType byte, players map[string]protocol.PlayerListEntry) packets.IPacket {
	var pk = p160.NewPlayerListPacket()
	pk.ListType = listType
	var entries = map[string]types.PlayerListEntry{}
	for _, player := range players {
		entries[player.GetName()] = types.PlayerListEntry{
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
