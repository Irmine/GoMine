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
	WorldGameMode int32
	Difficulty int32

	WorldSpawnPosition vectors.TripleVector
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
	pk.PutVarLong(0) // Entity Unique ID
	pk.PutUnsignedVarLong(0) // Entity runtime ID

	pk.PutVarInt(1) // Player game mode.

	pk.PutTripleVectorObject(vectors.TripleVector{0, 20, 0}) // Player pos.

	pk.PutLittleFloat(0) // Yaw
	pk.PutLittleFloat(0) // Pitch

	pk.PutVarInt(12345) // Seed
	pk.PutVarInt(0) // Dimension
	pk.PutVarInt(12345) // Generator
	pk.PutVarInt(1) // World gamemode
	pk.PutVarInt(0) // Difficulty

	pk.PutBlockPos(vectors.TripleVector{0, 20, 0}) // Spawn pos.
	pk.PutBool(false) // Achievements disabled
	pk.PutVarInt(0) // Time
	pk.PutBool(false) // Education mode

	pk.PutLittleFloat(0) // Rain level
	pk.PutLittleFloat(0) // Lightning level

	pk.PutBool(true) // Multi-player game
	pk.PutBool(true) // LAN Broadcast
	pk.PutBool(true) // XBOX Live Broadcast
	pk.PutBool(true) // Commands Enabled
	pk.PutBool(false) // Texture packs required

	pk.PutUnsignedVarInt(1) // Game rule count
	pk.PutString("showcoordinates") // Game rule name
	pk.PutByte(1) // Game rule value type
	pk.PutBool(true) // Game rule value

	pk.PutBool(true) // Bonus chest
	pk.PutBool(true) // Start map
	pk.PutBool(true) // Trust players
	pk.PutVarInt(2) // Default permission level
	pk.PutVarInt(1) // XBOX Broadcast mode

	pk.PutString(base64.RawStdEncoding.EncodeToString([]byte("world"))) // Level name base64 encoded
	pk.PutString("world") // Level name
	pk.PutString("") // Premium world template ID
	pk.PutBool(true) // Unknown
	pk.PutLittleLong(100) // Tick
	pk.PutVarInt(312904) // Enchantment seed
}

func (pk *StartGamePacket) Decode()  {

}