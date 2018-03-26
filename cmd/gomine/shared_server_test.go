package main

import (
	"github.com/irmine/gomine"
	"github.com/irmine/gomine/resources"
	"github.com/irmine/gomine/text"
	"testing"
	"time"
)

func TestSharedServer(t *testing.T) {
	ports := []uint16{19132, 19133, 19134, 19135, 19136}
	for _, port := range ports {
		go StartServer(port)
	}
	time.Sleep(time.Minute * 10)
}

func StartServer(port uint16) {
	text.DefaultLogger.Info("Starting server with port:", port)
	startTime := time.Now()
	path, err := GetServerPath()
	if err != nil {
		panic(err)
	}
	SetUpDirectories(path)
	config := resources.NewGoMineConfig(path)
	config.ServerPort = port
	server := gomine.NewServer(path, config)

	if err := server.Start(); err != nil {
		panic(err)
	}
	text.DefaultLogger.Info("Server startup done! Took:", time.Now().Sub(startTime))

	for range time.NewTicker(time.Second / 20).C {
		if !server.IsRunning() {
			break
		}
		server.Tick()
	}
}
