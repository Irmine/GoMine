package main

import (
	. "github.com/irmine/gomine"
	"github.com/irmine/gomine/resources"
	. "github.com/irmine/gomine/text"
	. "os"
	. "path/filepath"
	"time"
)

func main() {
	startTime := time.Now()
	path, err := GetServerPath()
	if err != nil {
		panic(err)
	}
	SetUpDirectories(path)
	config := resources.NewGoMineConfig(path)
	server := NewServer(path, config)

	if err := server.Start(); err != nil {
		panic(err)
	}
	DefaultLogger.Info("Server startup done! Took:", time.Now().Sub(startTime))

	for range time.NewTicker(time.Second / 20).C {
		if !server.IsRunning() {
			break
		}
		server.Tick()
	}
}

// getServerPath returns the server path.
func GetServerPath() (string, error) {
	executable, err := Executable()
	return Dir(executable) + "/", err
}

// setUpDirectories sets up all directories needed for GoMine.
func SetUpDirectories(path string) {
	Mkdir(path+"extensions", 0700)
	Mkdir(path+"extensions/plugins", 0700)
	Mkdir(path+"extensions/behavior_packs", 0700)
	Mkdir(path+"extensions/resource_packs", 0700)
}
