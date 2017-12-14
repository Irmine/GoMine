package packs

import (
	"os"
	"io/ioutil"
	"crypto/sha256"
	"archive/zip"
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	Behavior = "data"
	Resource = "resources"
)

type Pack struct {
	packPath string
	manifest *PackManifest
	content []byte
	size int64
	sha256 []byte
	packType string
}

type PackManifest struct {
	Header struct {
		Description 	string `json:"description"`
		Name 			string `json:"name"`
		UUID 			string `json:"uuid"`
		Version 		[]float64 `json:"version"`
		VersionString 	string
	} `json:"header"`
	Modules []struct {
		Description		string `json:"description"`
		Type			string `json:"type"`
		UUID 			string `json:"uuid"`
		Version 		[]float64 `json:"version"`
	} `json:"modules"`
	Dependencies []struct {
		Description		string `json:"description"`
		Type 			string `json:"type"`
		UUID 			string `json:"dependencies"`
		Version 		[]float64 `json:"version"`
	} `json:"dependencies"`
}

func NewPack(path string, packType string) *Pack {
	var reader, _ = os.Open(path)
	var content, _ = ioutil.ReadAll(reader)
	var sha = sha256.Sum256(content)

	var shaBytes []byte
	for _, b := range sha {
		shaBytes = append(shaBytes, b)
	}
	return &Pack{path, &PackManifest{}, content, int64(len(content)), shaBytes, packType}
}

/**
 * Loads the resource pack.
 */
func (pack *Pack) Load() error {
	var zipFile, err = zip.OpenReader(pack.packPath)
	if err != nil {
		panic(err)
	}

	for _, file := range zipFile.File {
		if file.Name != "manifest.json" && file.Name != "pack_manifest.json" {
			continue
		}
		reader, _ := file.Open()
		bytes, _ := ioutil.ReadAll(reader)

		manifest := &PackManifest{}
		err := json.Unmarshal(bytes, manifest)
		pack.manifest = manifest

		return err
	}
	return errors.New("No manifest.json or pack_manifest.json could be found in zip: " + pack.packPath)
}

/**
 * Returns the file path of this pack.
 */
func (pack *Pack) GetPath() string {
	return pack.packPath
}

/**
 * Returns a sha256 checksum of this pack.
 */
func (pack *Pack) GetSha256() string {
	return string(pack.sha256)
}

/**
 * Returns the size of the pack in bytes.
 */
func (pack *Pack) GetFileSize() int64 {
	return pack.size
}

/**
 * Returns the UUID of this pack.
 */
func (pack *Pack) GetUUID() string {
	return pack.manifest.Header.UUID
}

/**
 * Returns the version of this pack in a readable string.
 */
func (pack *Pack) GetVersion() string {
	return pack.manifest.Header.VersionString
}

/**
 * Returns the manifest of this pack.
 */
func (pack *Pack) GetManifest() *PackManifest {
	return pack.manifest
}

/**
 * Returns the complete content of the pack in a byte slice.
 */
func (pack *Pack) GetContent() []byte {
	return pack.content
}

/**
 * Verifies the manifest header of the pack.
 */
func (pack *Pack) ValidateManifest() error {
	var manifest = pack.manifest
	if manifest.Header.Description == "" {
		return errors.New("Pack at " + pack.packPath + " is missing a description.")
	}
	if manifest.Header.Name == "" {
		return errors.New("Pack at " + pack.packPath + " is missing a name.")
	}

	var regex = regexp.MustCompile("-")
	var occurrences = regex.FindAllStringIndex(manifest.Header.UUID, -1)
	if manifest.Header.UUID == "" || len(occurrences) != 4 {
		return errors.New("Resource pack at " + pack.packPath + " does not have a valid UUID.")
	}

	if len(manifest.Header.Version) < 2 {
		return errors.New("Pack at " + pack.packPath + " is missing a valid version.")
	}

	var versionStrings []string
	for _, versionNumber := range manifest.Header.Version {
		versionStrings = append(versionStrings, strconv.Itoa(int(versionNumber)))
	}
	manifest.Header.VersionString = strings.Join(versionStrings, ".")

	return pack.ValidateModules()
}

/**
 * Validates the modules of this pack.
 */
func (pack *Pack) ValidateModules() error {
	var modules = pack.manifest.Modules
	if len(modules) == 0 {
		return errors.New("Pack at " + pack.packPath + " doesn't have any modules.")
	}

	for index, module := range modules {
		if module.Description == "" {
			return errors.New("Module " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a description.")
		}

		var regex = regexp.MustCompile("-")
		var occurrences = regex.FindAllStringIndex(module.UUID, -1)
		if module.UUID == "" || len(occurrences) != 4 {
			return errors.New("Module " + strconv.Itoa(index) + " in pack at " + pack.packPath + " does not have a valid UUID.")
		}

		if len(module.Version) < 2 {
			return errors.New("Module " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a valid version.")
		}

		if module.Type != pack.packType {
			return errors.New("Module " + strconv.Itoa(index) + " in pack at " + pack.packPath + " does not have the correct type. Expected: '" + pack.packType + "', got: '" + module.Type + "'")
		}
	}

	return nil
}

/**
 * Returns a byte slice from the content of the pack at the given offset with the given length.
 */
func (pack *Pack) GetChunk(offset int, length int) []byte {
	if offset > len(pack.content) || offset < 0 || length < 1 {
		return []byte{}
	}
	if offset + length > len(pack.content) {
		length = int(pack.size) - offset
	}
	return pack.content[offset:offset + length]
}