package gomine

import (
	"errors"
	"os"
	"gomine/tasks"
	"gomine/utils"
	"gomine/resources"
	"gomine/worlds"
	"gomine/interfaces"
	"gomine/commands"
	"gomine/commands/defaults"
	"gomine/net"
	"gomine/net/info"
	"gomine/permissions"
	"gomine/players"
	"gomine/worlds/blocks"
	"gomine/worlds/generation"
	generatorDefaults "gomine/worlds/generation/defaults"
)

var levelId = 0

const (
	GoMineName = "GoMine"

	GoMineVersion = "0.0.1"
	ApiVersion = "0.0.1"
)

type Server struct {
	isRunning  bool
	tick int

	serverPath string
	scheduler  *tasks.Scheduler
	logger     interfaces.ILogger
	config 	   *resources.GoMineConfig
	consoleReader *ConsoleReader
	commandHolder interfaces.ICommandHolder
	permissionManager *permissions.PermissionManager

	levels map[int]interfaces.ILevel

	playerFactory *players.PlayerFactory

	rakLibAdapter *net.GoRakLibAdapter
}

/**
 * Creates a new server.
 * Will report an error if a server is already existent.
 */
func NewServer(serverPath string) *Server {
	var server = &Server{}
	server.serverPath = serverPath
	server.config = resources.NewGoMineConfig(serverPath)
	server.scheduler = tasks.NewScheduler()
	server.logger = utils.NewLogger(GoMineName, serverPath, server.GetConfiguration().DebugMode)
	server.levels = make(map[int]interfaces.ILevel)
	server.consoleReader = NewConsoleReader(server)
	server.commandHolder = commands.NewCommandHolder()
	server.rakLibAdapter = net.NewGoRakLibAdapter(server)

	server.playerFactory = players.NewPlayerFactory(server)

	server.permissionManager = permissions.NewPermissionManager(server)

	return server
}

/**
 * Registers all default commands.
 */
func (server *Server) RegisterDefaultCommands() {
	server.commandHolder.RegisterCommand(defaults.NewStop(server))
	server.commandHolder.RegisterCommand(defaults.NewTest(server))
}

/**
 * Registers all default commands.
 */
func (server *Server) RegisterGenerators() {
	generation.RegisterGenerator(generatorDefaults.NewFlatGenerator())
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

	server.RegisterDefaultCommands()
	blocks.InitBlockPool()
	generation.InitGeneratorList()

	server.LoadLevels()
	server.RegisterGenerators()

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

func (server *Server) LoadLevels() {
	server.LoadLevel(server.config.DefaultLevel)
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
		// server.GenerateLevel(level) We need file writing for this. TODO.
	}
	if server.IsLevelLoaded(levelName) {
		return false
	}
	server.levels[levelId] = worlds.NewLevel(levelName, levelId, server, make(map[int]interfaces.IChunk))
	levelId++
	return true
}

/**
 * Returns the default level and loads/generates it if needed.
 */
func (server *Server) GetDefaultLevel() interfaces.ILevel {
	if !server.IsLevelLoaded(server.config.DefaultLevel) {
		server.LoadLevel(server.config.DefaultLevel)
	}
	var level, _ = server.GetLevelByName(server.config.DefaultLevel)
	return level
}

/**
 * Returns a level by its ID. Returns an error if a level with the ID is not loaded.
 */
func (server *Server) GetLevelById(id int) (interfaces.ILevel, error) {
	var level interfaces.ILevel
	if level, ok := server.levels[id]; ok {
		return level, nil
	}
	return level, errors.New("level with given ID is not loaded")
}

/**
 * Returns a level by its name. Returns an error if the level is not loaded.
 */
func (server *Server) GetLevelByName(name string) (interfaces.ILevel, error) {
	var level interfaces.ILevel
	if !server.IsLevelGenerated(name) {
		return level, errors.New("level with given name is not generated")
	}
	if !server.IsLevelLoaded(name) {
		return level, errors.New("level with given name is not loaded")
	}
	for _, level := range server.GetLoadedLevels() {
		if level.GetName() == name {
			return level, nil
		}
	}
	return level, nil
}

/**
 * Returns the console command reader.
 */
func (server *Server) GetConsoleReader() *ConsoleReader {
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
 * Returns the GoMine Name.
 */
func (server *Server) GetName() string {
	return GoMineName
}

/**
 * Returns the name of the server specified in the configuration.
 */
func (server *Server) GetServerName() string {
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
func (server *Server) GetRakLibAdapter() interfaces.IGoRakLibAdapter {
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
 * Returns the player factory of the server.
 */
func (server *Server) GetPlayerFactory() interfaces.IPlayerFactory {
	return server.playerFactory
}

/**
 * Returns the current tick the server is on.
 */
func (server *Server) GetCurrentTick() int {
	return server.tick
}

/**
 * Internal. Not to be used by plugins.
 * Ticks the entire server. (Levels, scheduler, GoRakLib server etc.)
 */
func (server *Server) Tick(currentTick int) {
	server.tick = currentTick
	if !server.isRunning {
		return
	}
	server.GetScheduler().DoTick()
	for _, level := range server.levels {
		level.TickLevel()
	}

	for _, p := range server.playerFactory.GetPlayers() {
		p.Tick()
	}

	server.rakLibAdapter.Tick()

	server.rakLibAdapter.GetRakLibServer().SetConnectedSessionCount(server.GetPlayerFactory().GetPlayerCount())
}