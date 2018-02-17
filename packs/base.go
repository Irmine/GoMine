package packs

import (
	"os"
	"io/ioutil"
	"crypto/sha256"
	"archive/zip"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

const (
	Behavior PackType = "data"
	Resource PackType = "resources"
)

// PackType is a name of a pack type.
type PackType string

// Pack is the main interface which both Resource- and BehaviorPack satisfy.
type Pack interface {
	GetUUID() string
	GetVersion() string
	GetFileSize() int64
	GetSha256() string
	GetChunk(offset int, length int) []byte
	GetPath() string
}

// Base is a struct that forms the base of every pack.
// It has functions for loading, validating and pack data.
type Base struct {
	packPath string
	manifest *Manifest
	content  []byte
	size     int64
	sha256   []byte
	packType PackType
}

// Manifest is a struct that contains all information of a pack.
type Manifest struct {
	Header struct {
		Description   string    `json:"description"`
		Name          string    `json:"name"`
		UUID          string    `json:"uuid"`
		Version       []float64 `json:"version"`
		VersionString string
	} `json:"header"`
	Modules []struct {
		Description string    `json:"description"`
		Type        string    `json:"type"`
		UUID        string    `json:"uuid"`
		Version     []float64 `json:"version"`
	} `json:"modules"`
	Dependencies []struct {
		Description string    `json:"description"`
		Type        string    `json:"type"`
		UUID        string    `json:"uuid"`
		Version     []float64 `json:"version"`
	} `json:"dependencies"`
}

// newBase returns a new base at the given path and with the given pack type.
func newBase(path string, packType PackType) *Base {
	var reader, _ = os.Open(path)
	var content, _ = ioutil.ReadAll(reader)
	var sha = sha256.Sum256(content)

	var shaBytes []byte
	for _, b := range sha {
		shaBytes = append(shaBytes, b)
	}
	return &Base{path, &Manifest{}, content, int64(len(content)), shaBytes, packType}
}

// Load loads the pack, and returns an error if any.
func (pack *Base) Load() error {
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

		manifest := &Manifest{}
		err := json.Unmarshal(bytes, manifest)
		pack.manifest = manifest

		reader.Close()

		return err
	}
	return errors.New("No manifest.json or pack_manifest.json could be found in zip: " + pack.packPath)
}

// GetPath returns the path of the pack.
func (pack *Base) GetPath() string {
	return pack.packPath
}

// GetSha256 returns the Sha256 checksum of the pack.
func (pack *Base) GetSha256() string {
	return string(pack.sha256)
}

// GetFileSize returns the file size of the pack.
func (pack *Base) GetFileSize() int64 {
	return pack.size
}

// GetUUID returns the UUID of the pack.
func (pack *Base) GetUUID() string {
	return pack.manifest.Header.UUID
}

// GetVersion returns the version string of the pack.
func (pack *Base) GetVersion() string {
	return pack.manifest.Header.VersionString
}

// GetManifest returns the manifest of the pack.
func (pack *Base) GetManifest() *Manifest {
	return pack.manifest
}

// GetContent returns the full byte array of the data of the pack.
func (pack *Base) GetContent() []byte {
	return pack.content
}

// ValidateManifest validates the manifest, and returns an error if any.
func (pack *Base) ValidateManifest() error {
	var manifest = pack.manifest
	if manifest.Header.Description == "" {
		return errors.New("Pack at " + pack.packPath + " is missing a description.")
	}
	if manifest.Header.Name == "" {
		return errors.New("Pack at " + pack.packPath + " is missing a name.")
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

// ValidateModules validates the modules of the pack, and returns an error if any.
func (pack *Base) ValidateModules() error {
	var modules = pack.manifest.Modules
	if len(modules) == 0 {
		return errors.New("Pack at " + pack.packPath + " doesn't have any modules.")
	}

	for index, module := range modules {
		if module.Description == "" {
			return errors.New("Module " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a description.")
		}

		if len(module.Version) < 2 {
			return errors.New("Module " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a valid version.")
		}

		if module.Type == "" {
			return errors.New("Module " + strconv.Itoa(index) + " in pack at " + pack.packPath + " is missing a valid type.")
		}
	}

	return nil
}

// GetChunk returns a chunk of the pack at the given offset with the given length.
func (pack *Base) GetChunk(offset int, length int) []byte {
	if offset > len(pack.content) || offset < 0 || length < 1 {
		return []byte{}
	}
	if offset+length > len(pack.content) {
		length = int(pack.size) - offset
	}
	return pack.content[offset:offset+length]
}
