package main

import (
	. "github.com/irmine/gomine"
	. "github.com/irmine/gomine/text"
	. "os"
	. "path/filepath"
	"time"
)

func main() {
	startTime := time.Now()
	path, err := getServerPath()
	if err != nil {
		panic(err)
	}
	setUpDirectories(path)
	server := NewServer(path)

	server.Start()
	DefaultLogger.Info("Server startup done! Took:", time.Now().Sub(startTime))

	for range time.NewTicker(time.Second / 20).C {
		if !server.IsRunning() {
			break
		}
		server.Tick()
	}
}

// getServerPath returns the server path.
func getServerPath() (string, error) {
	executable, err := Executable()
	return Dir(executable) + "/", err
}

// setUpDirectories sets up all directories needed for GoMine.
func setUpDirectories(path string) {
	Mkdir(path+"extensions", 0700)
	Mkdir(path+"extensions/plugins", 0700)
	Mkdir(path+"extensions/behavior_packs", 0700)
	Mkdir(path+"extensions/resource_packs", 0700)
}
