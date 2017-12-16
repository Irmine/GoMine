package interfaces

import (
	"gomine/tasks"
	"gomine/resources"
)

type IServer interface {
	IsRunning() bool
	Start()
	Shutdown()
	GetScheduler() *tasks.Scheduler
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
	Tick(int)
	GetPermissionManager() IPermissionManager
	GetServerName() string
	GetVersion() string
	GetNetworkVersion() string
	GetRakLibAdapter() IGoRakLibAdapter
	GetPlayerFactory() IPlayerFactory
	GetPackHandler() IPackHandler
	GetDefaultLevel() ILevel
	GetLevelById(int) (ILevel, error)
	GetLevelByName(string) (ILevel, error)
	GetCurrentTick() int
}