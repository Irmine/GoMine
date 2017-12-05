package resources

import (
	"io/ioutil"
	"libraries/yaml"
	"os"
)

type GoMineConfig struct {
	ServerName string `yaml:"server-name"`
	ServerMotd string `yaml:"server-motd"`
	ServerIp string `yaml:"server-ip"`
	ServerPort uint16 `yaml:"server-port"`

	MaximumPlayers uint `yaml:"max-players"`
	DefaultGameMode byte `yaml:"default-gamemode"`

	DebugMode bool `yaml:"debug-mode"`

	DefaultLevel string `yaml:"default-level"`
	DefaultGenerator string `yaml:"default-generator"`
}

/**
 * Returns a new configuration struct.
 * Creates the file if it does not yet exist.
 */
func NewGoMineConfig(serverPath string) *GoMineConfig {
	initializeConfig(serverPath)
	return getGoMineConfig(serverPath)
}

/**
 * Initializes the configuration file if it does not yet exist.
 */
func initializeConfig(serverPath string) {
	var path = serverPath + "gomine.yml"
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var data, _ = yaml.Marshal(GoMineConfig{
			ServerName: "GoMine Server",
			ServerMotd: "GoMine Testing Server",
			ServerIp: "0.0.0.0",
			ServerPort: 19132,

			MaximumPlayers: 20,

			DebugMode: true,

			DefaultLevel: "world",
			DefaultGenerator: "Flat",
		})
		var file, _ = os.OpenFile(path, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		file.WriteString(string(data))
		file.Sync()
	}
}

/**
 * Parses the configuration file into a struct.
 */
func getGoMineConfig(serverPath string) *GoMineConfig {
	var yamlFile, _ = ioutil.ReadFile(serverPath + "gomine.yml")

	var config = &GoMineConfig{}
	yaml.Unmarshal(yamlFile, config)

	return config
}
