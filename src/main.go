package main

import (
	"fmt"
	"gomine"
	"time"
	"runtime"
	"os"
	"path/filepath"
	"strconv"
	"flag"
)

/**
 * Command line flags:
 * -stop-immediately : Stops the server immediately after starting and ticking once.
 */

var stopInstantly = false

var currentTick = 0

func main() {
	var startTime = time.Now()
	if !checkRequirements() {
		return
	}
	parseFlags()
	var serverPath = scanServerPath()

	var server, err = gomine.NewServer(serverPath)
	if err != nil {
		server.GetLogger().Critical("Another instance of the server is already running.")
		return
	}

	server.Start()
	var startupTime = time.Now().Sub(startTime)
	server.GetLogger().Info("Server startup done! Took: " + startupTime.String())

	var tickDrop = 20

	for {
		var tickDuration = int(1.0 / float32(server.GetTickRate()) * 1000) * int(time.Millisecond)
		var nextTime = time.Now().Add(time.Duration(tickDuration))

		server.Tick(currentTick)

		var diff = nextTime.Sub(time.Now()).Nanoseconds()

		if diff > 0 {
			tickDrop--

			if tickDrop < 0 && server.GetTickRate() != 20 && diff > 5 * int64(time.Millisecond) {
				server.SetTickRate(server.GetTickRate() + 1)

				server.GetLogger().Debug("Elevating tick rate to: " + strconv.Itoa(server.GetTickRate()))
			}

			time.Sleep(time.Duration(diff))
		} else {
			tickDrop++

			if tickDrop > 40 {
				server.SetTickRate(server.GetTickRate() - 1)
				server.GetLogger().Debug("Lowering tick rate to: " + strconv.Itoa(server.GetTickRate()))
			}
		}
		currentTick++

		if stopInstantly {
			server.Shutdown()
		}
		if !server.IsRunning() {
			break
		}
	}

	server.GetLogger().ProcessQueue() // Process the queue one last time synchronously to make sure everything gets written.
}

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
		fmt.Println("Please install the GoLang 1.9.2 release.")
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