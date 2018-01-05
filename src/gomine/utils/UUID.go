package utils

import (
	"encoding/hex"
	"strings"
	"math/rand"
	"time"
)

type UUID struct {
	parts [4]int32
	version  int
}

var validDigits = "abcdefghijklmnopqrstuvwxyz0123456789"

func NewUUID(parts [4]int32) UUID {
	v := UUID{}
	v.parts = parts
	v.version = (int(parts[1]) & 0xf000) >> 12
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
	var random = rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			uuid += "-"
		} else {
			offset := random.Intn(35)
			uuid += string(validDigits[offset])
		}
	}

	return uuid
}

func (uuid *UUID) GetParts() [4]int32 {
	return uuid.parts
}

func (uuid *UUID) GetVersion() int {
	return uuid.version
}

func (uuid *UUID) SetVersion(version int) {
	uuid.version = version
}

func (uuid *UUID) Equals(uuid2 UUID) bool {
	return uuid.parts == uuid2.parts
}

func UUIDFromString(str string) UUID {
	if !IsValidUUID(str) {
		return UUID{}
	}
	println(str)
	var bytes, _ = hex.DecodeString(strings.Replace(str, "-", "", -1))
	var offset = 0
	return UUIDFromBinary(&bytes, &offset)
}

func UUIDFromBinary(buffer *[]byte, offset *int) UUID {
	if len(*buffer) < 16 {
		panic("UUID is not 16 bytes long")
	}
	return NewUUID([4]int32{ReadInt(buffer, offset), ReadInt(buffer, offset), ReadInt(buffer, offset), ReadInt(buffer, offset)})
}

func (uuid *UUID) ToBinary() []byte {
	var buffer []byte
	for i := 0; i < 4; i++ {
		WriteInt(&buffer, uuid.parts[i])
	}
	return buffer
}

func (uuid *UUID) StringValue() string {
	v := uuid.ToBinary()
	out := hex.EncodeToString(v)
	res := out[0:4] + "-"
	res += out[4:8] + "-"
	res += out[8:12] + "-"
	res += out[12:16] + "-"
	res += out[16:20]
	return out
}