package gomine

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"os"

	"github.com/irmine/goraklib/server"

	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/commands/defaults"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/query"
	"github.com/irmine/gomine/packs"
	"github.com/irmine/gomine/permissions"
	"github.com/irmine/gomine/players"
	"github.com/irmine/gomine/plugins"
	"github.com/irmine/gomine/resources"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/gomine/worlds"
)

var levelId = 0

const (
	GoMineName    = "GoMine"
	GoMineVersion = "0.0.1"
)

type Server struct {
	isRunning         bool
	tick              int64
	privateKey        *ecdsa.PrivateKey
	token             []byte
	serverPath        string
	logger            interfaces.ILogger
	config            *resources.GoMineConfig
	consoleReader     *ConsoleReader
	commandHolder     *commands.Manager
	packManager       *packs.Manager
	permissionManager *permissions.Manager
	levels            map[int]interfaces.ILevel
	playerFactory     *players.PlayerFactory
	networkAdapter    *net.NetworkAdapter
	pluginManager     *plugins.PluginManager
	queryManager      query.QueryManager
}

// NewServer returns a new server with the given server path.
func NewServer(serverPath string) *Server {
	var s = &Server{}

	s.serverPath = serverPath
	s.config = resources.NewGoMineConfig(serverPath)
	s.logger = utils.NewLogger(GoMineName, serverPath, s.GetConfiguration().DebugMode)
	s.levels = make(map[int]interfaces.ILevel)
	s.consoleReader = NewConsoleReader(s)
	s.commandHolder = commands.NewManager()
	s.networkAdapter = net.NewNetworkAdapter(s)

	s.packManager = packs.NewManager(serverPath)

	s.playerFactory = players.NewPlayerFactory(s)
	s.permissionManager = permissions.NewManager()

	s.pluginManager = plugins.NewPluginManager(s)

	s.queryManager = query.NewQueryManager(s)

	if s.config.UseEncryption {
		var curve = elliptic.P384()

		var err error
		s.privateKey, err = ecdsa.GenerateKey(curve, rand.Reader)
		s.logger.LogError(err)

		if !curve.IsOnCurve(s.privateKey.X, s.privateKey.Y) {
			s.logger.Error("Invalid private key generated")
		}

		var token = make([]byte, 128)
		rand.Read(token)
		s.token = token
	}

	return s
}

// RegisterDefaultCommands registers all default commands of the server.
func (server *Server) RegisterDefaultCommands() {
	server.commandHolder.RegisterCommand(defaults.NewStop(server))
	server.commandHolder.RegisterCommand(defaults.NewList(server))
	server.commandHolder.RegisterCommand(defaults.NewTest())
	server.commandHolder.RegisterCommand(defaults.NewPing())
}

// IsRunning checks if the server is running.
func (server *Server) IsRunning() bool {
	return server.isRunning
}

// Start starts the server and loads levels, plugins, resource packs etc.
func (server *Server) Start() {
	if server.isRunning {
		return
	}
	server.GetLogger().Info("GoMine " + GoMineVersion + " is now starting...")

	server.RegisterDefaultCommands()

	server.LoadLevels()

	server.packManager.LoadResourcePacks() // Behavior packs may depend on resource packs, so always load resource packs first.
	server.packManager.LoadBehaviorPacks()

	server.pluginManager.LoadPlugins()

	server.isRunning = true
}

// Shutdown shuts down the server, saving and disabling everything.
func (server *Server) Shutdown() {
	if !server.isRunning {
		return
	}
	server.GetLogger().Info("Server is shutting down.")

	server.isRunning = false

	server.GetLogger().Notice("Server stopped.")
}

// GetMinecraftVersion returns the latest Minecraft game version.
// It is prefixed with a 'v', for example: "v1.2.10.1"
func (server *Server) GetMinecraftVersion() string {
	return info.LatestGameVersion
}

// GetMinecraftNetworkVersion returns the latest Minecraft network version.
// For example: "1.2.10.1"
func (server *Server) GetMinecraftNetworkVersion() string {
	return info.LatestGameVersionNetwork
}

// GetServerPath returns the server path the server is installed in.
func (server *Server) GetServerPath() string {
	return server.serverPath
}

// GetLogger returns the server's logger.
// It is prefixed by a [GoMine] prefix.
func (server *Server) GetLogger() interfaces.ILogger {
	return server.logger
}

// GetConfiguration returns the configuration file of GoMine.
func (server *Server) GetConfiguration() *resources.GoMineConfig {
	return server.config
}

// GetLoadedLevels returns all loaded levels of the server.
func (server *Server) GetLoadedLevels() map[int]interfaces.ILevel {
	return server.levels
}

func (server *Server) LoadLevels() {
	server.LoadLevel(server.config.DefaultLevel)
}

// IsLevelLoaded returns whether a level is loaded or not.
func (server *Server) IsLevelLoaded(levelName string) bool {
	for _, level := range server.levels {
		if level.GetName() == levelName {
			return true
		}
	}
	return false
}

// IsLevelGenerated returns whether a level is generated or not. (Includes loaded levels)
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

// LoadLevel loads a generated level. Returns true if the level was loaded successfully.
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

// GetDefaultLevel returns the default level and loads/generates it if needed.
func (server *Server) GetDefaultLevel() interfaces.ILevel {
	if !server.IsLevelLoaded(server.config.DefaultLevel) {
		server.LoadLevel(server.config.DefaultLevel)
	}
	var level, _ = server.GetLevelByName(server.config.DefaultLevel)
	return level
}

// GetLevelById returns a level by its ID. Returns an error if a level with the ID is not loaded.
func (server *Server) GetLevelById(id int) (interfaces.ILevel, error) {
	var level interfaces.ILevel
	if level, ok := server.levels[id]; ok {
		return level, nil
	}
	return level, errors.New("level with given ID is not loaded")
}

// GetLevelByName returns a level by its name. Returns an error if the level is not loaded.
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

// GetConsoleReader returns the console command reader.
func (server *Server) GetConsoleReader() *ConsoleReader {
	return server.consoleReader
}

// GetCommandHolder returns the command manager.
func (server *Server) GetCommandManager() *commands.Manager {
	return server.commandHolder
}

// HasPermission returns if the server has a given permission.
// Always returns true to satisfy the ICommandSender interface.
func (server *Server) HasPermission(string) bool {
	return true
}

// SendMessage sends a message to the server to satisfy the ICommandSender interface.
func (server *Server) SendMessage(message ...interface{}) {
	server.GetLogger().Notice(message)
}

// GetEngineName returns 'GoMine'.
func (server *Server) GetEngineName() string {
	return GoMineName
}

// GetName returns the LAN name of the server specified in the configuration.
func (server *Server) GetName() string {
	return server.config.ServerName
}

// GetPort returns the port of the server specified in the configuration.
func (server *Server) GetPort() uint16 {
	return server.config.ServerPort
}

// GetAddress returns the IP address specified in the configuration.
func (server *Server) GetAddress() string {
	return server.config.ServerIp
}

// GetMaximumPlayers returns the maximum amount of players on the server.
func (server *Server) GetMaximumPlayers() uint {
	return server.config.MaximumPlayers
}

// GetNetworkAdapter returns the NetworkAdapter of the server.
func (server *Server) GetNetworkAdapter() interfaces.INetworkAdapter {
	return server.networkAdapter
}

// Returns the Message Of The Day of the server.

func (server *Server) GetMotd() string {
	return server.config.ServerMotd
}

// GetPermissionManager returns the permission manager of the server.
func (server *Server) GetPermissionManager() *permissions.Manager {
	return server.permissionManager
}

// GetPlayerFactory returns the player factory of the server.
func (server *Server) GetPlayerFactory() interfaces.IPlayerFactory {
	return server.playerFactory
}

// GetCurrentTick returns the current tick the server is on.
func (server *Server) GetCurrentTick() int64 {
	return server.tick
}

// GetPackManager returns the resource and behavior pack manager.
func (server *Server) GetPackManager() *packs.Manager {
	return server.packManager
}

// GetPluginManager returns the plugin manager of the server.
func (server *Server) GetPluginManager() *plugins.PluginManager {
	return server.pluginManager
}

// BroadcastMessageTo broadcasts a message to all receivers.
func (server *Server) BroadcastMessageTo(message string, receivers []interfaces.IPlayer) {
	for _, player := range receivers {
		player.SendMessage(message)
	}
	server.logger.LogChat(message)
}

// Broadcast broadcasts a message to all players and the console in the server.
func (server *Server) BroadcastMessage(message string) {
	for _, player := range server.GetPlayerFactory().GetPlayers() {
		player.SendMessage(message)
	}
	server.logger.LogChat(message)
}

// GetPrivateKey returns the ECDSA private key of the server.
func (server *Server) GetPrivateKey() *ecdsa.PrivateKey {
	return server.privateKey
}

// GetPublicKey returns the ECDSA public key of the private key of the server.
func (server *Server) GetPublicKey() *ecdsa.PublicKey {
	return &server.privateKey.PublicKey
}

// GetServerToken returns the server token byte sequence.
func (server *Server) GetServerToken() []byte {
	return server.token
}

// GenerateQueryResult returns the query data of the server in a byte array.
func (server *Server) GenerateQueryResult(shortData bool) []byte {
	var plugs []string
	for _, plug := range server.pluginManager.GetPlugins() {
		plugs = append(plugs, plug.GetName()+" v"+plug.GetVersion())
	}

	var ps []string
	for _, player := range server.playerFactory.GetPlayers() {
		ps = append(ps, player.GetDisplayName())
	}

	var result = query.QueryResult{
		MOTD:           server.GetMotd(),
		ListPlugins:    server.config.AllowPluginQuery,
		PluginNames:    plugs,
		PlayerNames:    ps,
		GameMode:       "SMP",
		Version:        server.GetMinecraftVersion(),
		ServerEngine:   server.GetEngineName(),
		WorldName:      server.GetDefaultLevel().GetName(),
		OnlinePlayers:  int(server.GetPlayerFactory().GetPlayerCount()),
		MaximumPlayers: int(server.config.MaximumPlayers),
		Whitelist:      "off",
		Port:           server.config.ServerPort,
		Address:        server.config.ServerIp,
	}

	if shortData {
		return result.GetShort()
	}
	return result.GetLong()
}

// HandleRaw handles a raw packet, for instance a query packet.
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

// Tick ticks the entire server. (Levels, scheduler, GoRakLib server etc.)
// Internal. Not to be used by plugins.
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

	server.networkAdapter.Tick()

	server.networkAdapter.GetRakLibServer().SetConnectedSessionCount(server.GetPlayerFactory().GetPlayerCount())
}
