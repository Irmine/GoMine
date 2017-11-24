package gomine

import (
	"errors"
	"gomine/tasks"
	"gomine/utils"
	"gomine/resources"
	"gomine/worlds"
	"os"
	"gomine/interfaces"
	"gomine/commands"
	"gomine/commands/defaults"
	"gomine/net"
	"gomine/players"
	"gomine/net/info"
	"gomine/permissions"
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
	logger     interfaces.ILogger
	config 	   *resources.GoMineConfig
	consoleReader *utils.ConsoleReader
	commandHolder interfaces.ICommandHolder

	permissionManager *permissions.PermissionManager

	levels map[int]interfaces.ILevel
	players map[string]players.Player

	rakLibAdapter *net.GoRakLibAdapter
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
	server.consoleReader = utils.NewConsoleReader()
	server.commandHolder = commands.NewCommandHolder()
	server.rakLibAdapter = net.NewGoRakLibAdapter(server)

	server.permissionManager = permissions.NewPermissionManager(server)

	server.RegisterDefaultCommands()

	started = true

	return server, nil
}

/**
 * Registers all default commands.
 */
func (server *Server) RegisterDefaultCommands() {
	server.commandHolder.RegisterCommand(defaults.NewStop(server))
	server.commandHolder.RegisterCommand(defaults.NewTest(server))
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
 * Shuts down the server if it is running.
 */
func (server *Server) Shutdown() {
	if !server.isRunning {
		return
	}
	server.GetLogger().Info("Server is shutting down.")

	server.isRunning = false
}

/**
 * Returns the server version prefixed with 'v'.
 * EG: "v1.2.6.2"
 */
func (server *Server) GetVersion() string {
	return info.GameVersion
}

/**
 * Returns the server version used for networking.
 * This version string is not prefixed with a 'v'.
 */
func (server *Server) GetNetworkVersion() string {
	return info.GameVersionNetwork
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
func (server *Server) GetLogger() interfaces.ILogger {
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
	var _, err = os.Stat(path)
	if err != nil {
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
	levels[counter] = worlds.NewLevel(levelName, server, []interfaces.IChunk{})
	counter++
	return true
}

/**
 * Returns the console command reader.
 */
func (server *Server) GetConsoleReader() *utils.ConsoleReader {
	return server.consoleReader
}

/**
 * Returns the command holder.
 */
func (server *Server) GetCommandHolder() interfaces.ICommandHolder {
	return server.commandHolder
}

/**
 * Returns if the server has a given permission.
 * Always returns true to satisfy the ICommandSender interface.
 */
func (server *Server) HasPermission(string) bool {
	return true
}

/**
 * Sends a message to the server to satisfy the ICommandSender interface.
 */
func (server *Server) SendMessage(message string) {
	server.GetLogger().Notice(message)
}

/**
 * Returns the name of the server specified in the configuration.
 */
func (server *Server) GetName() string {
	return server.config.ServerName
}

/**
 * Returns the port of the server specified in the configuration.
 */
func (server *Server) GetPort() uint16 {
	return server.config.ServerPort
}

/**
 * Returns the IP address specified in the configuration.
 */
func (server *Server) GetAddress() string {
	return server.config.ServerIp
}

/**
 * Returns the maximum amount of players on the server.
 */
func (server *Server) GetMaximumPlayers() uint {
	return server.config.MaximumPlayers
}

/**
 * Returns the GoRakLibAdapter of the server.
 * This is used for network features.
 */
func (server *Server) GetRakLibAdapter() *net.GoRakLibAdapter {
	return server.rakLibAdapter
}

/**
 * Returns the Message Of The Day of the server.
 */
func (server *Server) GetMotd() string {
	return server.config.ServerMotd
}

/**
 * Returns the permission manager of the server.
 */
func (server *Server) GetPermissionManager() interfaces.IPermissionManager {
	return server.permissionManager
}

/**
 * Internal. Not to be used by plugins.
 * Ticks the entire server. (Levels, scheduler, GoRakLib server etc.)
 */
func (server *Server) Tick(currentTick int) {
	server.GetScheduler().DoTick()
	for _, level := range server.levels {
		level.TickLevel()
	}
	go server.consoleReader.ReadLine(server)
	server.rakLibAdapter.Tick()
}
