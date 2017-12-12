package packs

import (
	"archive/zip"
	"io/ioutil"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"strconv"
	"os"
)

type ResourcePack struct {
	packPath string
	manifest *PackManifest
	content []byte
}

type PackManifest struct {
	Header struct {
		Description string	`json:"description"`
		Name string			`json:"name"`
		UUID string			`json:"uuid"`
		Version []float64	`json:"version"`
		VersionString string
	}	`json:"header"`
	Modules []map[string]interface{}	`json:"modules"`
}

/**
 * Returns a new resource pack to the given path.
 */
func NewResourcePack(path string) *ResourcePack {
	var reader, _ = os.Open(path)
	var content, _ = ioutil.ReadAll(reader)
	return &ResourcePack{path, &PackManifest{}, content}
}

/**
 * Returns the file path of this resource pack.
 */
func (pack *ResourcePack) GetPath() string {
	return pack.packPath
}

/**
 * Returns the manifest of this resource pack.
 */
func (pack *ResourcePack) GetManifest() *PackManifest {
	return pack.manifest
}

/**
 * Returns the complete content of the resource pack in a byte slice.
 */
func (pack *ResourcePack) GetContent() []byte {
	return pack.content
}

/**
 * Loads the resource pack.
 */
func (pack *ResourcePack) Load() {
	var zipFile, _ = zip.OpenReader(pack.packPath)
	for _, file := range zipFile.File {
		if file.Name != "manifest.json" && file.Name != "pack_manifest.json" {
			continue
		}
		reader, _ := file.Open()
		bytes, _ := ioutil.ReadAll(reader)

		manifest := &PackManifest{}
		json.Unmarshal(bytes, manifest)
		pack.manifest = manifest
	}
}

/**
 * Verifies the manifest file of the resource pack.
 */
func (pack *ResourcePack) Validate() error {
	var manifest = pack.manifest
	if manifest.Header.Description == "" {
		return errors.New("Resource pack at " + pack.packPath + " is missing a description.")
	}
	if manifest.Header.Name == "" {
		return errors.New("Resource pack at " + pack.packPath + " is missing a name.")
	}

	var regex = regexp.MustCompile("-")
	var occurrences = regex.FindAllStringIndex(manifest.Header.UUID, -1)
	if manifest.Header.UUID == "" || len(occurrences) != 4 {
		return errors.New("Resource pack at " + pack.packPath + " does not have a valid UUID.")
	}

	if len(manifest.Header.Version) == 0 {
		return errors.New("Resource pack at " + pack.packPath + " is missing a version.")
	}

	var versionStrings []string
	for _, versionNumber := range manifest.Header.Version {
		versionStrings = append(versionStrings, strconv.Itoa(int(versionNumber)))
	}
	manifest.Header.VersionString = strings.Join(versionStrings, ".")

	return nil
}

/**
 * Returns a byte slice from the content of the pack at the given offset with the given length.
 */
func (pack *ResourcePack) GetChunk(offset int, length int) []byte {
	if offset > len(pack.content) || offset + length > len(pack.content) || offset < 0 || length < 1 {
		return []byte{}
	}
	return pack.content[offset:offset + length]
}