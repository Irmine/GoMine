package proto

import (
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/p201"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/permissions"
	"github.com/irmine/gomine/vectors"
)

type Protocol201 struct {
	*Protocol160
}

func NewProtocol201() *Protocol201 {
	var proto = &Protocol201{NewProtocol160()}
	proto.protocolNumber = 201

	return proto
}

func (protocol *Protocol201) GetStartGame(player interfaces.IPlayer) interfaces.IPacket {
	var pk = p201.NewStartGamePacket()
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
