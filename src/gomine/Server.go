package gomine

import (
	"errors"
	"gomine/tasks"
	"gomine/utils"
)

type Server struct {
	isRunning  bool
	tickRate   int
	serverPath string
	scheduler  tasks.Scheduler
	logger     utils.Logger
}

var started bool = false

/**
 * Creates a new server.
 * Will report an error if a server is already existent.
 */
func NewServer(serverPath string) (Server, error) {
	var errorServer Server
	if started {
		return errorServer, errors.New("cannot create a second server")
	}

	var server = Server{
		false,
		20,
		serverPath,
		tasks.NewScheduler(),
		utils.NewLogger("GoMine", serverPath),
	}

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
	server.GetLogger().Log("Server is starting...", utils.Info)

	server.isRunning = true
}

/**
 * Shuts down the server.
 */
func (server *Server) Shutdown() {
	server.GetLogger().Log("Server is shutting down.", utils.Info)

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
	return &server.scheduler
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
func (server *Server) GetLogger() utils.Logger {
	return server.logger
}

/**
 * Internal. Not to be used by plugins.
 * Ticks the entire server. (Entities, block entities etc.)
 */
func (server *Server) Tick(currentTick int) {
	server.GetScheduler().DoTick()
}
