package packs

import (
	"gomine/interfaces"
	"io/ioutil"
	"path/filepath"
)

type PackHandler struct {
	server interfaces.IServer

	resourcePacks map[string]interfaces.IPack
	resourceStack interfaces.IPackStack

	behaviorPacks map[string]interfaces.IPack
	behaviorStack interfaces.IPackStack
}

/**
 * Returns a new pack handler.
 */
func NewPackHandler(server interfaces.IServer) *PackHandler {
	return &PackHandler{server, make(map[string]interfaces.IPack), NewPackStack(), make(map[string]interfaces.IPack), NewPackStack()}
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
 * Returns the resource pack stack.
 */
func (handler *PackHandler) GetResourceStack() interfaces.IPackStack {
	return handler.resourceStack
}

/**
 * Returns the behavior pack stack.
 */
func (handler *PackHandler) GetBehaviorStack() interfaces.IPackStack {
	return handler.behaviorStack
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
			handler.GetResourceStack().AddPackOnTop(resourcePack)
		} else {
			handler.GetResourceStack().AddPackOnBottom(resourcePack)
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

		err = behaviorPack.ValidateDependencies(handler)
		if err != nil {
			handler.server.GetLogger().LogError(err)
			continue
		}

		handler.resourcePacks[behaviorPack.manifest.Header.UUID] = behaviorPack
		handler.server.GetLogger().Debug("Loaded behavior pack: " + behaviorPack.manifest.Header.Name)

		if file.Name() == handler.server.GetConfiguration().SelectedResourcePack {
			handler.server.GetLogger().Info("Selected behavior pack: " + behaviorPack.manifest.Header.Name)
			handler.GetBehaviorStack().AddPackOnTop(behaviorPack)
		} else {
			handler.GetBehaviorStack().AddPackOnBottom(behaviorPack)
		}
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