package plugins

import "github.com/irmine/gomine/interfaces"

type Manifest struct {
	Name         string
	Description  string
	Version      string
	APIVersion   string
	Author       string
	Organisation string
}

type IManifest interface {
	GetName() string
	GetDescription() string
	GetVersion() string
	GetAPIVersion() string
	GetAuthor() string
	GetOrganisation() string
}

type IPlugin interface {
	GetServer() interfaces.IServer
	OnEnable()

	GetName() string
	GetVersion() string
	GetAuthor() string
	GetOrganisation() string
	GetAPIVersion() string
	setManifest(IManifest)
}

type Plugin struct {
	server interfaces.IServer

	manifest IManifest
}

func NewPlugin(server interfaces.IServer) *Plugin {
	return &Plugin{server, Manifest{}}
}

/**
 * Returns the name of the manifest.
 */
func (manifest Manifest) GetName() string {
	return manifest.Name
}

/**
 * Returns the version of the manifest.
 */
func (manifest Manifest) GetVersion() string {
	return manifest.Version
}

/**
 * Returns the author of the manifest.
 */
func (manifest Manifest) GetOrganisation() string {
	return manifest.Organisation
}

/**
 * Returns the API Version of the manifest.
 */
func (manifest Manifest) GetAPIVersion() string {
	return manifest.APIVersion
}

/**
 * Returns the author of the manifest.
 */
func (manifest Manifest) GetAuthor() string {
	return manifest.Author
}

/**
 * Returns the description of the manifest.
 */
func (manifest Manifest) GetDescription() string {
	return manifest.Description
}

/**
 * Returns the name of the plugin.
 */
func (plug *Plugin) GetName() string {
	return plug.manifest.GetName()
}

/**
 * Returns the version of the plugin.
 */
func (plug *Plugin) GetVersion() string {
	return plug.manifest.GetVersion()
}

/**
 * Returns the author of the plugin.
 */
func (plug *Plugin) GetOrganisation() string {
	return plug.manifest.GetOrganisation()
}

/**
 * Returns the API Version of the plugin.
 */
func (plug *Plugin) GetAPIVersion() string {
	return plug.manifest.GetAPIVersion()
}

/**
 * Returns the author of the plugin.
 */
func (plug *Plugin) GetAuthor() string {
	return plug.manifest.GetAuthor()
}

/**
 * Returns the description of the plugin.
 */
func (plug *Plugin) GetDescription() string {
	return plug.manifest.GetDescription()
}

/**
 * Sets the manifest of this plugin.
 */
func (plug *Plugin) setManifest(manifest IManifest) {
	plug.manifest = manifest
}

/**
 * Returns the main server.
 */
func (plug *Plugin) GetServer() interfaces.IServer {
	return plug.server
}
