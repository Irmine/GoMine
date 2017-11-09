package utils

import (
	"errors"
	"fmt"
)

func get(buffer *[]byte, offset *int, length int) (byte, error) {

	if (len(*buffer)-1) == *offset {
		return 0, errors.New("no bytes left to read")
	}

	if length > 1 {
		var out byte
		for i := 0; i < length; i++ {
			out += (*buffer)[*offset]
			*offset++
		}
		return out, nil
	}

	out := (*buffer)[*offset]
	*offset++
	return out, nil
}

func WriteBool(buffer *[]byte, bool bool) {
	if bool {
		*buffer = append(*buffer, 0x01)
		return
	}
	*buffer = append(*buffer, 0x00)
}

func ReadBool(buffer *[]byte, offset *int, length int) (bool) {
	out, err := get(buffer, offset, length)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		panic("Aborting...")
	}
	return out != 0x00
}

func ReadByte(buffer *[]byte, offset *int, length int) (byte) {
	out, err := get(buffer, offset, length)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		panic("Aborting...")
	}
	return byte(out)
}

func WriteByte(buffer *[]byte, byte byte) {
	*buffer = append(*buffer, byte)
}

func ReadInt(buffer *[]byte, offset *int, length int) (int) {
	out, err := get(buffer, offset, length)
	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		panic("Aborting...")
	}
	return int(out)
}

func WriteInt(buffer *[]byte, int int) {
	*buffer = append(*buffer, byte(int))
}

//todo: add left methods
