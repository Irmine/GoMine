package main

import (
	"github.com/irmine/gomine"
	"github.com/irmine/gomine/resources"
	"github.com/irmine/gomine/text"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	path, err := GetServerPath()
	must(err)
	SetUpDirectories(path)

	config := resources.NewGoMineConfig(path)
	server := gomine.NewServer(path, config)

	must(server.Start())
	text.DefaultLogger.Info("Server startup done! Took:", time.Now().Sub(startTime))

	for range time.NewTicker(time.Second / 20).C {
		if !server.IsRunning() {
			break
		}
		server.Tick()
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// GetServerPath returns the server path.
func GetServerPath() (string, error) {
	executable, err := os.Executable()
	return strings.Replace(filepath.Dir(executable)+"/", `\`, "/", -1), err
}

// SetUpDirectories sets up all directories needed for GoMine.
func SetUpDirectories(path string) {
	os.Mkdir(path+"extensions", 0700)
	os.Mkdir(path+"extensions/plugins", 0700)
	os.Mkdir(path+"extensions/behavior_packs", 0700)
	os.Mkdir(path+"extensions/resource_packs", 0700)
}
