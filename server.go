package gomine

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"encoding/hex"
	"errors"
	"fmt"
	"github.com/irmine/gomine/commands"
	"github.com/irmine/gomine/net"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets/data"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/packs"
	"github.com/irmine/gomine/permissions"
	"github.com/irmine/gomine/resources"
	"github.com/irmine/gomine/text"
	"github.com/irmine/goraklib/server"
	"github.com/irmine/query"
	"github.com/irmine/worlds"
	"github.com/irmine/worlds/providers"
	net2 "net"
	"os"
	"strings"
)

const (
	GoMineName    = "GoMine"
	GoMineVersion = "0.0.1"
)

type Server struct {
	isRunning         bool
	tick              int64
	privateKey        *ecdsa.PrivateKey
	token             []byte
	ServerPath        string
	Config            *resources.GoMineConfig
	CommandReader     *text.CommandReader
	CommandManager    *commands.Manager
	PackManager       *packs.Manager
	PermissionManager *permissions.Manager
	LevelManager      *worlds.Manager
	SessionManager    *net.SessionManager
	NetworkAdapter    *net.NetworkAdapter
	PluginManager     *PluginManager
	QueryManager      query.Manager
}

// AlreadyStarted gets returned during server startup,
// if the server has already been started.
var AlreadyStarted = errors.New("server is already started")

// NewServer returns a new server with the given server path.
func NewServer(serverPath string, config *resources.GoMineConfig) *Server {
	var s = &Server{}

	s.ServerPath = serverPath
	s.Config = config
	text.DefaultLogger.DebugMode = config.DebugMode
	file, _ := os.OpenFile(serverPath+"gomine.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0700)
	text.DefaultLogger.AddOutput(func(message []byte) {
		_, err := file.WriteString(text.ColoredString(message).StripAll())
		if err != nil {
			text.DefaultLogger.LogError(err)
		}
	})

	s.LevelManager = worlds.NewManager(serverPath)
	s.CommandReader = text.NewCommandReader(os.Stdin)
	s.CommandReader.AddReadFunc(s.attemptReadCommand)

	s.CommandManager = commands.NewManager()

	s.SessionManager = net.NewSessionManager()
	s.NetworkAdapter = net.NewNetworkAdapter(NewPacketManager(s), s.SessionManager)
	s.NetworkAdapter.GetRakLibManager().PongData = s.GeneratePongData()
	s.NetworkAdapter.GetRakLibManager().RawPacketFunction = s.HandleRaw
	s.NetworkAdapter.GetRakLibManager().DisconnectFunction = s.HandleDisconnect

	s.PackManager = packs.NewManager(serverPath)
	s.PermissionManager = permissions.NewManager()
	s.PluginManager = NewPluginManager(s)
	s.QueryManager = query.NewManager()

	if config.UseEncryption {
		var curve = elliptic.P384()

		var err error
		s.privateKey, err = ecdsa.GenerateKey(curve, rand.Reader)
		text.DefaultLogger.LogError(err)

		if !curve.IsOnCurve(s.privateKey.X, s.privateKey.Y) {
			text.DefaultLogger.Error("Invalid private key generated")
		}

		var token = make([]byte, 128)
		_, err = rand.Read(token)
		if err != nil {
			text.DefaultLogger.Error(err)
		}
		s.token = token
	}

	return s
}

// RegisterDefaultCommands registers all default commands of the server.
func (server *Server) RegisterDefaultCommands() {
	server.CommandManager.RegisterCommand(NewStop(server))
	server.CommandManager.RegisterCommand(NewList(server))
	server.CommandManager.RegisterCommand(NewPing())
	server.CommandManager.RegisterCommand(NewTest(server))
}

// IsRunning checks if the server is running.
func (server *Server) IsRunning() bool {
	return server.isRunning
}

// Start starts the server and loads levels, plugins, resource packs etc.
// Start returns an error if one occurred during starting.
func (server *Server) Start() error {
	if server.isRunning {
		return AlreadyStarted
	}
	text.DefaultLogger.Info("GoMine "+GoMineVersion+" is now starting...", "("+server.ServerPath+")")

	server.LevelManager.SetDefaultLevel(worlds.NewLevel("world", server.ServerPath))
	var dimension = worlds.NewDimension("overworld", server.LevelManager.GetDefaultLevel(), worlds.OverworldId)
	server.LevelManager.GetDefaultLevel().SetDefaultDimension(dimension)
	dimension.SetChunkProvider(providers.NewAnvil(server.ServerPath + "worlds/world/overworld/region/"))
	dimension.SetGenerator(Flat{})

	server.RegisterDefaultCommands()

	server.PackManager.LoadResourcePacks() // Behavior packs may depend on resource packs, so always load resource packs first.
	server.PackManager.LoadBehaviorPacks()

	server.PluginManager.LoadPlugins()

	server.isRunning = true
	return server.NetworkAdapter.GetRakLibManager().Start(server.Config.ServerIp, int(server.Config.ServerPort))
}

// Shutdown shuts down the server, saving and disabling everything.
func (server *Server) Shutdown() {
	if !server.isRunning {
		return
	}
	text.DefaultLogger.Info("Server is shutting down.")

	text.DefaultLogger.Notice("Server stopped.")
	text.DefaultLogger.Wait()

	server.isRunning = false
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

// HasPermission returns if the server has a given permission.
// Always returns true to satisfy the ICommandSender interface.
func (server *Server) HasPermission(string) bool {
	return true
}

// SendMessage sends a message to the server to satisfy the ICommandSender interface.
func (server *Server) SendMessage(message ...interface{}) {
	text.DefaultLogger.Notice(message)
}

// GetEngineName returns 'GoMine'.
func (server *Server) GetEngineName() string {
	return GoMineName
}

// GetName returns the LAN name of the server specified in the configuration.
func (server *Server) GetName() string {
	return server.Config.ServerName
}

// GetPort returns the port of the server specified in the configuration.
func (server *Server) GetPort() uint16 {
	return server.Config.ServerPort
}

// GetAddress returns the IP address specified in the configuration.
func (server *Server) GetAddress() string {
	return server.Config.ServerIp
}

// GetMaximumPlayers returns the maximum amount of players on the server.
func (server *Server) GetMaximumPlayers() uint {
	return server.Config.MaximumPlayers
}

// Returns the Message Of The Day of the server.
func (server *Server) GetMotd() string {
	return server.Config.ServerMotd
}

// GetCurrentTick returns the current tick the server is on.
func (server *Server) GetCurrentTick() int64 {
	return server.tick
}

// BroadcastMessageTo broadcasts a message to all receivers.
func (server *Server) BroadcastMessageTo(receivers []*net.MinecraftSession, message ...interface{}) {
	for _, session := range receivers {
		session.SendMessage(message)
	}
	text.DefaultLogger.LogChat(message)
}

// Broadcast broadcasts a message to all players and the console in the server.
func (server *Server) BroadcastMessage(message ...interface{}) {
	for _, session := range server.SessionManager.GetSessions() {
		session.SendMessage(message)
	}
	text.DefaultLogger.LogChat(message)
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
func (server *Server) GenerateQueryResult() query.Result {
	var plugs []string
	for _, plug := range server.PluginManager.GetPlugins() {
		plugs = append(plugs, plug.GetName()+" v"+plug.GetVersion())
	}

	var ps []string
	for name := range server.SessionManager.GetSessions() {
		ps = append(ps, name)
	}

	var result = query.Result{
		MOTD:           server.GetMotd(),
		ListPlugins:    server.Config.AllowPluginQuery,
		PluginNames:    plugs,
		PlayerNames:    ps,
		GameMode:       "SMP",
		Version:        server.GetMinecraftVersion(),
		ServerEngine:   server.GetEngineName(),
		WorldName:      server.LevelManager.GetDefaultLevel().GetName(),
		OnlinePlayers:  int(server.SessionManager.GetSessionCount()),
		MaximumPlayers: int(server.Config.MaximumPlayers),
		Whitelist:      "off",
		Port:           server.Config.ServerPort,
		Address:        server.Config.ServerIp,
	}

	return result
}

// HandleRaw handles a raw packet, for instance a query packet.
func (server *Server) HandleRaw(packet []byte, addr *net2.UDPAddr) {
	if string(packet[0:2]) == string(query.Header) {
		if !server.Config.AllowQuery {
			return
		}

		var q = query.NewFromRaw(packet, addr)
		q.DecodeServer()

		server.QueryManager.HandleQuery(q)
		return
	}
	text.DefaultLogger.Debug("Unhandled raw packet:", hex.EncodeToString(packet))
}

// HandleDisconnect handles a disconnection from a session.
func (server *Server) HandleDisconnect(s *server.Session) {
	text.DefaultLogger.Debug(s, "disconnected!")
	session, ok := server.SessionManager.GetSessionByRakNetSession(s)

	server.SessionManager.RemoveMinecraftSession(session)
	if !ok {
		return
	}

	if session.GetPlayer().Dimension != nil {
		for _, online := range server.SessionManager.GetSessions() {
			online.SendPlayerList(data.ListTypeRemove, map[string]protocol.PlayerListEntry{online.GetPlayer().GetName(): online.GetPlayer()})
		}

		session.GetPlayer().DespawnFromAll()

		session.GetPlayer().Close()

		server.BroadcastMessage(text.Yellow+session.GetDisplayName(), "has left the server")
	}
}

// GeneratePongData generates the GoRakLib pong data for the UnconnectedPong RakNet packet.
func (server *Server) GeneratePongData() string {
	return fmt.Sprint("MCPE;", server.GetMotd(), ";", info.LatestProtocol, ";", server.GetMinecraftNetworkVersion(), ";", server.SessionManager.GetSessionCount(), ";", server.Config.MaximumPlayers, ";", server.NetworkAdapter.GetRakLibManager().ServerId, ";", server.GetEngineName(), ";Creative;")
}

// Tick ticks the entire server. (Levels, scheduler, GoRakLib server etc.)
// Internal. Not to be used by plugins.
func (server *Server) Tick() {
	if !server.isRunning {
		return
	}
	if server.tick%20 == 0 {
		server.QueryManager.SetQueryResult(server.GenerateQueryResult())
		server.NetworkAdapter.GetRakLibManager().PongData = server.GeneratePongData()
	}

	for _, session := range server.SessionManager.GetSessions() {
		session.Tick()
	}

	for range server.LevelManager.GetLevels() {
		//level.Tick()
	}
	server.tick++
}

func (server *Server) attemptReadCommand(commandText string) {
	args := strings.Split(commandText, " ")
	commandName := args[0]
	i := 1
	for !server.CommandManager.IsCommandRegistered(commandName) {
		if i == len(args) {
			break
		}
		commandName += " " + args[i]
		i++
	}
	manager := server.CommandManager

	if !manager.IsCommandRegistered(commandName) {
		text.DefaultLogger.Error("Command could not be found.")
		return
	}
	args = args[i:]

	command, _ := manager.GetCommand(commandName)
	command.Execute(server, args)
}
