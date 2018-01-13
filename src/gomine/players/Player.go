package players

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/entities"
	"gomine/vectors"
	"gomine/entities/math"
	"gomine/utils"
	math2 "math"
)

type Player struct {
	*entities.Human
	session *server.Session

	attributeMap *entities.AttributeMap
	runtimeId uint64
	playerName  string
	displayName string

	closed bool

	permissions map[string]interfaces.IPermission
	permissionGroup interfaces.IPermissionGroup

	server interfaces.IServer

	language string

	uuid utils.UUID
	xuid string
	clientId int

	onGround bool
	viewDistance int32

	skinId string
	skinData []byte
	capeData []byte
	geometryName string
	geometryData string

	finalized bool

	usedChunks map[int]interfaces.IChunk
}

/**
 * Returns a new player with the given credentials.
 */
func NewPlayer(server interfaces.IServer, session *server.Session, name string, uuid utils.UUID, xuid string, clientId int) *Player {
	var player = &Player{}

	player.runtimeId = entities.RuntimeId
	player.playerName = name
	player.displayName = name

	player.usedChunks = make(map[int]interfaces.IChunk)

	player.uuid = uuid
	player.xuid = xuid
	player.clientId = clientId

	player.permissions = make(map[string]interfaces.IPermission)
	player.permissionGroup = server.GetPermissionManager().GetDefaultGroup()

	player.server = server
	player.session = session

	return player
}

/**
 * Returns a new player.
 */
func (player *Player) New(server interfaces.IServer, session *server.Session, name string, uuid utils.UUID, xuid string, clientId int) interfaces.IPlayer {
	return NewPlayer(server, session, name, uuid, xuid, clientId)
}

/**
 * Returns if this player has been placed in a world.
 */
func (player *Player) IsInWorld() bool {
	return player.Dimension != nil && player.Level != nil
}

/**
 * Places this player inside of a level and dimension.
 */
func (player *Player) PlaceInWorld(position *vectors.TripleVector, rotation *math.Rotation, level interfaces.ILevel, dimension interfaces.IDimension) {
	player.Human = entities.NewHuman(player.GetDisplayName(), position, rotation, vectors.NewTripleVector(0, 0, 0), level, dimension)
}

/**
 * Checks if this player is finalized.
 */
func (player *Player) IsFinalized() bool {
	return player.finalized
}

/**
 * Sets this player finalized.
 */
func (player *Player) SetFinalized() {
	player.finalized = true
}

/**
 * Spawns this player to the given other player.
 */
func (player *Player) SpawnPlayerTo(player2 interfaces.IPlayer) {
	player.GetLevel().GetEntityHelper().SpawnPlayerTo(player, player2)
}

/**
 * Spawns this player to all other players.
 */
func (player *Player) SpawnPlayerToAll() {
	for _, p := range player.GetServer().GetPlayerFactory().GetPlayers() {
		if player == p {
			continue
		}
		player.SpawnPlayerTo(p)
	}
}

/**
 * Returns the UUID of this player.
 */
func (player *Player) GetUUID() utils.UUID {
	return player.uuid
}

/**
 * Returns the XUID of this player.
 */
func (player *Player) GetXUID() string {
	return player.xuid
}

/**
 * Sets the language (locale) of this player.
 */
func (player *Player) SetLanguage(language string) {
	player.language = language
}

/**
 * Returns the language (locale) of this player.
 */
func (player *Player) GetLanguage() string {
	return player.language
}

/**
 * Returns the client ID of this player.
 */
func (player *Player) GetClientId() int {
	return player.clientId
}

/**
 * Sets the view distance of this player.
 */
func (player *Player) SetViewDistance(distance int32) {
	player.viewDistance = distance
}

/**
 * Returns the view distance of this player.
 */
func (player *Player) GetViewDistance() int32 {
	return player.viewDistance
}

/**
 * Returns the main server.
 */
func (player *Player) GetServer() interfaces.IServer {
	return player.server
}

/**
 * Returns the username the player used to join the server.
 */
func (player *Player) GetName() string {
	return player.playerName
}

/**
 * Returns the name the player shows in-game.
 */
func (player *Player) GetDisplayName() string {
	return player.displayName
}

/**
 * Sets the name other players can see in-game.
 */
func (player *Player) SetDisplayName(name string) {
	player.displayName = name
}

/**
 * Returns the permission group this player is in.
 */
func (player *Player) GetPermissionGroup() interfaces.IPermissionGroup {
	return player.permissionGroup
}

/**
 * Sets the permission group of this player.
 */
func (player *Player) SetPermissionGroup(group interfaces.IPermissionGroup) {
	player.permissionGroup = group
}

/**
 * Checks if this player has a permission.
 */
func (player *Player) HasPermission(permission string) bool {
	if player.GetPermissionGroup().HasPermission(permission) {
		return true
	}
	var _, exists = player.permissions[permission]
	return exists
}

/**
 * Adds a permission to the player.
 * Returns true if a permission with the same name was overwritten.
 */
func (player *Player) AddPermission(permission interfaces.IPermission) bool {
	var hasPermission = player.HasPermission(permission.GetName())

	player.permissions[permission.GetName()] = permission

	return hasPermission
}

/**
 * Deletes a permission from the player.
 * This does not delete the permission from the group the player is in.
 */
func (player *Player) RemovePermission(permission string) bool {
	if !player.HasPermission(permission) {
		return false
	}

	delete(player.permissions, permission)

	return true
}

/**
 * Teleport player to a new position
 */
func (player *Player) Teleport(v *vectors.TripleVector, rot *math.Rotation)  {
	pk := packets.NewMovePlayerPacket()
	pk.EntityId = player.runtimeId
	pk.Position = *v
	pk.Rotation = *rot
	pk.OnGround = player.onGround
	pk.RidingEid = 0
	player.SendPacket(pk)

	player.Position = v
}

/**
 * Sets the skin ID/name of the player.
 */
func (player *Player) SetSkinId(id string) {
	player.skinId = id
}

/**
 * Returns the skin ID/name of the player.
 */
func (player *Player) GetSkinId() string {
	return player.skinId
}

/**
 * Returns the skin data of the player. (RGBA byte array)
 */
func (player *Player) GetSkinData() []byte {
	return player.skinData
}

/**
 * Sets the skin data of the player. (RGBA byte array)
 */
func (player *Player) SetSkinData(data []byte) {
	player.skinData = data
}

/**
 * Returns the cape data of the player. (RGBA byte array)
 */
func (player *Player) GetCapeData() []byte {
	return player.capeData
}

/**
 * Sets the cape data of the player. (RGBA byte array)
 */
func (player *Player) SetCapeData(data []byte) {
	player.capeData = data
}

/**
 * Returns the geometry name of the player.
 */
func (player *Player) GetGeometryName() string {
	return player.geometryName
}

/**
 * Sets the geometry name of the player.
 */
func (player *Player) SetGeometryName(name string) {
	player.geometryName = name
}

/**
 * Returns the geometry data (json string) of the player.
 */
func (player *Player) GetGeometryData() string {
	return player.geometryData
}

/**
 * Sets the geometry data (json string) of the player.
 */
func (player *Player) SetGeometryData(data string) {
	player.geometryData = data
}

/**
 * Returns the GoRakLib session of this player.
 */
func (player *Player) GetSession() *server.Session {
	return player.session
}

/**
 * Returns the ping of the player in milliseconds.
 */
func (player *Player) GetPing() uint64 {
	return player.session.GetPing()
}

/**
 * Sends a chunk to the player.
 */
func (player *Player) SendChunk(chunk interfaces.IChunk, index int)  {
	var pk = packets.NewFullChunkPacket()
	pk.Chunk = chunk
	player.usedChunks[index] = chunk

	player.SendPacket(pk)
}

/**
 * Synchronizes the server's player movement with the client movement and adjusts chunks.
 */
func (player *Player) SyncMove(x, y, z, pitch, yaw, headYaw float32) {
	player.SetPosition(vectors.NewTripleVector(x, y, z))
	player.Rotation.Pitch += pitch
	player.Rotation.Yaw += yaw
	player.Rotation.HeadYaw += headYaw

	var chunkX = int32(math2.Floor(float64(x))) >> 4
	var chunkZ = int32(math2.Floor(float64(z))) >> 4

	var rs = player.GetViewDistance() * player.GetViewDistance()

	for index, chunk := range player.usedChunks {
		xDist := chunkX - chunk.GetX()
		zDist := chunkZ - chunk.GetZ()
		if xDist * xDist + zDist * zDist > rs {
			delete(player.usedChunks, index)
			for _, entity := range chunk.GetEntities() {
				entity.DespawnFrom(player)
			}
		}
	}
}

/**
 * Checks if the player has a chunk with the given index in use.
 */
func (player *Player) HasChunkInUse(index int) bool {
	_, ok := player.usedChunks[index]
	return ok
}

/**
 * Checks if the player has any used chunks.
 */
func (player *Player) HasAnyChunkInUse() bool {
	return len(player.usedChunks) > 0
}

/**
 * Sends a packet to this player.
 */
func (player *Player) SendPacket(packet interfaces.IPacket) {
	player.server.GetRakLibAdapter().SendPacket(packet, player.session, server.PriorityMedium)
}

func (player *Player) Tick() {

}

/**
 * Sends a message to this player.
 */
func (player *Player) SendMessage(message string) {
	var pk = packets.NewTextPacket()
	pk.Message = message
	player.SendPacket(pk)
}