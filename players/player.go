package players

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/worlds/entities"
)

type Player struct {
	*entities.Entity
	uuid     uuid.UUID
	xuid     string
	platform int32

	playerName  string
	displayName string

	skinId       string
	skinData     []byte
	capeData     []byte
	geometryName string
	geometryData string
}

// NewPlayer returns a new player with the given name.
func NewPlayer(uuid uuid.UUID, xuid string, platform int32, name string) *Player {
	var player = &Player{Entity: entities.New(entities.Player)}

	player.uuid = uuid
	player.xuid = xuid
	player.platform = platform

	player.playerName = name
	player.displayName = name

	return player
}

// GetName returns the username the player used to join the server.
func (player *Player) GetName() string {
	return player.playerName
}

// SetName sets the player name of this player.
// Note: This function is internal, and should not be used by plugins.
func (player *Player) SetName(name string) {
	player.playerName = name
}

// GetDisplayName returns the name the player shows in-game.
func (player *Player) GetDisplayName() string {
	return player.displayName
}

// SetDisplayName sets the name other players can see in-game.
func (player *Player) SetDisplayName(name string) {
	player.displayName = name
}

// GetUUID returns the UUID of the player.
func (player *Player) GetUUID() uuid.UUID {
	return player.uuid
}

// GetXUID returns the XUID of the player.
func (player *Player) GetXUID() string {
	return player.xuid
}

// GetPlatform returns the platform of the player.
func (player *Player) GetPlatform() int32 {
	return player.platform
}

// SpawnPlayerTo spawns this player to the given other player.
func (player *Player) SpawnPlayerTo(viewer entities.Viewer) {
	viewer.SendAddPlayer(player.GetUUID(), player.platform, player)
}

// SpawnPlayerToAll spawns this player to all other players.
func (player *Player) SpawnPlayerToAll() {
	for _, p := range player.Dimension.GetViewers() {
		if p.GetUUID() == player.GetUUID() {
			continue
		}
		if viewer, ok := p.(entities.Viewer); ok {
			player.SpawnPlayerTo(viewer)
		}
	}
}

// SetSkinId sets the skin ID/name of the player.
func (player *Player) SetSkinId(id string) {
	player.skinId = id
}

// GetSkinId returns the skin ID/name of the player.
func (player *Player) GetSkinId() string {
	return player.skinId
}

// GetSkinData returns the skin data of the player. (RGBA byte array)
func (player *Player) GetSkinData() []byte {
	return player.skinData
}

// SetSkinData sets the skin data of the player. (RGBA byte array)
func (player *Player) SetSkinData(data []byte) {
	player.skinData = data
}

// GetCapeData returns the cape data of the player. (RGBA byte array)
func (player *Player) GetCapeData() []byte {
	return player.capeData
}

// SetCapeData sets the cape data of the player. (RGBA byte array)
func (player *Player) SetCapeData(data []byte) {
	player.capeData = data
}

// GetGeometryName returns the geometry name of the player.
func (player *Player) GetGeometryName() string {
	return player.geometryName
}

// SetGeometryName sets the geometry name of the player.
func (player *Player) SetGeometryName(name string) {
	player.geometryName = name
}

// GetGeometryData returns the geometry data (json string) of the player.
func (player *Player) GetGeometryData() string {
	return player.geometryData
}

// SetGeometryData sets the geometry data (json string) of the player.
func (player *Player) SetGeometryData(data string) {
	player.geometryData = data
}

// SyncMove synchronizes the server's player movement with the client movement.
func (player *Player) SyncMove(x, y, z float64, pitch, yaw, headYaw float64, onGround bool) {
	player.SetPosition(r3.Vector{x, y, z})
	player.Rotation.Pitch += float64(pitch)
	player.Rotation.Yaw += float64(yaw)
	player.Rotation.HeadYaw = float64(headYaw)
	player.OnGround = onGround
}
