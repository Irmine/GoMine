package utils

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type UUID struct {
	parts [2]int64
	version  int
}

func NewUUID(parts [2]int64, version int) UUID {
	v := UUID{}
	v.parts = parts
	v.version = version
	return v
}

func (uuid *UUID) Parts() [2]int64 {
	return uuid.parts
}

func (uuid *UUID) Version() int {
	return uuid.version
}

func (uuid *UUID) Equal(uuid2 UUID) bool {
	return uuid.parts == uuid2.parts
}

func (uuid *UUID) OutOfString(uuid2 string, version int) UUID {
	return uuid.OutOfBinary(strings.Replace(uuid2, "-", "", -1), version)
}

func (uuid *UUID) OutOfBinary(uuid2 string, version int) UUID {
	if len(uuid2) != 16 {
		panic("UUID is not 16 bytes long")
	}
	out := NewUUID([2]int64{}, version)
	bytes, err := hex.DecodeString(uuid2)
	offset := 0

	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		panic("Aborting...")
	}

	out.parts[0] = ReadLong(&bytes, &offset)
	out.parts[1] = ReadLong(&bytes, &offset)

	return out
}

func (uuid *UUID) StringValue() string {
	v := uuid.BinaryValue()
	out := hex.EncodeToString(v)
	res := out[0:4] + "-"
	res += out[4:8] + "-"
	res += out[8:12] + "-"
	res += out[12:16] + "-"
	res += out[16:20]
	return out
}

func (uuid *UUID) BinaryValue() []byte {
	var offset int
	offset = 0
	buffer := make([]byte, offset)
	WriteLong(&buffer, uuid.parts[0])
	WriteLong(&buffer, uuid.parts[1])
	return buffer
}