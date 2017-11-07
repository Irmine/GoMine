package gomine

import "errors"

type server struct {
	isRunning bool
	tickRate int
}

var started bool = false

func NewServer() (server, error) {
	var errorServer server
	if started {
		return errorServer, errors.New("cannot create a second server")
	}

	var server = server{}
	server.tickRate = 20
	started = true
	return server, nil
}

func (server *server) IsRunning() bool {
	return server.isRunning
}

/*
 *
 */
func (server *server) Start() {
	server.isRunning = true
}

/*
 * Shuts down the server.
 */
func (server *server) Shutdown() {
	server.isRunning = false
}

/*
 * Returns the tick rate of the server.
 */
func (server *server) GetTickRate() int {
	return server.tickRate
}

/*
 * Resets the tick value back to the default. (20)
 */
func (server *server) ResetTickRate() {
	server.tickRate = 20
}

/*
 * Internal. Not to be used by plugins.
 */
func (server *server) SetTickRate(tickRate int) {
	server.tickRate = tickRate
}

/*
 * Internal. Not to be used by plugins.
 */
func (server *server) Tick(currentTick int) {

}