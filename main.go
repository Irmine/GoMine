package main

import (
	"fmt"
	"gomine"
	time2 "time"
	"runtime"
)

var currentTick int = 0

func main() {
	if !checkRequirements() {
		return
	}

	var server, error = gomine.NewServer()
	if error != nil {
		fmt.Println("Another instance of the server is already running.")
		return
	}
	server.Start()

	var tickDrop = 20

	for {
		var tickDuration = int(1.0 / float32(server.GetTickRate()) * 1000) * int(time2.Millisecond)
		var nextTime = time2.Now().Add(time2.Duration(tickDuration))

		server.Tick(currentTick)

		var diff = nextTime.Sub(time2.Now()).Nanoseconds()

		if diff > 0 {
			tickDrop--

			if tickDrop < 0 && server.GetTickRate() != 20 && diff > 5 * int64(time2.Millisecond) {
				server.SetTickRate(server.GetTickRate() + 1)
				fmt.Println("Elevating tick rate to:", server.GetTickRate())
			}

			time2.Sleep(time2.Duration(diff))
		} else {
			tickDrop++

			if tickDrop > 40 {
				server.SetTickRate(server.GetTickRate() - 1)
				fmt.Println("Lowering tick rate to:", server.GetTickRate())
			}
		}

		if !server.IsRunning() {
			server.Shutdown()
			fmt.Println("Server shut down.")
			break
		}

		currentTick++
	}

	// Other shutdown code.
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
