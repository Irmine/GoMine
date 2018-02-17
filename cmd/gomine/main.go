package main

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/irmine/gomine"
)

/*
 * Command line flags:
 * -stop-immediately : Stops the server immediately after starting and ticking once.
 */

var stopInstantly = false

var ticker = time.NewTicker(time.Second / 20)
var currentTick int64 = 0

func main() {
	var startTime = time.Now()

	parseFlags()

	var serverPath = scanServerPath()
	setUpDirectories(serverPath)

	var server = gomine.NewServer(serverPath)

	server.Start()
	var startupTime = time.Now().Sub(startTime)

	server.GetLogger().Info("Server startup done! Took:", startupTime.String())

	for range ticker.C {
		server.Tick(currentTick)

		if stopInstantly {
			server.Shutdown()
		}
		if !server.IsRunning() {
			break
		}
		currentTick++
	}

	server.GetLogger().Terminate()        // Terminate the logger to stop writing asynchronously.
	server.GetLogger().ProcessQueue(true) // Process the logger queue one last time forced and synchronously to make sure everything gets written.
	server.GetLogger().Sync()
}

/**
 * Scans and returns the server path.
 */
func scanServerPath() string {
	var executable, err = os.Executable()
	if err != nil {
		panic(err)
	}
	var serverPath = filepath.Dir(executable) + "/"

	return serverPath
}

/**
 * Parses all command line flags.
 */
func parseFlags() {
	var instantStop = flag.Bool("stop-immediately", false, "instant stop")

	flag.Parse()

	stopInstantly = *instantStop
}

/**
 * Sets up all directories needed for GoMine.
 */
func setUpDirectories(path string) {
	os.Mkdir(path+"extensions", os.ModeDir)
	os.Mkdir(path+"extensions/plugins", os.ModeDir)
	os.Mkdir(path+"extensions/behavior_packs", os.ModeDir)
	os.Mkdir(path+"extensions/resource_packs", os.ModeDir)
}
