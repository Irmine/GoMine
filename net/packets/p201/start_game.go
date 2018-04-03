package p201

import (
	"encoding/base64"

	"github.com/golang/geo/r3"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/worlds/blocks"
)

type StartGamePacket struct {
	*packets.Packet
	EntityUniqueId         int64
	EntityRuntimeId        uint64
	PlayerGameMode         int32
	PlayerPosition         r3.Vector
	Yaw                    float32
	Pitch                  float32
	LevelSeed              int32
	Dimension              int32
	Generator              int32
	LevelGameMode          int32
	Difficulty             int32
	LevelSpawnPosition     blocks.Position
	AchievementsDisabled   bool
	Time                   int32
	EduMode                bool
	RainLevel              float32
	LightningLevel         float32
	MultiPlayerGame        bool
	BroadcastToLan         bool
	BroadcastToXbox        bool
	CommandsEnabled        bool
	ForcedResourcePacks    bool
	GameRules              map[string]types.GameRuleEntry
	BonusChest             bool
	StartMap               bool
	TrustPlayers           bool
	DefaultPermissionLevel int32
	XBOXBroadcastMode      int32
	LevelName              string
	IsTrial                bool
	CurrentTick            int64
	EnchantmentSeed        int32
	ServerChunkTickRange   int32
}

func NewStartGamePacket() *StartGamePacket {
	return &StartGamePacket{Packet: packets.NewPacket(info.PacketIds200[info.StartGamePacket]), GameRules: make(map[string]types.GameRuleEntry)}
}

func (pk *StartGamePacket) Encode() {
	pk.PutEntityUniqueId(pk.EntityUniqueId)   // Entity Unique ID
	pk.PutEntityRuntimeId(pk.EntityRuntimeId) // Entity runtime ID

	pk.PutVarInt(pk.PlayerGameMode) // Player game mode.

	pk.PutVector(pk.PlayerPosition) // Player pos.

	pk.PutLittleFloat(pk.Yaw)   // Yaw
	pk.PutLittleFloat(pk.Pitch) // Pitch

	pk.PutVarInt(pk.LevelSeed)     // Seed
	pk.PutVarInt(pk.Dimension)     // Dimension
	pk.PutVarInt(pk.Generator)     // Generator
	pk.PutVarInt(pk.LevelGameMode) // World gamemode
	pk.PutVarInt(pk.Difficulty)    // Difficulty

	pk.PutBlockPosition(pk.LevelSpawnPosition) // Spawn pos.
	pk.PutBool(pk.AchievementsDisabled)        // Achievements disabled
	pk.PutVarInt(pk.Time)                      // Time
	pk.PutBool(pk.EduMode)                     // Education mode

	pk.PutLittleFloat(pk.RainLevel)      // Rain level
	pk.PutLittleFloat(pk.LightningLevel) // Lightning level

	pk.PutBool(pk.MultiPlayerGame)     // Multi-player game
	pk.PutBool(pk.BroadcastToLan)      // LAN Broadcast
	pk.PutBool(pk.BroadcastToXbox)     // XBOX Live Broadcast
	pk.PutBool(pk.CommandsEnabled)     // Commands Enabled
	pk.PutBool(pk.ForcedResourcePacks) // Texture packs required

	pk.PutGameRules(pk.GameRules) // Game rules

	pk.PutBool(pk.BonusChest)               // Bonus chest
	pk.PutBool(pk.StartMap)                 // Start map
	pk.PutBool(pk.TrustPlayers)             // Trust players
	pk.PutVarInt(pk.DefaultPermissionLevel) // Default permission level
	pk.PutVarInt(pk.XBOXBroadcastMode)      // XBOX Broadcast mode
	pk.PutLittleInt(pk.ServerChunkTickRange)

	pk.PutString(base64.RawStdEncoding.EncodeToString([]byte(pk.LevelName))) // Level name base64 encoded
	pk.PutString(pk.LevelName)                                               // Level name
	pk.PutString("")                                                         // Premium world template ID
	pk.PutBool(pk.IsTrial)                                                   // Is Trial
	pk.PutLittleLong(pk.CurrentTick)                                         // Tick
	pk.PutVarInt(pk.EnchantmentSeed)                                         // Enchantment seed
}

func (pk *StartGamePacket) Decode() {

}
