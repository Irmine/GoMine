// GoNBT is a Go package for reading/writing and manipulating the NBT storage format.
// It is used by both Minecraft Java and Bedrock Editions.
package gonbt

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io/ioutil"
)

const (
	CompressionNone = 0
	CompressionGzip = 1
	CompressionZlib = 2
)

// The NBTReader is used for reading compressed/uncompressed NBT into a compound.
// Network can be set to true to compact values, so ints become varints for example.
// EndianType should be either LittleEndian (0) or BigEndian (1), depending on the byte order of the NBT read.
type NBTReader struct {
	*BinaryStream
}

func NewNBTReader(buffer []byte, network bool, endianType byte) *NBTReader {
	return &NBTReader{NewStream(buffer, network, endianType & 0x01)}
}


// ReadUncompressedIntoCompound reads an entire uncompressed NBT buffer into a Compound.
// Returns nil if the first tag was not a compound.
func (reader *NBTReader) ReadUncompressedIntoCompound() *Compound {
	var tag = reader.GetTag()
	if compound, ok := tag.(*Compound); ok {
		compound.Read(reader)
		return compound
	}
	return nil
}

// ReadIntoCompound reads the entire NBT buffer into a compound.
// If the buffer is gzip compressed, it will decompress it.
// This function returns either a compound or nil, when the NBT is not valid.
func (reader *NBTReader) ReadIntoCompound(compression int) *Compound {
	if compression == CompressionNone {
		return reader.ReadUncompressedIntoCompound()
	}

	var data []byte
	if compression == CompressionGzip {
		var gz, _ = gzip.NewReader(bytes.NewBuffer(reader.Buffer))
		data, _ = ioutil.ReadAll(gz)
		gz.Close()
	} else {
		var zl, _ = zlib.NewReader(bytes.NewBuffer(reader.Buffer))
		data, _ = ioutil.ReadAll(zl)
		zl.Close()
	}

	reader.Buffer = data
	return reader.ReadUncompressedIntoCompound()
}


// GetTag returns the named tag at the current offset in the buffer.
// This does not read tag data into the tag on its own.
func (reader *NBTReader) GetTag() INamedTag {
	if reader.Feof() {
		return NewEnd("")
	}

	var tagId = reader.GetByte()
	var tagCheck = GetTagById(tagId, "")

	if tagCheck == nil {
		return nil
	}

	var name = reader.GetString()

	return GetTagById(tagId, name)
}
