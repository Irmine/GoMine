package gomine

import (
	"errors"
	"os"
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
	"gomine/packs"
	"gomine/plugins"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"gomine/net/query"
	"goraklib/server"
	"strings"
)

var levelId = 0

const (
	GoMineName = "GoMine"
	GoMineVersion = "0.0.1"
)

type Server struct {
	isRunning  bool
	tick int64
	privateKey *ecdsa.PrivateKey
	token []byte

	serverPath string
	logger     interfaces.ILogger
	config 	   *resources.GoMineConfig
	consoleReader *ConsoleReader
	commandHolder interfaces.ICommandHolder

	packHandler *packs.PackHandler
	permissionManager *permissions.PermissionManager

	levels map[int]interfaces.ILevel

	playerFactory *players.PlayerFactory

	rakLibAdapter *net.GoRakLibAdapter

	pluginManager *plugins.PluginManager

	queryManager query.QueryManager
}

/**
 * Creates a new server.
 * Will report an error if a server is already existent.
 */
func NewServer(serverPath string) *Server {
	var server = &Server{}

	server.serverPath = serverPath
	server.config = resources.NewGoMineConfig(serverPath)
	server.logger = utils.NewLogger(GoMineName, serverPath, server.GetConfiguration().DebugMode)
	server.levels = make(map[int]interfaces.ILevel)
	server.consoleReader = NewConsoleReader(server)
	server.commandHolder = commands.NewCommandHolder()
	server.rakLibAdapter = net.NewGoRakLibAdapter(server)

	server.packHandler = packs.NewPackHandler(server)

	server.playerFactory = players.NewPlayerFactory(server)
	server.permissionManager = permissions.NewPermissionManager(server)

	server.pluginManager = plugins.NewPluginManager(server)

	server.queryManager = query.NewQueryManager(server)

	if server.config.UseEncryption {
		var curve = elliptic.P384()

		var err error
		server.privateKey, err = ecdsa.GenerateKey(curve, rand.Reader)
		server.logger.LogError(err)

		if !curve.IsOnCurve(server.privateKey.X, server.privateKey.Y) {
			server.logger.Error("Invalid private key generated")
		}

		var token = make([]byte, 128)
		rand.Read(token)
		server.token = token
	}

	return server
}

/**
 * Registers all default commands.
 */
func (server *Server) RegisterDefaultCommands() {
	server.commandHolder.RegisterCommand(defaults.NewStop(server))
	server.commandHolder.RegisterCommand(defaults.NewList(server))
	server.commandHolder.RegisterCommand(defaults.NewTest(server))
	server.commandHolder.RegisterCommand(defaults.NewPing())
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

	server.LoadLevels()

	server.packHandler.LoadResourcePacks() // Behavior packs may depend on resource packs, so always load resource packs first.
	server.packHandler.LoadBehaviorPacks()

	server.pluginManager.LoadPlugins()

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

	server.GetLogger().Notice("Server stopped.")
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
func (server *Server) GetEngineName() string {
	return GoMineName
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
func (server *Server) GetCurrentTick() int64 {
	return server.tick
}

/**
 * Returns the resource and behavior pack handler.
 */
func (server *Server) GetPackHandler() interfaces.IPackHandler {
	return server.packHandler
}

/**
 * Returns the plugin manager of the server.
 */
func (server *Server) GetPluginManager() *plugins.PluginManager {
	return server.pluginManager
}

/**
 * Broadcasts a message to all receivers.
 */
func (server *Server) BroadcastMessageTo(message string, receivers []interfaces.IPlayer) {
	for _, player := range receivers {
		player.SendMessage(message)
	}
	server.logger.LogChat(message)
}

/**
 * Broadcasts a message to all players and the console in the server.
 */
func (server *Server) BroadcastMessage(message string) {
	for _, player := range server.GetPlayerFactory().GetPlayers() {
		player.SendMessage(message)
	}
	server.logger.LogChat(message)
}

/**
 * Returns the ECDSA private key of the server.
 */
func (server *Server) GetPrivateKey() *ecdsa.PrivateKey {
	return server.privateKey
}

/**
 * Returns the ECDSA public key of the private key of the server.
 */
func (server *Server) GetPublicKey() *ecdsa.PublicKey {
	return &server.privateKey.PublicKey
}

/**
 * Returns the server token byte sequence.
 */
func (server *Server) GetServerToken() []byte {
	return server.token
}

/**
 * Returns the query data of the server in a byte array.
 */
func (server *Server) GenerateQueryResult(shortData bool) []byte {
	var plugs []string
	for _, plug := range server.pluginManager.GetPlugins() {
		plugs = append(plugs, plug.GetName() + " v" + plug.GetVersion())
	}

	var ps []string
	for _, player := range server.playerFactory.GetPlayers() {
		ps = append(ps, player.GetDisplayName())
	}

	var result = query.QueryResult{
		MOTD: server.GetMotd(),
		ListPlugins: server.config.AllowPluginQuery,
		PluginNames: plugs,
		PlayerNames: ps,
		GameMode: "SMP",
		Version: server.GetVersion(),
		ServerEngine: server.GetEngineName(),
		WorldName: server.GetDefaultLevel().GetName(),
		OnlinePlayers: int(server.GetPlayerFactory().GetPlayerCount()),
		MaximumPlayers: int(server.config.MaximumPlayers),
		Whitelist: "off",
		Port: server.config.ServerPort,
		Address: server.config.ServerIp,
	}

	if shortData {
		return result.GetShort()
	}
	return result.GetLong()
}

/**
 * Handles a raw packet, for instance a query packet.
 */
func (server *Server) HandleRaw(packet server.RawPacket) {
	if string(packet.Buffer[0:2]) == string(query.QueryHeader) {
		if !server.config.AllowQuery {
			return
		}

		var q = query.NewQueryFromRaw(packet)
		q.DecodeServer()

		server.queryManager.HandleQuery(q)
	}
}

/**
 * Internal. Not to be used by plugins.
 * Ticks the entire server. (Levels, scheduler, GoRakLib server etc.)
 */
func (server *Server) Tick(currentTick int64) {
	server.tick = currentTick
	if !server.isRunning {
		return
	}
	for _, level := range server.levels {
		level.TickLevel()
	}

	for _, p := range server.playerFactory.GetPlayers() {
		p.Tick()
	}

	server.rakLibAdapter.Tick()

	server.rakLibAdapter.GetRakLibServer().SetConnectedSessionCount(server.GetPlayerFactory().GetPlayerCount())
}