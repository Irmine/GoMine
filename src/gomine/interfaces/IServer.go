package interfaces

import (
	"gomine/tasks"
	"gomine/utils"
	"gomine/resources"
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
	GetLogger() *utils.Logger
	GetConfiguration() *resources.GoMineConfig
	GetLoadedLevels() map[int]ILevel
	IsLevelLoaded(string) bool
	IsLevelGenerated(string) bool
	LoadLevel(string) bool
	Tick(int)
}
