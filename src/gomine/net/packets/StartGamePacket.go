package packets

import (
	"gomine/net/info"
	"gomine/vectors"
	"encoding/base64"
	"gomine/interfaces"
)

type StartGamePacket struct {
	*Packet
	EntityUniqueId int64
	EntityRuntimeId uint64

	PlayerGameMode int32
	PlayerPosition vectors.TripleVector

	Yaw float32
	Pitch float32

	LevelSeed int32
	Dimension int32
	Generator int32
	LevelGameMode int32
	Difficulty int32

	LevelSpawnPosition vectors.TripleVector
	AchievementsDisabled bool
	Time int32
	EduMode bool

	RainLevel float32
	LightningLevel float32

	MultiPlayerGame bool
	BroadcastToLan bool
	BroadcastToXbox bool

	CommandsEnabled bool
	ForcedResourcePacks bool

	GameRules map[string]interfaces.IGameRule

	BonusChest bool
	StartMap bool

	TrustPlayers bool
	DefaultPermissionLevel int32
	XboxBroadcastMode int32

	LevelName string
	CurrentTick int64
	EnchantmentSeed int32
}

func NewStartGamePacket() *StartGamePacket {
	return &StartGamePacket{Packet: NewPacket(info.StartGamePacket), GameRules: make(map[string]interfaces.IGameRule)}
}

func (pk *StartGamePacket) Encode()  {
	pk.PutVarLong(pk.EntityUniqueId) // Entity Unique ID
	pk.PutUnsignedVarLong(pk.EntityRuntimeId) // Entity runtime ID

	pk.PutVarInt(pk.PlayerGameMode) // Player game mode.

	pk.PutTripleVectorObject(pk.PlayerPosition) // Player pos.

	pk.PutLittleFloat(pk.Yaw) // Yaw
	pk.PutLittleFloat(pk.Pitch) // Pitch

	pk.PutVarInt(pk.LevelSeed) // Seed
	pk.PutVarInt(pk.Dimension) // Dimension
	pk.PutVarInt(pk.Generator) // Generator
	pk.PutVarInt(pk.LevelGameMode) // World gamemode
	pk.PutVarInt(pk.Difficulty) // Difficulty

	pk.PutBlockPos(pk.LevelSpawnPosition) // Spawn pos.
	pk.PutBool(pk.AchievementsDisabled) // Achievements disabled
	pk.PutVarInt(pk.Time) // Time
	pk.PutBool(pk.EduMode) // Education mode

	pk.PutLittleFloat(pk.RainLevel) // Rain level
	pk.PutLittleFloat(pk.LightningLevel) // Lightning level

	pk.PutBool(pk.MultiPlayerGame) // Multi-player game
	pk.PutBool(pk.BroadcastToLan) // LAN Broadcast
	pk.PutBool(pk.BroadcastToXbox) // XBOX Live Broadcast
	pk.PutBool(pk.CommandsEnabled) // Commands Enabled
	pk.PutBool(pk.ForcedResourcePacks) // Texture packs required

	// TODO: Fix. Something causes the client to crash with game rules.
	//pk.PutUnsignedVarInt(0)
	pk.PutGameRules(pk.GameRules) // Game rules

	pk.PutBool(pk.BonusChest) // Bonus chest
	pk.PutBool(pk.StartMap) // Start map
	pk.PutBool(pk.TrustPlayers) // Trust players
	pk.PutVarInt(pk.DefaultPermissionLevel) // Default permission level
	pk.PutVarInt(pk.XboxBroadcastMode) // XBOX Broadcast mode

	pk.PutString(base64.RawStdEncoding.EncodeToString([]byte(pk.LevelName))) // Level name base64 encoded
	pk.PutString(pk.LevelName) // Level name
	pk.PutString("") // Premium world template ID
	pk.PutBool(true) // Unknown
	pk.PutLittleLong(pk.CurrentTick) // Tick
	pk.PutVarInt(pk.EnchantmentSeed) // Enchantment seed
}

func (pk *StartGamePacket) Decode()  {

}