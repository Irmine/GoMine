package packs

import (
	"errors"
	"strconv"
)

// BehaviorPack is a pack used to modify the behavior of entities.
type BehaviorPack struct {
	*Base
}

// NewBehaviorPack returns a new behavior pack at the given path.
func NewBehaviorPack(path string) *BehaviorPack {
	return &BehaviorPack{newBase(path, Behavior)}
}

// ValidateDependencies validates all dependencies of the behavior pack, and returns an error if any.
func (pack *BehaviorPack) ValidateDependencies(manager *Manager) error {
	var dependencies = pack.manifest.Dependencies
	for index, dependency := range dependencies {
		if dependency.Description == "" {
			return errors.New("Dependency " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a description.")
		}

		if !manager.IsResourcePackLoaded(dependency.UUID) {
			return errors.New("Dependency with UUID: " + dependency.UUID + " is not loaded.")
		}

		if len(dependency.Version) < 2 {
			return errors.New("Dependency " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a valid version.")
		}

		if dependency.Type != string(Resource) {
			return errors.New("Dependency " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing the correct type. Expected: 'resources', got: '" + dependency.Type + "'")
		}
	}
	return nil
}
