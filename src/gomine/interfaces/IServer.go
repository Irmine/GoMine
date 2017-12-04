package interfaces

import (
	"gomine/tasks"
	"gomine/resources"
	"gomine/worlds/generation"
)

type IServer interface {
	IsRunning() bool
	Start()
	Shutdown()
	GetTickRate() int
	SetTickRate(int)
	ResetTickRate()
	GetScheduler() *tasks.Scheduler
	GetServerPath() string
	GetLogger() ILogger
	GetConfiguration() *resources.GoMineConfig
	GetCommandHolder() ICommandHolder
	GetLoadedLevels() map[int]ILevel
	IsLevelLoaded(string) bool
	IsLevelGenerated(string) bool
	LoadLevel(string, generation.IGenerator) bool
	HasPermission(string) bool
	SendMessage(string)
	GetName() string
	GetAddress() string
	GetPort() uint16
	GetMaximumPlayers() uint
	GetMotd() string
	Tick()
	GetPermissionManager() IPermissionManager
	GetServerName() string
	GetVersion() string
	GetNetworkVersion() string
	GetRakLibAdapter() IGoRakLibAdapter
	GetPlayerFactory() IPlayerFactory
	GenerateLevel(ILevel)
	GetDefaultLevel() ILevel
	GetLevelById(int) (ILevel, error)
	GetLevelByName(string) (ILevel, error)
}
