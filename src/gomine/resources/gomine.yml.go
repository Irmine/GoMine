package resources

import (
	"io/ioutil"
	"libraries/yaml"
	"os"
)

type GoMineConfig struct {
	ServerName string `yaml:"Server LAN Name"`
	ServerMotd string `yaml:"Server MOTD"`
	ServerIp string `yaml:"Server IP"`
	ServerPort uint16 `yaml:"Server Port"`

	MaximumPlayers uint `yaml:"Maximum Players"`
	DefaultGameMode byte `yaml:"Default Gamemode"`

	DebugMode bool `yaml:"Debug Mode"`

	DefaultLevel string `yaml:"Default Level"`
	DefaultGenerator string `yaml:"Default Generator"`

	ForceResourcePacks bool `yaml:"Forced Resource Packs"`
	SelectedResourcePack string `yaml:"Selected Resource Pack"`

	XBOXLiveAuth bool `yaml:"XBOX Live Auth"`
	UseEncryption bool `yaml:"Use Encryption"`
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
			DefaultGameMode: 1,

			DebugMode: true,

			DefaultLevel: "world",
			DefaultGenerator: "Flat",

			ForceResourcePacks: false,
			SelectedResourcePack: "",

			XBOXLiveAuth: true,
			UseEncryption: false,
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