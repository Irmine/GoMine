package players

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/worlds"
	"gomine/vectors"
	"gomine/net/packets"
	"gomine/entities"
)

type Player struct {
	session *server.Session

	runtimeId uint64
	playerName  string
	displayName string

	permissions map[string]interfaces.IPermission
	permissionGroup interfaces.IPermissionGroup

	position *vectors.TripleVector
	yaw, headYaw, pitch float32

	server interfaces.IServer
	dimension worlds.Dimension

	language string

	uuid string
	xuid string
	clientId int

	viewDistance int32

	skinId string
	skinData []byte
	capeData []byte
	geometryName string
	geometryData string
}

/**
 * Returns a new player with the given credentials.
 */
func NewPlayer(server interfaces.IServer, session *server.Session, name string, uuid string, xuid string, clientId int) *Player {
	var player = &Player{}
	entities.RuntimeId++
	player.runtimeId = entities.RuntimeId
	player.playerName = name
	player.displayName = name

	player.uuid = uuid
	player.xuid = xuid
	player.clientId = clientId

	player.permissions = make(map[string]interfaces.IPermission)
	player.permissionGroup = server.GetPermissionManager().GetDefaultGroup()
	player.position = vectors.NewTripleVector(0, 0, 0)

	player.server = server
	player.session = session

	return player
}

/**
 * Returns the UUID of this player.
 */
func (player *Player) GetUUID() string {
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

func (player *Player) SetPosition(v *vectors.TripleVector) {
	pk := packets.NewMovePlayerPacket()
	pk.EntityId = player.runtimeId
	pk.Position = *v
	pk.Pitch = player.pitch
	pk.Yaw = player.yaw
	pk.HeadYaw = player.headYaw
	pk.Mode = packets.Teleport
	pk.OnGround = false
	pk.RidingEid = 0
	player.server.GetRakLibAdapter().SendPacket(pk, player.GetSession())
	*player.position = *v
}

func (player *Player) GetPosition() *vectors.TripleVector {
	return player.position
}

//func (player *Player) SetDimension(dimension worlds.Dimension) {
//	player.dimension = dimension
//}

func (player *Player) GetDimension() worlds.Dimension {
	return player.dimension
}

func (player *Player) SetSkinId(id string) {
	player.skinId = id
}

func (player *Player) GetSkinId() string {
	return player.skinId
}

func (player *Player) GetSkinData() []byte {
	return player.skinData
}

func (player *Player) SetSkinData(data []byte) {
	player.skinData = data
}

func (player *Player) GetCapeData() []byte {
	return player.capeData
}

func (player *Player) SetCapeData(data []byte) {
	player.capeData = data
}

func (player *Player) GetGeometryName() string {
	return player.geometryName
}

func (player *Player) SetGeometryName(name string) {
	player.geometryName = name
}

func (player *Player) GetGeometryData() string {
	return player.geometryData
}

func (player *Player) SetGeometryData(data string) {
	player.geometryData = data
}

/**
 * Returns the GoRakLib session of this player.
 */
func (player *Player) GetSession() *server.Session {
	return player.session
}

func (player *Player) SendChunk(chunk interfaces.IChunk)  {
	var pk = packets.NewFullChunkPacket()
	pk.ChunkX = int32(chunk.GetX())
	pk.ChunkZ = int32(chunk.GetZ())
	pk.Chunk = chunk
	player.server.GetRakLibAdapter().SendPacket(pk, player.session)
}

func (player *Player) Tick() {
	player.dimension.RequestChunks(player)
}