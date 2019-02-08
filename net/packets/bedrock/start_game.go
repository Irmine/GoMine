package bedrock

import (
	"encoding/base64"

	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/worlds/blocks"
)

const (
	GameBroadcastSettingNone = iota
	GameBroadcastSettingInviteOnly
	GameBroadcastSettingFriendsOnly
	GameBroadcastSettingFriendsOfFriends
	GameBroadcastSettingPublic
)

type StartGamePacket struct {
	*packets.Packet
	EntityUniqueId                 int64
	EntityRuntimeId                uint64
	PlayerGameMode                 int32
	PlayerPosition                 r3.Vector
	Yaw                            float32
	Pitch                          float32
	LevelSeed                      int32
	Dimension                      int32
	Generator                      int32
	LevelGameMode                  int32
	Difficulty                     int32
	LevelSpawnPosition             blocks.Position
	AchievementsDisabled           bool
	Time                           int32
	EduMode                        bool
	EduFeaturesEnabled             bool
	RainLevel                      float32
	LightningLevel                 float32
	Bool1                          bool
	MultiPlayerGame                bool
	BroadcastToLan                 bool
	CommandsEnabled                bool
	ForcedResourcePacks            bool
	GameRules                      map[string]types.GameRuleEntry
	BonusChest                     bool
	StartMap                       bool
	DefaultPermissionLevel         int32
	LevelName                      string
	IsTrial                        bool
	CurrentTick                    int64
	EnchantmentSeed                int32
	ServerChunkTickRange           int32
	PlatformBroadcast              bool
	XBOXBroadcastIntent            int32
	PlatformBroadcastIntent        int32
	LockedBehaviorPack             bool
	LockedResourcePack             bool
	FromLockedWorldTemplate        bool
	UseMsaGamertagsOnly            bool
	FromWorldTemplate              bool
	WorldTemplateOptionLocked      bool
	RuntimeIdsTable                []byte
	MultiplayerCorrelationID       string
}

func NewStartGamePacket() *StartGamePacket {
	return &StartGamePacket{Packet: packets.NewPacket(info.PacketIds[info.StartGamePacket]), GameRules: make(map[string]types.GameRuleEntry)}
}

func (pk *StartGamePacket) Encode() {
	pk.PutEntityUniqueId(pk.EntityUniqueId)   // Entity Unique ID
	pk.PutEntityRuntimeId(pk.EntityRuntimeId) // Entity runtime ID

	pk.PutVarInt(pk.PlayerGameMode) // Player game mode.
	pk.PutVector(pk.PlayerPosition) // Player pos.

	pk.PutLittleFloat(pk.Pitch) // Pitch
	pk.PutLittleFloat(pk.Yaw)   // Yaw

	pk.PutVarInt(pk.LevelSeed)     // Seed
	pk.PutVarInt(pk.Dimension)     // Dimension
	pk.PutVarInt(pk.Generator)     // Generator
	pk.PutVarInt(pk.LevelGameMode) // World gamemode
	pk.PutVarInt(pk.Difficulty)    // Difficulty

	pk.PutBlockPosition(pk.LevelSpawnPosition) // Spawn pos.
	pk.PutBool(pk.AchievementsDisabled)        // Achievements disabled
	pk.PutVarInt(pk.Time)                      // Time
	pk.PutBool(pk.EduMode)                     // Education mode
	pk.PutBool(pk.EduFeaturesEnabled)          // Education mode features enabled

	pk.PutLittleFloat(pk.RainLevel)      // Rain level
	pk.PutLittleFloat(pk.LightningLevel) // Lightning level

	pk.PutBool(pk.Bool1)
	pk.PutBool(pk.MultiPlayerGame)     // Multi-player game
	pk.PutBool(pk.BroadcastToLan)      // LAN Broadcast
	pk.PutVarInt(pk.XBOXBroadcastIntent)
	pk.PutVarInt(pk.PlatformBroadcastIntent)

	pk.PutBool(pk.CommandsEnabled)     // Commands Enabled
	pk.PutBool(pk.ForcedResourcePacks) // Texture packs required

	pk.PutGameRules(pk.GameRules) // Game rules

	pk.PutBool(pk.BonusChest)                // Bonus chest
	pk.PutBool(pk.StartMap)                  // Start map
	pk.PutVarInt(pk.DefaultPermissionLevel)  // Default permission level
	pk.PutLittleInt(pk.ServerChunkTickRange) // Server chunk tick range
	pk.PutBool(pk.LockedBehaviorPack)        // Has Locked Behavior Pack
	pk.PutBool(pk.LockedResourcePack)        // Has Locked Resource Pack
	pk.PutBool(pk.FromLockedWorldTemplate)   // From World Locked Template
	pk.PutBool(pk.UseMsaGamertagsOnly)       // Use Msa Gamertags Only
	pk.PutBool(pk.FromWorldTemplate)         // From World Template
	pk.PutBool(pk.WorldTemplateOptionLocked) // World template option locked

	pk.PutString(base64.RawStdEncoding.EncodeToString([]byte(pk.LevelName))) // Level name base64 encoded
	pk.PutString(pk.LevelName)                                               // Level name
	pk.PutString("")                                                     // Premium world template ID
	pk.PutBool(pk.IsTrial)                                                   // Is Trial
	pk.PutLittleLong(pk.CurrentTick)                                         // Tick
	pk.PutVarInt(pk.EnchantmentSeed)                                         // Enchantment seed
	pk.PutBytes(pk.RuntimeIdsTable)
	pk.PutString(pk.MultiplayerCorrelationID)
}

func (pk *StartGamePacket) Decode() {

}
