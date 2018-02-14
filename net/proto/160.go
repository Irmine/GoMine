package proto

import (
	"gomine/net/info"
	"gomine/interfaces"
	"gomine/net/packets/p160"
	"gomine/net/packets/types"
	"gomine/permissions"
	"gomine/vectors"
	p160handler "gomine/players/handlers/p160"
)

type Protocol160 struct {
	*Protocol200
}

func NewProtocol160() *Protocol160 {
	var proto = &Protocol160{NewProtocol200()}
	proto.protocolNumber = 160
	proto.DeregisterPacketHandlers(info.TextPacket, 8)
	proto.RegisterHandler(info.TextPacket, p160handler.NewTextHandler(), 8)
	proto.RegisterPacket(info.PacketIds200[info.TextPacket], func() interfaces.IPacket { return p160.NewTextPacket() })

	return proto
}

func (protocol *Protocol160) GetAddPlayer(player interfaces.IPlayer) interfaces.IPacket {
	var pk = p160.NewAddPlayerPacket()
	pk.UUID = player.GetUUID()
	pk.Username = player.GetDisplayName()
	pk.EntityRuntimeId = player.GetRuntimeId()
	pk.EntityUniqueId = player.GetUniqueId()
	pk.Position = *player.GetPosition()
	pk.Rotation = *player.GetRotation()

	return pk
}

func (protocol *Protocol160) GetStartGame(player interfaces.IPlayer) interfaces.IPacket {
	var pk = p160.NewStartGamePacket()
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

func (protocol *Protocol160) GetText(text types.Text) interfaces.IPacket {
	var pk = p160.NewTextPacket()
	pk.TextType = text.TextType
	pk.IsTranslation = text.IsTranslation
	pk.TranslationParameters = text.TranslationParameters
	pk.SourceName = text.SourceName
	pk.XUID = text.SourceXUID
	pk.Message = text.Message

	return pk
}

func (protocol *Protocol160) GetPlayerList(listType byte, players map[string]interfaces.IPlayer) interfaces.IPacket {
	var pk = p160.NewPlayerListPacket()
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