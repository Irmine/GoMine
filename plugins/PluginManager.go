package plugins

import (
	"gomine/interfaces"
	"io/ioutil"
	"path/filepath"
	"plugin"
	"errors"
	"strings"
	"os/exec"
	"os"
	"gomine/utils"
)

const (
	ApiVersion = "0.0.1"

	OutdatedPlugin = "plugin.Open: plugin was built with a different version of package"
	NoPluginsSupported = "plugin: not implemented"
)

type PluginManager struct {
	server interfaces.IServer
	plugins map[string]IPlugin
}

func NewPluginManager(server interfaces.IServer) *PluginManager {
	return &PluginManager{server, make(map[string]IPlugin)}
}

/**
 * Returns all plugins currently loaded on the server.
 */
func (manager *PluginManager) GetPlugins() map[string]IPlugin {
	return manager.plugins
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
		if err != nil {
			if err.Error() == NoPluginsSupported {
				manager.server.GetLogger().Error("Go does currently not support plugins for your operating system.")
				return
			}
		}
		manager.server.GetLogger().LogError(err)
	}
}

/**
 * Compiles a plugin.go at the given path during runtime, and opens it. This action is extremely time consuming.
 */
func (manager *PluginManager) CompilePlugin(filePath string) (*plugin.Plugin, error) {
	var compiledPath = strings.Replace(strings.Replace(filePath, ".go", "", 1), "\\", "/", -1)
	compiledPath += "~" + utils.GenerateRandomUUID() + ".so"

	var cmd = exec.Command("go", "build", "-buildmode=plugin", "-i", "-o", compiledPath, filePath)
	var output, err = cmd.CombinedOutput()

	if err != nil {
		manager.server.GetLogger().LogError(err)
		manager.server.GetLogger().Error(string(output))
	}

	plug, err := plugin.Open(compiledPath)

	return plug, err
}

/**
 * Recompiles a plugin.so at the given path, provided the main source file is at the same location suffixed with .go.
 */
func (manager *PluginManager) RecompilePlugin(filePath string) (*plugin.Plugin, error) {
	var decompiledPath = strings.Replace(strings.Replace(filePath, ".so", ".go", 1), "\\", "/", -1)
	if strings.Contains(filePath, "~") {
		decompiledPath = strings.Split(decompiledPath, "~")[0] + ".go"
	}

	os.Remove(filePath)

	return manager.CompilePlugin(decompiledPath)
}

/**
 * Loads a plugin at the given file path and returns an error if applicable.
 */
func (manager *PluginManager) LoadPlugin(filePath string) error {
	var plug, err = plugin.Open(filePath)

	if err != nil {
		if strings.Contains(err.Error(), OutdatedPlugin) {
			manager.server.GetLogger().Notice("Outdated plugin. Recompiling plugin... This might take a bit.")
			var newPlugin, newErr = manager.RecompilePlugin(filePath)
			if newErr != nil {
				return newErr
			}
			plug = newPlugin
		} else {
			return err
		}
	}

	manifestSymbol, err := plug.Lookup("Manifest")
	if err != nil {
		return errors.New("Plugin at '" + filePath + "' does not have a Manifest.")
	}

	manifest, ok := manifestSymbol.(IManifest)
	if !ok {
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

	pluginFunc, ok := newPluginSymbol.(func(server interfaces.IServer) IPlugin)
	if !ok {
		return errors.New("Plugin at '" + filePath + "' does not have a valid NewPlugin function.")
	}

	var finalPlugin = pluginFunc(manager.server)
	finalPlugin.setManifest(manifest)

	manager.plugins[finalPlugin.GetName()] = finalPlugin
	finalPlugin.OnEnable()

	return nil
}

/**
 * Validates the plugin manifest and checks for duplicated plugins.
 */
func (manager *PluginManager) ValidateManifest(manifest IManifest, path string) error {
	if manifest.GetName() == "" {
		return errors.New("Plugin manifest at " + path + " is missing a name.")
	}
	if manager.IsPluginLoaded(manifest.GetName()) {
		return errors.New("Found duplicated plugin at " + path)
	}

	if manifest.GetDescription() == "" {
		return errors.New("Plugin manifest at " + path + " is missing a description.")
	}

	var dotCount = strings.Count(manifest.GetVersion(), ".")
	if dotCount < 1 {
		return errors.New("Plugin manifest at " + path + " is missing a valid version.")
	}

	var digits = strings.Split(manifest.GetAPIVersion(), ".")
	if len(digits) < 2 {
		return errors.New("Plugin manifest at " + path + " is missing a valid API version.")
	}
	var currentDigits = strings.Split(ApiVersion, ".")

	if digits[0] != currentDigits[0] {
		return errors.New("Plugin manifest at " + path + " has an incompatible greater API version. Got: " + digits[0] + ".~, Expected: " + currentDigits[0] + ".~")
	}

	return nil
}
