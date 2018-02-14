package packs

import (
	"errors"
	"strconv"

	"github.com/irmine/gomine/utils"
)

type BehaviorPack struct {
	*Pack
}

/**
 * Returns a new behaviour pack to the given path.
 */
func NewBehaviorPack(path string) *BehaviorPack {
	return &BehaviorPack{NewPack(path, Behavior)}
}

/**
 * Validates all dependencies of this pack.
 */
func (pack *BehaviorPack) ValidateDependencies(handler *PackHandler) error {
	var dependencies = pack.manifest.Dependencies
	for index, dependency := range dependencies {
		if dependency.Description == "" {
			return errors.New("Dependency " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a description.")
		}

		if !utils.IsValidUUID(dependency.UUID) {
			return errors.New("Dependency " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a valid UUID.")
		}

		if !handler.IsResourcePackLoaded(dependency.UUID) {
			return errors.New("Dependency with UUID: " + dependency.UUID + " is not loaded.")
		}

		if len(dependency.Version) < 2 {
			return errors.New("Dependency " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a valid version.")
		}

		if dependency.Type != Resource {
			return errors.New("Dependency " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing the correct type. Expected: 'resources', got: '" + dependency.Type + "'")
		}
	}
	return nil
}
