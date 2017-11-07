package gomine

import (
	"errors"
	"gomine/tasks"
)

type server struct {
	isRunning bool
	tickRate int
	scheduler tasks.Scheduler
}

var started bool = false

/**
 * Creates a new server.
 * Will report an error if a server is already existent.
 */
func NewServer() (server, error) {
	var errorServer server
	if started {
		return errorServer, errors.New("cannot create a second server")
	}

	var server = server{}
	server.tickRate = 20
	server.scheduler = tasks.NewScheduler()

	return server, nil
}

/**
 * Returns whether the server is running or not.
 */
func (server *server) IsRunning() bool {
	return server.isRunning
}

/**
 * Starts the server.
 */
func (server *server) Start() {
	server.isRunning = true
}

/**
 * Shuts down the server.
 */
func (server *server) Shutdown() {
	server.isRunning = false
}

/**
 * Returns the tick rate of the server.
 */
func (server *server) GetTickRate() int {
	return server.tickRate
}

/**
 * Resets the tick value back to the default. (20)
 */
func (server *server) ResetTickRate() {
	server.tickRate = 20
}

/**
 * Internal. Not to be used by plugins.
 */
func (server *server) SetTickRate(tickRate int) {
	server.tickRate = tickRate
}

/**
 * Returns the scheduler used for scheduling tasks.
 */
func (server *server) GetScheduler() tasks.Scheduler {
	return server.scheduler
}

/**
 * Internal. Not to be used by plugins.
 * Ticks the entire server. (Entities, block entities etc.)
 */
func (server *server) Tick(currentTick int) {
	server.GetScheduler().DoTick()
}