package utils

import (
	"encoding/hex"
	"fmt"
	"strings"
	"math/rand"
	"time"
)

type UUID struct {
	parts [2]int64
	version  int
}

var validDigits = "abcdefghijklmnopqrstuvwxyz0123456789"

func NewUUID(parts [2]int64, version int) UUID {
	v := UUID{}
	v.parts = parts
	v.version = version
	return v
}

func IsValidUUID(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}
	var parts = []string{uuid[0:7], uuid[9:12], uuid[14:17], uuid[19:22], uuid[24:35]}
	var separators = string(uuid[8]) + string(uuid[13]) + string(uuid[18]) + string(uuid[23])

	for _, char := range separators {
		if string(char) != "-" {
			return false
		}
	}

	for _, part := range parts {
		for _, char := range part {
			if !strings.Contains(validDigits, string(char)) {
				return false
			}
		}
	}

	return true
}

func GenerateRandomUUID() string {
	var uuid = ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			uuid += "-"
		} else {
			offset := rand.Intn(35)
			uuid += string(validDigits[offset])
		}
	}
	return uuid
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