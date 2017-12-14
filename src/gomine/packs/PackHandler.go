package packs

import (
	"gomine/interfaces"
	"io/ioutil"
	"path/filepath"
)

type PackHandler struct {
	server interfaces.IServer

	resourcePacks map[string]interfaces.IPack
	selectedResourcePack *ResourcePack

	behaviorPacks map[string]interfaces.IPack
	selectedBehaviorPack *BehaviorPack
}

/**
 * Returns a new pack handler.
 */
func NewPackHandler(server interfaces.IServer) *PackHandler {
	return &PackHandler{server, make(map[string]interfaces.IPack), nil, make(map[string]interfaces.IPack), nil}
}

/**
 * Returns all loaded resource packs.
 */
func (handler *PackHandler) GetResourcePacks() map[string]interfaces.IPack {
	return handler.resourcePacks
}

/**
 * Returns all loaded behavior packs.
 */
func (handler *PackHandler) GetBehaviorPacks() map[string]interfaces.IPack {
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
		if extension != ".mcpack" && extension != ".zip" {
			continue
		}

		filePath := path + file.Name()

		resourcePack := NewResourcePack(filePath)
		err := resourcePack.Load()
		if err != nil {
			handler.server.GetLogger().LogError(err)
			continue
		}

		err = resourcePack.ValidateManifest()
		if err != nil {
			handler.server.GetLogger().LogError(err)
			continue
		}

		handler.resourcePacks[resourcePack.manifest.Header.UUID] = resourcePack
		handler.server.GetLogger().Debug("Loaded resource pack: " + resourcePack.manifest.Header.Name)

		if file.Name() == handler.server.GetConfiguration().SelectedResourcePack {
			handler.server.GetLogger().Info("Selected resource pack: " + resourcePack.manifest.Header.Name)
			handler.selectedResourcePack = resourcePack
		}
	}
}

/**
 * Loads all behavior packs in the extensions/behavior_packs/ folder.
 */
func (handler *PackHandler) LoadBehaviorPacks() {
	var path = handler.server.GetServerPath() + "extensions/behavior_packs/"
	var files, _ = ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		extension := filepath.Ext(file.Name())
		if extension != ".mcpack" && extension != ".zip" {
			continue
		}

		filePath := path + file.Name()

		behaviorPack := NewBehaviorPack(filePath)
		err := behaviorPack.Load()
		if err != nil {
			handler.server.GetLogger().LogError(err)
			continue
		}

		err = behaviorPack.ValidateManifest()
		if err != nil {
			handler.server.GetLogger().LogError(err)
			continue
		}

		handler.resourcePacks[behaviorPack.manifest.Header.UUID] = behaviorPack
		handler.server.GetLogger().Debug("Loaded behavior pack: " + behaviorPack.manifest.Header.Name)

		if file.Name() == handler.server.GetConfiguration().SelectedResourcePack {
			handler.server.GetLogger().Info("Selected behavior pack: " + behaviorPack.manifest.Header.Name)
			handler.selectedBehaviorPack = behaviorPack
		}
	}
}

/**
 * Returns the selected resource pack, or nil if none is available.
 */
func (handler *PackHandler) GetSelectedResourcePack() interfaces.IPack {
	return handler.selectedResourcePack
}

/**
 * Returns the selected behavior pack, or nil if none is available.
 */
func (handler *PackHandler) GetSelectedBehaviorPack() interfaces.IPack {
	return handler.selectedBehaviorPack
}

/**
 * Checks if a resource pack with the given UUID exists.
 */
func (handler *PackHandler) IsResourcePackLoaded(uuid string) bool {
	var _, exists = handler.resourcePacks[uuid]
	return exists
}

/**
 * Checks if a behavior pack with the given UUID exists.
 */
func (handler *PackHandler) IsBehaviorPackLoaded(uuid string) bool {
	var _, exists = handler.behaviorPacks[uuid]
	return exists
}

/**
 * Returns if there's a resource OR behavior pack with the given UUID loaded.
 */
func (handler *PackHandler) IsPackLoaded(uuid string) bool {
	return handler.IsResourcePackLoaded(uuid) || handler.IsBehaviorPackLoaded(uuid)
}

/**
 * Returns a resource pack with the given UUID, or nil if it doesn't exist.
 */
func (handler *PackHandler) GetResourcePack(uuid string) interfaces.IPack {
	if !handler.IsResourcePackLoaded(uuid) {
		return nil
	}
	return handler.resourcePacks[uuid]
}

/**
 * Returns a behavior pack with the given UUID, or nil if it doesn't exist.
 */
func (handler *PackHandler) GetBehaviorPack(uuid string) interfaces.IPack {
	if !handler.IsBehaviorPackLoaded(uuid) {
		return nil
	}
	return handler.behaviorPacks[uuid]
}

/**
 * Returns either a behavior or resource pack with the given UUID, or nil if none exist.
 */
func (handler *PackHandler) GetPack(uuid string) interfaces.IPack {
	if handler.GetResourcePack(uuid) != nil {
		return handler.GetResourcePack(uuid)
	}
	return handler.GetBehaviorPack(uuid)
}

/**
 * Returns all behavior packs in a slice.
 */
func (handler *PackHandler) GetBehaviorPackSlice() []interfaces.IPack {
	var packs = handler.behaviorPacks
	var packsSlice []interfaces.IPack

	for _, pack := range packs {
		packsSlice = append(packsSlice, pack)
	}

	return packsSlice
}

/**
 * Returns all resource packs in a slice.
 */
func (handler *PackHandler) GetResourcePackSlice() []interfaces.IPack {
	var packs = handler.resourcePacks
	var packsSlice []interfaces.IPack

	for _, pack := range packs {
		packsSlice = append(packsSlice, pack)
	}

	return packsSlice
}