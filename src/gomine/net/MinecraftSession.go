package net

import (
	"goraklib/server"
	"gomine/utils"
	"gomine/interfaces"
	"goraklib/protocol"
	"gomine/net/packets"
)

type MinecraftSession struct {
	server interfaces.IServer
	session *server.Session
	uuid utils.UUID
	xuid string
	clientId int
	minecraftProtocol int32
	minecraftVersion string
	language string

	clientPlatform int32
	
	encryptionHandler *utils.EncryptionHandler
	usesEncryption bool
	xboxLiveAuthenticated bool

	initialized bool
}

func NewMinecraftSession(server interfaces.IServer, session *server.Session, pk *packets.LoginPacket) *MinecraftSession {
	return &MinecraftSession{
		server,
		session,
		pk.ClientUUID,
		pk.ClientXUID,
		pk.ClientId,
		pk.Protocol,
		pk.ClientData.GameVersion,
		pk.Language,
		int32(pk.ClientData.DeviceOS),
		utils.NewEncryptionHandler(),
		false,
		false,
		true,
	}
}

/**
 * Returns the platform the client uses to player the game.
 */
func (session *MinecraftSession) GetPlatform() int32 {
	return session.clientPlatform
}

/**
 * Returns the protocol the client used to join the server.
 */
func (session *MinecraftSession) GetProtocol() int32 {
	return session.minecraftProtocol
}

/**
 * Sets the protocol of this minecraft session.
 */
func (session *MinecraftSession) SetProtocol(protocol int32) {
	session.minecraftProtocol = protocol
}

/**
 * Returns the Minecraft version the player used to join the server.
 */
func (session *MinecraftSession) GetGameVersion() string {
	return session.minecraftVersion
}

/**
 * Returns the main GoMine server.
 */
func (session *MinecraftSession) GetServer() interfaces.IServer {
	return session.server
}

/**
 * Returns the GoRakLib session of this session.
 */
func (session *MinecraftSession) GetSession() *server.Session {
	return session.session
}

/**
 * Returns the ping of the session in milliseconds.
 */
func (session *MinecraftSession) GetPing() uint64 {
	return session.session.GetPing()
}

/**
 * Returns the UUID of this player.
 */
func (session *MinecraftSession) GetUUID() utils.UUID {
	return session.uuid
}

/**
 * Returns the XUID of this session.
 */
func (session *MinecraftSession) GetXUID() string {
	return session.xuid
}

/**
 * Sets the language (locale) of this session.
 */
func (session *MinecraftSession) SetLanguage(language string) {
	session.language = language
}

/**
 * Returns the language (locale) of this session.
 */
func (session *MinecraftSession) GetLanguage() string {
	return session.language
}

/**
 * Returns the client ID of this session.
 */
func (session *MinecraftSession) GetClientId() int {
	return session.clientId
}

/**
 * Returns the handler used for encryption.
 */
func (session *MinecraftSession) GetEncryptionHandler() *utils.EncryptionHandler {
	return session.encryptionHandler
}

/**
 * Checks if the session uses encryption or not.
 */
func (session *MinecraftSession) UsesEncryption() bool {
	return session.usesEncryption
}

/**
 * Enables encryption for this session and computes secret key bytes.
 */
func (session *MinecraftSession) EnableEncryption() {
	session.usesEncryption = true
	session.encryptionHandler.Data.ComputeSharedSecret()
	session.encryptionHandler.Data.ComputeSecretKeyBytes()
}

/**
 * Checks if the session logged in while being logged into XBOX Live.
 */
func (session *MinecraftSession) IsXBOXLiveAuthenticated() bool {
	return session.xboxLiveAuthenticated
}

/**
 * Sets the session XBOX Live authenticated.
 */
func (session *MinecraftSession) SetXBOXLiveAuthenticated(value bool) {
	session.xboxLiveAuthenticated = value
}

/**
 * Sends a packet to this session.
 */
func (session *MinecraftSession) SendPacket(packet interfaces.IPacket) {
	var b = NewMinecraftPacketBatch(session, session.server.GetLogger())
	b.AddPacket(packet)

	session.SendBatch(b)
}

/**
 * Sends a batch to this session.
 */
func (session *MinecraftSession) SendBatch(batch interfaces.IMinecraftPacketBatch) {
	session.session.SendConnectedPacket(batch, protocol.ReliabilityReliableOrdered, server.PriorityMedium)
}

/**
 * Checks if the session is initialized.
 */
func (session *MinecraftSession) IsInitialized() bool {
	return session.initialized
}

/**
 * Handles packets after the initial LoginPacket.
 */
func (session *MinecraftSession) HandlePacket(packet interfaces.IPacket, player interfaces.IPlayer) {
	priorityHandlers := GetPacketHandlers(packet.GetId())

	var handled = false
	handling:
		for _, h := range priorityHandlers {
			for _, handler := range h {
				if packet.IsDiscarded() {
					break handling
				}

				ret := handler.Handle(packet, player, session.session, session.server)
				if !handled {
					handled = ret
				}
			}
		}
	if !handled {
		session.server.GetLogger().Debug("Unhandled Minecraft packet with ID:", packet.GetId())
	}
}