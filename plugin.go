package gomine

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
	GetServer() *Server
	OnEnable()

	GetName() string
	GetVersion() string
	GetAuthor() string
	GetOrganisation() string
	GetAPIVersion() string
	setManifest(IManifest)
}

type Plugin struct {
	server *Server

	manifest IManifest
}

func NewPlugin(server *Server) *Plugin {
	return &Plugin{server, Manifest{}}
}

// GetName returns the name of the manifest.
func (manifest Manifest) GetName() string {
	return manifest.Name
}

// GetVersion returns the version of the manifest.
func (manifest Manifest) GetVersion() string {
	return manifest.Version
}

// GetOrganisation returns the author of the manifest.
func (manifest Manifest) GetOrganisation() string {
	return manifest.Organisation
}

// GetAPIVersion returns the API Version of the manifest.
func (manifest Manifest) GetAPIVersion() string {
	return manifest.APIVersion
}

// GetAuthor returns the author of the manifest.
func (manifest Manifest) GetAuthor() string {
	return manifest.Author
}

// GetDescription returns the description of the manifest.
func (manifest Manifest) GetDescription() string {
	return manifest.Description
}

// GetName returns the name of the plugin.
func (plug *Plugin) GetName() string {
	return plug.manifest.GetName()
}

// GetVersion returns the version of the plugin.
func (plug *Plugin) GetVersion() string {
	return plug.manifest.GetVersion()
}

// GetAuthor returns the author of the plugin.
func (plug *Plugin) GetOrganisation() string {
	return plug.manifest.GetOrganisation()
}

// GetAPIVersion returns the API Version of the plugin.
func (plug *Plugin) GetAPIVersion() string {
	return plug.manifest.GetAPIVersion()
}

// GetAuthor returns the author of the plugin.
func (plug *Plugin) GetAuthor() string {
	return plug.manifest.GetAuthor()
}

// GetDescription returns the description of the plugin.
func (plug *Plugin) GetDescription() string {
	return plug.manifest.GetDescription()
}

// SetManifest sets the manifest of this plugin.
func (plug *Plugin) setManifest(manifest IManifest) {
	plug.manifest = manifest
}

// GetServer returns the main server.
func (plug *Plugin) GetServer() *Server {
	return plug.server
}
