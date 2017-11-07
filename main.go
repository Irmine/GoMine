package main

import (
	"fmt"
	"gomine"
	time2 "time"
)

var currentTick int = 0

func main() {

	var server = gomine.NewServer()
	server.Start()
	var tickDuration = 50 * int64(time2.Millisecond)

	for {
		var nextTime = time2.Now().Add(time2.Duration(tickDuration))

		server.Tick(currentTick)

		var diff = nextTime.Sub(time2.Now()).Nanoseconds()

		if diff > 0 {
			time2.Sleep(time2.Duration(diff))
		} else {
			// Tick time took longer than it should. Lower tick rate?
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
