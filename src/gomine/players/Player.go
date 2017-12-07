package players

import (
	"gomine/interfaces"
	"goraklib/server"
	"gomine/net/packets"
	"gomine/entities"
	"gomine/worlds/locations"
	"fmt"
)

type Player struct {
	*entities.Human
	session *server.Session

	runtimeId uint64
	playerName  string
	displayName string

	permissions map[string]interfaces.IPermission
	permissionGroup interfaces.IPermissionGroup

	position locations.EntityPosition

	server interfaces.IServer
	level interfaces.ILevel
	dimension interfaces.IDimension

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
	player.position = locations.NewEntityPosition(0, 0, 0, 0, 0, 0, player.level, player.dimension)

	player.server = server
	player.session = session

	player.level = server.GetDefaultLevel()
	player.dimension = player.level.GetDefaultDimension()

	return player
}

func (player *Player) New(server interfaces.IServer, session *server.Session, name string, uuid string, xuid string, clientId int) interfaces.IPlayer {
	return NewPlayer(server, session, name, uuid, xuid, clientId)
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

/**
 * Sets the player position
 */
func (player *Player) SetPosition(v locations.EntityPosition) {
	player.position = v
}

/**
 * Returns the player position
 */
func (player *Player) GetPosition() locations.EntityPosition{
	return player.position
}

/**
 * Teleport player to a new position
 */
func (player *Player) Teleport(v locations.EntityPosition)  {
	pk := packets.NewMovePlayerPacket()
	pk.EntityId = player.runtimeId
	pk.Position = v
	pk.OnGround = false
	pk.RidingEid = 0
	player.server.GetRakLibAdapter().SendPacket(pk, player.GetSession())
	player.SetPosition(v)
}

//func (player *Player) SetDimension(dimension worlds.Dimension) {
//	player.dimension = dimension
//}

/**
 * Returns the dimension this player is currently in.
 */
func (player *Player) GetDimension() interfaces.IDimension {
	return player.dimension
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
func (player *Player) SendChunk(chunk interfaces.IChunk)  {
	var pk = packets.NewFullChunkPacket()

	pk.ChunkX = chunk.GetX()
	pk.ChunkZ = chunk.GetZ()
	pk.Chunk = chunk

	player.server.GetRakLibAdapter().SendPacket(pk, player.session)
}

/**
 * Called on every player move
 */
func (player *Player) Move(x, y, z, pitch, yaw, headYaw float32) {
	//fmt.Println("X : ", x)
	//fmt.Println("Y : ", x)
	//fmt.Println("Z : ", x)
}

func (player *Player) Tick() {

}