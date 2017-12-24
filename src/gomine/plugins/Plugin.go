package plugins

import "gomine/interfaces"

type Manifest struct {
	Name string
	Description string
	Version string
	APIVersion string
	Author string
	Organisation string
}

type IPlugin interface {
	GetServer() interfaces.IServer
	OnEnable()

	GetName() string
	GetVersion() string
	GetAuthor() string
	GetOrganisation() string
	GetAPIVersion() string
	setManifest(Manifest)
}

type Plugin struct {
	server interfaces.IServer

	manifest Manifest
}

func NewPlugin(server interfaces.IServer) *Plugin {
	return &Plugin{server, nil}
}

/**
 * Returns the name of the plugin.
 */
func (plug *Plugin) GetName() string {
	return plug.manifest.Name
}

/**
 * Returns the version of the plugin.
 */
func (plug *Plugin) GetVersion() string {
	return plug.manifest.Version
}

/**
 * Returns the author of the plugin.
 */
func (plug *Plugin) GetOrganisation() string {
	return plug.manifest.Organisation
}

/**
 * Returns the API Version of the plugin.
 */
func (plug *Plugin) GetAPIVersion() string {
	return plug.manifest.APIVersion
}

/**
 * Returns the author of the plugin.
 */
func (plug *Plugin) GetAuthor() string {
	return plug.manifest.Author
}

/**
 * Returns the description of the plugin.
 */
func (plug *Plugin) GetDescription() string {
	return plug.manifest.Description
}

/**
 * Sets the manifest of this plugin.
 */
func (plug *Plugin) setManifest(manifest Manifest) {
	plug.manifest = manifest
}

/**
 * Returns the main server.
 */
func (plug *Plugin) GetServer() interfaces.IServer {
	return plug.server
}