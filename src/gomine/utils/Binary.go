package utils

import (
	"errors"
)

func WriteBool(buffer *[]byte, bool bool){
	if bool{
		*buffer = append(*buffer, 0x01)
		return
	}
	*buffer = append(*buffer, 0x00)
}

func ReadBool(buffer *[]byte, offset *int) (bool, error){
	if len(*buffer) == 0{
		return false, errors.New("No bytes left to read")
	}
	*offset++
	return (*buffer)[*offset] != 0x00, nil
}

func ReadByte(buffer *[]byte, offset *int) (byte, error){
	if len(*buffer) == 0{
		return 0, errors.New("No bytes left to read")
	}
	*offset++
	return byte((*buffer)[*offset]), nil
}

func WriteByte(buffer *[]byte, byte byte){
	*buffer = append(*buffer, byte)
}

func ReadInt(buffer *[]byte, offset *int) (int, error){
	if len(*buffer) == 0{
		return 0, errors.New("No bytes left to read")
	}
	*offset++
	return int((*buffer)[*offset]), nil
}

func WriteInt(buffer *[]byte, int int){
	*buffer = append(*buffer, byte(int))
}

//todo: add left methods