package interfaces

import (
	"crypto/ecdsa"

	"github.com/irmine/gomine/resources"
	"github.com/irmine/goraklib/server"
	"github.com/irmine/gomine/packs"
	"github.com/irmine/gomine/permissions"
	"github.com/irmine/gomine/commands"
)

type IServer interface {
	IsRunning() bool
	Start()
	Shutdown()
	GetServerPath() string
	GetLogger() ILogger
	GetConfiguration() *resources.GoMineConfig
	GetCommandManager() *commands.Manager
	GetLoadedLevels() map[int]ILevel
	IsLevelLoaded(string) bool
	IsLevelGenerated(string) bool
	LoadLevel(string) bool
	HasPermission(string) bool
	SendMessage(...interface{})
	GetName() string
	GetAddress() string
	GetPort() uint16
	GetMaximumPlayers() uint
	GetMotd() string
	Tick(int64)
	GetPermissionManager() *permissions.Manager
	GetEngineName() string
	GetMinecraftVersion() string
	GetMinecraftNetworkVersion() string
	GetNetworkAdapter() INetworkAdapter
	GetPlayerFactory() IPlayerFactory
	GetPackManager() *packs.Manager
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
