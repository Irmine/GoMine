package serverPkg

type server struct {
	isRunning bool
}

func NewServer() server {
	return server{}
}

func (server *server) IsRunning() bool {
	return server.isRunning
}

func (server *server) Start() {
	server.isRunning = true
}

func (server *server) Shutdown() {
	server.isRunning = false
}

/*
 * Internal. Do not use manually.
 */
func (server *server) Tick(currentTick int) {

}