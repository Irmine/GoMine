package interfaces

import (
	"gomine/resources"
	"crypto/ecdsa"
	"goraklib/server"
)

type IServer interface {
	IsRunning() bool
	Start()
	Shutdown()
	GetServerPath() string
	GetLogger() ILogger
	GetConfiguration() *resources.GoMineConfig
	GetCommandHolder() ICommandHolder
	GetLoadedLevels() map[int]ILevel
	IsLevelLoaded(string) bool
	IsLevelGenerated(string) bool
	LoadLevel(string) bool
	HasPermission(string) bool
	SendMessage(string)
	GetName() string
	GetAddress() string
	GetPort() uint16
	GetMaximumPlayers() uint
	GetMotd() string
	Tick(int64)
	GetPermissionManager() IPermissionManager
	GetEngineName() string
	GetMinecraftVersion() string
	GetMinecraftNetworkVersion() string
	GetNetworkAdapter() INetworkAdapter
	GetPlayerFactory() IPlayerFactory
	GetPackHandler() IPackHandler
	GetDefaultLevel() ILevel
	GetLevelById(int) (ILevel, error)
	GetLevelByName(string) (ILevel, error)
	GetCurrentTick() int64
	BroadcastMessageTo(message string, receivers []IPlayer)
	BroadcastMessage(message string)
	GetPrivateKey() *ecdsa.PrivateKey
	GetPublicKey() *ecdsa.PublicKey
	GetServerToken() []byte
	HandleRaw(server.RawPacket)
	GenerateQueryResult(bool) []byte
}
