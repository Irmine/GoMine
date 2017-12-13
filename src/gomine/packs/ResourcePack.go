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
	"crypto/sha256"
)

type ResourcePack struct {
	packPath string
	manifest *PackManifest
	content []byte
	size int64
	sha256 []byte
}

type PackManifest struct {
	Header struct {
		Description 	string				`json:"description"`
		Name 			string				`json:"name"`
		UUID 			string				`json:"uuid"`
		Version 		[]float64			`json:"version"`
		VersionString 	string
	}									`json:"header"`
	Modules []map[string]interface{}	`json:"modules"`
}

/**
 * Returns a new resource pack to the given path.
 */
func NewResourcePack(path string) *ResourcePack {
	var reader, _ = os.Open(path)
	var content, _ = ioutil.ReadAll(reader)
	var sha = sha256.Sum256(content)

	var shaBytes []byte
	for _, b := range sha {
		shaBytes = append(shaBytes, b)
	}
	return &ResourcePack{path, &PackManifest{}, content, int64(len(content)), shaBytes}
}

/**
 * Returns the file path of this resource pack.
 */
func (pack *ResourcePack) GetPath() string {
	return pack.packPath
}

/**
 * Returns a sha256 checksum of this resource pack.
 */
func (pack *ResourcePack) GetSha256() string {
	return string(pack.sha256)
}

/**
 * Returns the size of the pack in bytes.
 */
func (pack *ResourcePack) GetFileSize() int64 {
	return pack.size
}

/**
 * Returns the UUID of this pack.
 */
func (pack *ResourcePack) GetUUID() string {
	return pack.manifest.Header.UUID
}

/**
 * Returns the version of this pack in a readable string.
 */
func (pack *ResourcePack) GetVersion() string {
	return pack.manifest.Header.VersionString
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
func (pack *ResourcePack) Load() error {
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
	if offset > len(pack.content) || offset < 0 || length < 1 {
		return []byte{}
	}
	if offset + length > len(pack.content) {
		length = int(pack.size) - offset
	}
	return pack.content[offset:offset + length]
}