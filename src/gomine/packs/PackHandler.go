package packs

import (
	"gomine/interfaces"
	"io/ioutil"
	"path/filepath"
)

type PackHandler struct {
	server interfaces.IServer

	resourcePacks map[string]*ResourcePack
	behaviorPacks map[string]*BehaviorPack
}

/**
 * Returns a new pack handler.
 */
func NewPackHandler(server interfaces.IServer) *PackHandler {
	return &PackHandler{server, make(map[string]*ResourcePack), make(map[string]*BehaviorPack)}
}

/**
 * Returns all loaded resource packs.
 */
func (handler *PackHandler) GetResourcePacks() map[string]*ResourcePack {
	return handler.resourcePacks
}

/**
 * Returns all loaded behavior packs.
 */
func (handler *PackHandler) GetBehaviorPacks() map[string]*BehaviorPack {
	return handler.behaviorPacks
}

/**
 * Loads all resource packs in the extensions/resource_packs/ folder.
 */
func (handler *PackHandler) LoadResourcePacks() {
	var path = handler.server.GetServerPath() + "extensions/resource_packs/"
	var files, _ = ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		extension := filepath.Ext(file.Name())
		if extension != "mcpack" && extension != "zip" {
			continue
		}

		fileName := file.Name()

		resourcePack := NewResourcePack(fileName)
		resourcePack.Load()
		err := resourcePack.Validate()
		if err != nil {
			handler.server.GetLogger().LogError(err)
			continue
		}
		handler.resourcePacks[resourcePack.manifest.Header.UUID] = resourcePack
	}
}

/**
 * Checks if a resource pack with the given UUID exists.
 */
func (handler *PackHandler) IsResourcePackLoaded(uuid string) bool {
	var _, exists = handler.resourcePacks[uuid]
	return exists
}

/**
 * Returns a resource pack with the given UUID, or nil if it doesn't exist.
 */
func (handler *PackHandler) GetResourcePack(uuid string) *ResourcePack {
	if !handler.IsResourcePackLoaded(uuid) {
		return nil
	}
	return handler.resourcePacks[uuid]
}

func (handler *PackHandler) LoadBehaviorPacks() {
	//var path = handler.server.GetServerPath() + "extensions/behavior_packs/"
}