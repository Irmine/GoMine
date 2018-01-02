package plugins

import (
	"gomine/interfaces"
	"io/ioutil"
	"path/filepath"
	"plugin"
	"errors"
	"strings"
	"gomine"
)

type PluginManager struct {
	server interfaces.IServer
	plugins map[string]IPlugin
}

func NewPluginManager(server interfaces.IServer) *PluginManager {
	return &PluginManager{server, make(map[string]IPlugin)}
}

/**
 * Returns the main server.
 */
func (manager *PluginManager) GetServer() interfaces.IServer {
	return manager.server
}

/**
 * Returns a plugin with the given name, or nil if none could be found.
 */
func (manager *PluginManager) GetPlugin(name string) IPlugin {
	if !manager.IsPluginLoaded(name) {
		return nil
	}
	return manager.plugins[name]
}

/**
 * Checks if a plugin with the given name is loaded.
 */
func (manager *PluginManager) IsPluginLoaded(name string) bool {
	var _, exists = manager.plugins[name]
	return exists
}

/**
 * Loads all plugins in the 'extensions/plugins' folder.
 */
func (manager *PluginManager) LoadPlugins() {
	var path = manager.server.GetServerPath() + "extensions/plugins/"
	var files, _ = ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := path + file.Name()
		extension := filepath.Ext(filePath)

		if extension != ".so" {
			continue
		}

		err := manager.LoadPlugin(filePath)
		manager.server.GetLogger().LogError(err)
	}
}

/**
 * Loads a plugin at the given file path and returns an error if applicable.
 */
func (manager *PluginManager) LoadPlugin(filePath string) error {
	plug, err := plugin.Open(filePath)
	if err != nil {
		return err
	}

	manifestSymbol, err := plug.Lookup("Manifest")
	if err != nil {
		return errors.New("Plugin at '" + filePath + "' does not have a Manifest.")
	}

	manifest, err := manifestSymbol.(Manifest)
	if err != nil {
		return errors.New("Plugin at '" + filePath + "' does not have a valid Manifest.")
	}

	err = manager.ValidateManifest(manifest, filePath)
	if err != nil {
		return err
	}

	newPluginSymbol, err := plug.Lookup("NewPlugin")
	if err != nil {
		return errors.New("Plugin at '" + filePath + "' does not have a NewPlugin function.")
	}

	pluginFunc, err := newPluginSymbol.(func(server interfaces.IServer) IPlugin)
	if err != nil {
		return errors.New("Plugin at '" + filePath + "' does not have a valid NewPlugin function.")
	}
	var finalPlugin = pluginFunc(manager.server)
	finalPlugin.setManifest(manifest)

	manager.plugins[finalPlugin.GetName()] = finalPlugin

	return nil
}

/**
 * Validates the plugin manifest.
 */
func (manager *PluginManager) ValidateManifest(manifest Manifest, path string) error {
	if manifest.Name == "" {
		return errors.New("Plugin manifest at " + path + " is missing a name.")
	}
	if manifest.Description == "" {
		return errors.New("Plugin manifest at " + path + " is missing a description.")
	}
	var dotCount = strings.Count(manifest.Version, ".")
	if dotCount < 1 {
		return errors.New("Plugin manifest at " + path + " is missing a valid version.")
	}

	var digits = strings.Split(manifest.APIVersion, ".")
	if len(digits) < 2 {
		return errors.New("Plugin manifest at " + path + " is missing a valid API version.")
	}
	var currentDigits = strings.Split(gomine.ApiVersion, ".")

	if digits[0] != currentDigits[0] {
		return errors.New("Plugin manifest at " + path + " has an incompatible greater API version. Got: " + digits[0] + ".~, Expected: " + currentDigits[0] + ".~")
	}

	return nil
}
