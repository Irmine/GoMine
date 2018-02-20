package main

import (
	"time"
	"os"
	"path/filepath"
	"flag"
	"github.com/irmine/gomine"
)

var ticker = time.NewTicker(time.Second / 20)
var currentTick int64 = 0

func main() {
	var startTime = time.Now()

	parseFlags()

	var serverPath = getServerPath()
	setUpDirectories(serverPath)

	var server = gomine.NewServer(serverPath)

	server.Start()
	var startupTime = time.Now().Sub(startTime)

	server.GetLogger().Info("Server startup done! Took:", startupTime.String())

	for range ticker.C {
		server.Tick(currentTick)

		if !server.IsRunning() {
			break
		}
		currentTick++
	}

	server.GetLogger().Terminate()        // Terminate the logger to stop writing asynchronously.
	server.GetLogger().ProcessQueue(true) // Process the logger queue one last time forced and synchronously to make sure everything gets written.
	server.GetLogger().Sync()
}

// getServerPath returns the server path.
func getServerPath() string {
	var executable, err = os.Executable()
	if err != nil {
		panic(err)
	}
	var serverPath = filepath.Dir(executable) + "/"

	return serverPath
}

// parseFlags parses all command line flags.
func parseFlags() {
	flag.Parse()
}

// setUpDirectories sets up all directories needed for GoMine.
func setUpDirectories(path string) {
	os.Mkdir(path+"extensions", 0777)
	os.Mkdir(path+"extensions/plugins", 0777)
	os.Mkdir(path+"extensions/behavior_packs", 0777)
	os.Mkdir(path+"extensions/resource_packs", 0777)
}
