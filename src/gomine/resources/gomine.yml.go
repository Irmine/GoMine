package resources

import (
	"io/ioutil"
	"libraries/yaml"
	"os"
)

type GoMineConfig struct {
	Version string `yaml:"version"`
	DebugMode bool `yaml:"debug-mode"`
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
	var _, error = os.Stat(path)

	if os.IsNotExist(error) {
		var data, _ = yaml.Marshal(GoMineConfig{
			Version: "0.0.1",
			DebugMode: true,
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
