package net

import (
	"goraklib/server"
	"gomine/utils"
	"gomine/interfaces"
	"goraklib/protocol"
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
	
	encryptionHandler *utils.EncryptionHandler
	usesEncryption bool
	xboxLiveAuthenticated bool

	initialized bool
}

func NewMinecraftSession(server interfaces.IServer, session *server.Session, protocol int32, version string, uuid utils.UUID, xuid string, clientId int) *MinecraftSession {
	return &MinecraftSession{
		server,
		session,
		uuid,
		xuid,
		clientId,
		protocol,
		version,
		"en_US",
		utils.NewEncryptionHandler(),
		false,
		false,
		true,
	}
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