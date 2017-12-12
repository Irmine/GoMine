package main

import (
	"fmt"
	"gomine"
	"time"
	"runtime"
	"os"
	"path/filepath"
	"flag"
)

/**
 * Command line flags:
 * -stop-immediately : Stops the server immediately after starting and ticking once.
 */

var stopInstantly = false

var ticker = time.NewTicker(time.Second / 20)
var currentTick = 0

func main() {
	var startTime = time.Now()
	if !checkRequirements() {
		return
	}

	parseFlags()

	var serverPath = scanServerPath()
	setUpDirectories(serverPath)

	var server = gomine.NewServer(serverPath)

	server.Start()
	var startupTime = time.Now().Sub(startTime)

	server.GetLogger().Info("Server startup done! Took: " + startupTime.String())

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

	server.GetLogger().Terminate() // Terminate the logger to stop writing asynchronously.
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
 * Checks if the Go installation meets the requirements of GoMine.
 */
func checkRequirements() bool {
	var version = runtime.Version()
	if version != "go1.9.2" {
		fmt.Println("Please install the Go 1.9.2 release.")
		return false
	}

	return true
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
	os.Mkdir(path + "extensions", os.ModeDir)
	os.Mkdir(path + "extensions/plugins", os.ModeDir)
	os.Mkdir(path + "extensions/behavior_packs", os.ModeDir)
	os.Mkdir(path + "extensions/resource_packs", os.ModeDir)
}