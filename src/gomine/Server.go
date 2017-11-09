package gomine

import (
	"errors"
	"gomine/tasks"
	"gomine/utils"
	"gomine/resources"
	"gomine/worlds"
	"os"
	"gomine/interfaces"
)

const (
	GoMineVersion = "0.0.1"
	ApiVersion = "0.0.1"
)

type Server struct {
	isRunning  bool
	tickRate   int
	serverPath string
	scheduler  *tasks.Scheduler
	logger     *utils.Logger
	config 	   *resources.GoMineConfig

	levels map[int]interfaces.ILevel
}

var started bool = false

var counter int = 0

/**
 * Creates a new server.
 * Will report an error if a server is already existent.
 */
func NewServer(serverPath string) (*Server, error) {
	var errorServer Server
	if started {
		return &errorServer, errors.New("cannot create a second server")
	}

	var server = &Server{}
	server.tickRate = 20
	server.serverPath = serverPath
	server.config = resources.NewGoMineConfig(serverPath)
	server.scheduler = tasks.NewScheduler()
	server.logger = utils.NewLogger("GoMine", serverPath, server.GetConfiguration().DebugMode)
	server.levels = make(map[int]interfaces.ILevel)

	return server, nil
}

/**
 * Returns whether the server is running or not.
 */
func (server *Server) IsRunning() bool {
	return server.isRunning
}

/**
 * Starts the server.
 */
func (server *Server) Start() {
	server.GetLogger().Info("GoMine " + GoMineVersion + " is now starting...")

	server.isRunning = true
}

/**
 * Shuts down the server.
 */
func (server *Server) Shutdown() {
	server.GetLogger().Info("Server is shutting down.")

	server.isRunning = false
}

/**
 * Returns the tick rate of the server.
 */
func (server *Server) GetTickRate() int {
	return server.tickRate
}

/**
 * Resets the tick value back to the default. (20)
 */
func (server *Server) ResetTickRate() {
	server.tickRate = 20
}

/**
 * Internal. Not to be used by plugins.
 */
func (server *Server) SetTickRate(tickRate int) {
	server.tickRate = tickRate
}

/**
 * Returns the scheduler used for scheduling tasks.
 */
func (server *Server) GetScheduler() *tasks.Scheduler {
	return server.scheduler
}

/**
 * Returns the path the src folder is located in.
 */
func (server *Server) GetServerPath() string {
	return server.serverPath
}

/**
 * Returns the server logger. Logs with a [GoMine] prefix.
 */
func (server *Server) GetLogger() *utils.Logger {
	return server.logger
}

/**
 * Returns the configuration of GoMine.
 */
func (server *Server) GetConfiguration() *resources.GoMineConfig {
	return server.config
}

/**
 * Returns all loaded levels in the server.
 */
func (server *Server) GetLoadedLevels() map[int]interfaces.ILevel {
	return server.levels
}

/**
 * Returns whether a level is loaded or not.
 */
func (server *Server) IsLevelLoaded(levelName string) bool {
	for _, level := range server.levels  {
		if level.GetName() == levelName {
			return true
		}
	}
	return false
}

/**
 * Returns whether a level is generated or not. (Includes loaded levels)
 */
func (server *Server) IsLevelGenerated(levelName string) bool {
	if server.IsLevelLoaded(levelName) {
		return true
	}
	var path = server.GetServerPath() + "worlds/" + levelName
	var _, error = os.Stat(path)
	if error != nil {
		return false
	}
	return true
}

/**
 * Loads a generated world. Returns true if the level was loaded successfully.
 */
func (server *Server) LoadLevel(levelName string) bool {
	if !server.IsLevelGenerated(levelName) {
		return false
	}
	if server.IsLevelLoaded(levelName) {
		return false
	}
	var levels = server.levels
	levels[counter] = worlds.NewLevel(levelName, server)
	counter++
	return true
}

/**
 * Internal. Not to be used by plugins.
 * Ticks the entire server. (Entities, block entities etc.)
 */
func (server *Server) Tick(currentTick int) {
	server.GetScheduler().DoTick()
}
