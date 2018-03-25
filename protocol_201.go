package gomine

import (
	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/p201"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/permissions"
)

type P201 struct {
	*P160
}

func NewP201(server *Server) *P201 {
	var proto = &P201{NewP160(server)}
	proto.ProtocolNumber = 201

	return proto
}

func (protocol *P201) GetStartGame(player protocol.StartGameEntry) packets.IPacket {
	var pk = p201.NewStartGamePacket()
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
