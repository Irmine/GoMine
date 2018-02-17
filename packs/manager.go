package packs

import (
	"io/ioutil"
	"path/filepath"
)

// Manager manages the loading of packs.
// It provides helper functions for both types of packs.
type Manager struct {
	serverPath string

	resourcePacks map[string]*ResourcePack
	resourceStack *Stack

	behaviorPacks map[string]*BehaviorPack
	behaviorStack *Stack
}

// NewManager returns a new pack manager with the given path.
func NewManager(serverPath string) *Manager {
	return &Manager{serverPath, make(map[string]*ResourcePack), NewStack(), make(map[string]*BehaviorPack), NewStack()}
}

// GetResourcePacks returns all resource maps in a UUID => pack map.
func (manager *Manager) GetResourcePacks() map[string]*ResourcePack {
	return manager.resourcePacks
}

// GetBehaviorPacks returns all behavior packs in a UUID => pack map.
func (manager *Manager) GetBehaviorPacks() map[string]*BehaviorPack {
	return manager.behaviorPacks
}

// GetResourceStack returns the resource pack stack.
func (manager *Manager) GetResourceStack() *Stack {
	return manager.resourceStack
}

// GetBehaviorStack returns the behavior pack stack.
func (manager *Manager) GetBehaviorStack() *Stack {
	return manager.behaviorStack
}

// LoadResourcePacks loads all resource packs in the `serverPath/extensions/resource_packs/` folder.
// It returns an array of errors that occurred during the loading of all resource packs.
func (manager *Manager) LoadResourcePacks() []error {
	var path = manager.serverPath + "extensions/resource_packs/"
	var files, _ = ioutil.ReadDir(path)
	var errors []error
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
			errors = append(errors, err)
			continue
		}

		err = resourcePack.ValidateManifest()
		if err != nil {
			errors = append(errors, err)
			continue
		}

		manager.resourcePacks[resourcePack.manifest.Header.UUID] = resourcePack
		manager.GetResourceStack().Push(resourcePack)
	}
	return errors
}

// LoadBehaviorPacks loads all behavior packs in the `serverPath/extensions/behavior_packs/` folder.
// It returns an array of errors that occurred during the loading of all behavior packs.
func (manager *Manager) LoadBehaviorPacks() []error {
	var path = manager.serverPath + "extensions/behavior_packs/"
	var files, _ = ioutil.ReadDir(path)
	var errors []error
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
			errors = append(errors, err)
			continue
		}

		err = behaviorPack.ValidateManifest()
		if err != nil {
			errors = append(errors, err)
			continue
		}

		err = behaviorPack.ValidateDependencies(manager)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		manager.behaviorPacks[behaviorPack.manifest.Header.UUID] = behaviorPack
		manager.GetBehaviorStack().Push(behaviorPack)
	}
	return errors
}

// IsResourcePackLoaded checks if a resource pack with the given UUID is loaded.
func (manager *Manager) IsResourcePackLoaded(uuid string) bool {
	var _, exists = manager.resourcePacks[uuid]
	return exists
}

// IsBehaviorPackLoaded checks if a behavior pack with the given UUID is loaded.
func (manager *Manager) IsBehaviorPackLoaded(uuid string) bool {
	var _, exists = manager.behaviorPacks[uuid]
	return exists
}

// IsPackLoaded checks if any pack with the given UUID is loaded.
func (manager *Manager) IsPackLoaded(uuid string) bool {
	return manager.IsResourcePackLoaded(uuid) || manager.IsBehaviorPackLoaded(uuid)
}

// GetResourcePack returns a resource pack by its UUID, or nil of none was found.
func (manager *Manager) GetResourcePack(uuid string) *ResourcePack {
	if !manager.IsResourcePackLoaded(uuid) {
		return nil
	}
	return manager.resourcePacks[uuid]
}

// GetBehaviorPack returns a behavior pack by its UUID, or nil if none was found.
func (manager *Manager) GetBehaviorPack(uuid string) *BehaviorPack {
	if !manager.IsBehaviorPackLoaded(uuid) {
		return nil
	}
	return manager.behaviorPacks[uuid]
}

// GetPack returns any pack that has the given UUID, or nil if none was found.
func (manager *Manager) GetPack(uuid string) Pack {
	if manager.GetResourcePack(uuid) != nil {
		return manager.GetResourcePack(uuid)
	}
	return manager.GetBehaviorPack(uuid)
}
