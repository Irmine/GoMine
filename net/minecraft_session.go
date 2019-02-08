package net

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/permissions"
	"github.com/irmine/gomine/players"
	"github.com/irmine/gomine/text"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/goraklib/protocol"
	"github.com/irmine/goraklib/server"
	"github.com/irmine/worlds"
	"github.com/irmine/worlds/blocks"
	"github.com/irmine/worlds/chunks"
	"math"
	"strings"
)

type MinecraftSession struct {
	adapter *NetworkAdapter
	session *server.Session

	player *players.Player

	uuid     uuid.UUID
	xuid     string
	clientId int

	protocolNumber   int32
	minecraftVersion string

	language string

	clientPlatform int32

	encryptionHandler     *utils.EncryptionHandler
	usesEncryption        bool
	xboxLiveAuthenticated bool

	viewDistance int32
	chunkLoader  *worlds.Loader

	permissions     map[string]*permissions.Permission
	permissionGroup *permissions.Group

	Connected         bool
}

// NewMinecraftSession returns a new Minecraft session with the given RakNet session.
func NewMinecraftSession(adapter *NetworkAdapter, session *server.Session) *MinecraftSession {
	return &MinecraftSession{adapter, session, nil, uuid.New(), "", 0, 0, "", "", 0, utils.NewEncryptionHandler(), false, false, 0, nil, nil, nil, false}
}

// SetData sets the basic session data of the Minecraft Session
func (session *MinecraftSession) SetData(permissionManager *permissions.Manager, data types.SessionData) {
	session.permissions = make(map[string]*permissions.Permission)
	session.permissionGroup = permissionManager.GetDefaultGroup()

	session.uuid = data.ClientUUID
	session.xuid = data.ClientXUID
	session.clientId = data.ClientId
	session.protocolNumber = data.ProtocolNumber
	session.minecraftVersion = data.GameVersion
	session.language = data.Language
	session.clientPlatform = int32(data.DeviceOS)
	session.chunkLoader = worlds.NewLoader(nil, 0, 0)
	session.chunkLoader.PublisherUpdateFunction = func() {
		var vector = session.player.Position
		var position = blocks.NewPosition(int32(vector.X), uint32(vector.Y), int32(vector.Z))
		session.SendNetworkChunkPublisherUpdate(position, uint32(session.GetViewDistance() * 16))
	}
	session.chunkLoader.LoadFunction = func(chunk *chunks.Chunk) {
		session.SendFullChunkData(chunk)
		chunk.AddViewer(session)
		chunk.AddEntity(session.player)
	}
	session.chunkLoader.UnloadFunction = func(chunk *chunks.Chunk) {
		chunk.RemoveViewer(session)
		chunk.RemoveEntity(session.player.GetRuntimeId())
	}
}

// GetPlayer returns the player associated with the Minecraft session.
// This player may not yet exist during the login sequence, and this function may return nil.
func (session *MinecraftSession) GetPlayer() *players.Player {
	return session.player
}

// SetPlayer sets the player associated with the Minecraft session.
// Network actions will be executed on this player.
func (session *MinecraftSession) SetPlayer(player *players.Player) {
	session.player = player
}

// GetName returns the name of the player under the session.
func (session *MinecraftSession) GetName() string {
	if session.player == nil {
		return ""
	}
	return session.player.GetName()
}

// GetDisplayName returns the display name of the player under the session.
func (session *MinecraftSession) GetDisplayName() string {
	if session.player == nil {
		return ""
	}
	return session.player.GetDisplayName()
}

// HasSpawned checks if the player of the session has spawned.
func (session *MinecraftSession) HasSpawned() bool {
	return session.GetPlayer().GetDimension() != nil
}

// SetViewDistance sets the view distance of this player.
func (session *MinecraftSession) SetViewDistance(distance int32) {
	session.viewDistance = distance
}

// GetViewDistance returns the view distance of this player.
func (session *MinecraftSession) GetViewDistance() int32 {
	return session.viewDistance
}

// GetChunkLoader returns the chunk loader of the session.
func (session *MinecraftSession) GetChunkLoader() *worlds.Loader {
	return session.chunkLoader
}

// GetPlatform returns the platform the client uses to player the game.
func (session *MinecraftSession) GetPlatform() int32 {
	return session.clientPlatform
}

// GetProtocolNumber returns the bedrock number the client used to join the server.
func (session *MinecraftSession) GetProtocolNumber() int32 {
	return session.protocolNumber
}

// GetGameVersion returns the Minecraft version the player used to join the server.
func (session *MinecraftSession) GetGameVersion() string {
	return session.minecraftVersion
}

// GetSession returns the GoRakLib session of this session.
func (session *MinecraftSession) GetSession() *server.Session {
	return session.session
}

// GetPing returns the ping of the session in milliseconds.
func (session *MinecraftSession) GetPing() int64 {
	return session.session.CurrentPing
}

// GetUUID returns the UUID of this session.
func (session *MinecraftSession) GetUUID() uuid.UUID {
	return session.uuid
}

// GetXUID returns the XUID of this session.
func (session *MinecraftSession) GetXUID() string {
	return session.xuid
}

// SetLanguage sets the language (locale) of this session.
func (session *MinecraftSession) SetLanguage(language string) {
	session.language = language
}

// GetLanguage returns the language (locale) of this session.
func (session *MinecraftSession) GetLanguage() string {
	return session.language
}

// GetClientId returns the client ID of this session.
func (session *MinecraftSession) GetClientId() int {
	return session.clientId
}

// GetEncryptionHandler returns the handler used for encryption.
func (session *MinecraftSession) GetEncryptionHandler() *utils.EncryptionHandler {
	return session.encryptionHandler
}

// UsesEncryption checks if the session uses encryption or not.
func (session *MinecraftSession) UsesEncryption() bool {
	return session.usesEncryption
}

// EnableEncryption enables encryption for this session and computes secret key bytes.
func (session *MinecraftSession) EnableEncryption() {
	session.usesEncryption = true
	session.encryptionHandler.Data.ComputeSharedSecret()
	session.encryptionHandler.Data.ComputeSecretKeyBytes()
}

// IsXBOXLiveAuthenticated checks if the session logged in while being logged into XBOX Live.
func (session *MinecraftSession) IsXBOXLiveAuthenticated() bool {
	return session.xboxLiveAuthenticated
}

// SetXBOXLiveAuthenticated sets the session XBOX Live authenticated.
func (session *MinecraftSession) SetXBOXLiveAuthenticated(value bool) {
	session.xboxLiveAuthenticated = value
}

// SendMessage sends a text message to the Minecraft session.
func (session *MinecraftSession) SendMessage(message ...interface{}) {
	session.SendText(types.Text{Message: strings.Trim(fmt.Sprint(message), "[]")})
}

// GetPermissionGroup returns the permission group this session is in.
func (session *MinecraftSession) GetPermissionGroup() *permissions.Group {
	return session.permissionGroup
}

// SetPermissionGroup sets the permission group of this session.
func (session *MinecraftSession) SetPermissionGroup(group *permissions.Group) {
	session.permissionGroup = group
}

// HasPermission checks if this session has a permission.
func (session *MinecraftSession) HasPermission(permission string) bool {
	if session.GetPermissionGroup().HasPermission(permission) {
		return true
	}
	var _, exists = session.permissions[permission]
	return exists
}

// AddPermission adds a permission to the session.
// Returns true if a permission with the same name was overwritten.
func (session *MinecraftSession) AddPermission(permission *permissions.Permission) bool {
	var hasPermission = session.HasPermission(permission.GetName())

	session.permissions[permission.GetName()] = permission

	return hasPermission
}

// RemovePermission deletes a permission from the session.
// This does not delete the permission from the group the session is in.
func (session *MinecraftSession) RemovePermission(permission string) bool {
	if !session.HasPermission(permission) {
		return false
	}
	delete(session.permissions, permission)

	return true
}

func (session *MinecraftSession) SendSkin(target *MinecraftSession) {
	var player = session.GetPlayer()
	target.SendPlayerSkin(player.GetUUID(), player.GetSkinId(), player.GetGeometryName(), player.GetGeometryData(), player.GetSkinData(), player.GetCapeData())
}

// SendPacket sends a packet to this session.
func (session *MinecraftSession) SendPacket(packet packets.IPacket) {
	if session.session == nil {
		return
	}
	var b = NewMinecraftPacketBatch(session)
	b.AddPacket(packet)

	session.SendBatch(b)
}

// SendBatch sends a batch to this session.
func (session *MinecraftSession) SendBatch(batch *MinecraftPacketBatch) {
	if session.session == nil {
		return
	}
	session.session.SendPacket(batch, protocol.ReliabilityReliable, server.PriorityMedium)
}

// HandlePacket handles packets of this session.
func (session *MinecraftSession) HandlePacket(packet packets.IPacket) {
	priorityHandlers := session.adapter.packetManager.GetHandlersById(packet.GetId())

	var handled = false
handling:
	for _, h := range priorityHandlers {
		for _, iHandler := range h {
			if handler, ok := iHandler.(*PacketHandler); ok {
				if packet.IsDiscarded() {
					break handling
				}

				ret := handler.function(packet, session)
				if !handled {
					handled = ret
				}
			}
		}
	}
	if !handled {
		text.DefaultLogger.Debug("Unhandled Minecraft packet with ID:", packet.GetId())
	}
}

func (session *MinecraftSession) Close(reason string, hideDisconnectionScreen bool) {
	if session.Connected {
		loadedChunks := session.GetChunkLoader().GetLoadedChunks()

		for _, online := range session.adapter.sessionManager.GetSessions() {
			online.SendRemoveEntity(session.player.Entity.GetUniqueId())
			online.player.RemoveViewer(session)
		}

		for _, chunk := range loadedChunks {
			chunk.RemoveViewer(session)
			chunk.RemoveEntity(session.player.Entity.GetRuntimeId())
		}

		session.player.Close()
	}
	session.SendDisconnect(reason, hideDisconnectionScreen)
}

func (session *MinecraftSession) Kick(reason string, hideDisconnectionScreen bool, isAdmin bool) {
	if isAdmin {
		reason = "Kicked By Admin. Reason: " + reason
		session.Close(reason, hideDisconnectionScreen)
		text.DefaultLogger.Info(session.GetDisplayName() + " Disconnected.", reason)
	}else{
		session.Close(reason, hideDisconnectionScreen)
		text.DefaultLogger.Info(session.GetDisplayName() + " Disconnected.", reason)
	}
}

// SyncMove synchronizes the server's player movement with the client movement.
func (session *MinecraftSession) SyncMove(x, y, z float64, pitch, yaw, headYaw float64, onGround bool) {
	session.player.SyncMove(x, y, z, pitch, yaw, headYaw, onGround)
}

func (session *MinecraftSession) Tick() {
	if session.Connected {
		session.GetChunkLoader().Warp(session.GetPlayer().GetDimension(), int32(math.Floor(session.player.Position.X))>>4, int32(math.Floor(session.player.Position.Z))>>4)
		session.GetChunkLoader().Request(session.GetViewDistance(), 40)
	}
}
