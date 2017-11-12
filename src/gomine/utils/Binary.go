package utils

import (
	"fmt"
)

func Read(buffer *[]byte, offset *int, length int) ([]byte) {
	bytes := make([]byte, length)
	if *offset == (len(*buffer)-1) {
		fmt.Printf("An error occurred: %v", "no bytes left to Write")
		panic("Aborting...")
	}
	if length > 1 {
		for i := 0; i < length; i++ {
			bytes = append(bytes, (*buffer)[*offset])
			*offset++
		}
		return bytes
	}
	bytes = append(bytes, (*buffer)[*offset])
	*offset++
	return bytes
}

func Write(buffer *[]byte, v byte){
	*buffer = append(*buffer, v)
}

func WriteBool(buffer *[]byte, bool bool) {
	if bool {
		WriteByte(buffer, 0x01)
		return
	}
	WriteByte(buffer, 0x00)
}

func ReadBool(buffer *[]byte, offset *int) (bool) {
	out := Read(buffer, offset, 1)
	return out[0] != 0x00
}

func WriteByte(buffer *[]byte, byte byte) {
	Write(buffer, byte)
}

func ReadByte(buffer *[]byte, offset *int) (byte) {
	out := Read(buffer, offset, 1)
	return byte(out[0])
}

func WriteUnsignedByte(buffer *[]byte, unsigned uint8) {
	WriteByte(buffer, byte(unsigned))
}

func ReadUnsignedByte(buffer *[]byte, offset *int) (byte) {
	out := Read(buffer, offset, 1)
	return byte(out[0])
}

func WriteShort(buffer *[]byte, signed int16) {
	var i uint
	len2 := 2
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(signed >> i))
	}
}

func ReadShort(buffer *[]byte, offset *int) (int16) {
	var v int
	var i uint
	var out int
	bytes := Read(buffer, offset, 2)
	len2 := len(bytes)
	v = len2
	for i = 0; i < uint(len2) * 8; i += 8 {
		if i == 0 {
			out = int(bytes[v])
			continue
		}
		out |= int(bytes[v]) << i
		v--
	}
	return int16(out)
}

func WriteInt(buffer *[]byte, int int32) {
	var i uint
	len2 := 4
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(int >> i))
	}
}

func ReadInt(buffer *[]byte, offset *int) (int32) {
	var v int
	var i uint
	var out int
	bytes := Read(buffer, offset, 4)
	len2 := len(bytes)
	v = len2
	for i = 0; i < uint(len2) * 8; i += 8 {
		if i == 0 {
			out = int(bytes[v])
			continue
		}
		out |= int(bytes[v]) << i
		v--
	}
	return int32(out)
}

func WriteLong(buffer *[]byte, int int64) {
	var i uint
	len2 := 8
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(int >> i))
	}
}

func ReadLong(buffer *[]byte, offset *int) (int64) {
	var v int
	var i uint
	var out int
	bytes := Read(buffer, offset, 8)
	len2 := len(bytes)
	v = len2
	for i = 0; i < uint(len2) * 8; i += 8 {
		if i == 0 {
			out = int(bytes[v])
			continue
		}
		out |= int(bytes[v]) << i
		v--
	}
	return int64(out)
}

func WriteUnsignedLong(buffer *[]byte, int uint64) {
	var i uint
	len2 := 8
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(int >> i))
	}
}

func ReadUnsignedLong(buffer *[]byte, offset *int) (uint64) {
	var v int
	var i uint
	var out int
	bytes := Read(buffer, offset, 8)
	len2 := len(bytes)
	v = len2
	for i = 0; i < uint(len2) * 8; i += 8 {
		if i == 0 {
			out = int(bytes[v])
			continue
		}
		out |= int(bytes[v]) << i
		v--
	}
	return uint64(out)
}

func WriteFloat(buffer *[]byte, float float32) {
	var i uint
	len2 := 4
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(float) >> i))
	}
}

func ReadFloat(buffer *[]byte, offset *int) (float32) {
	var v int
	var i uint
	var out int
	bytes := Read(buffer, offset, 4)
	len2 := len(bytes)
	v = len2
	for i = 0; i < uint(len2) * 8; i += 8 {
		if i == 0 {
			out = int(bytes[v])
			continue
		}
		out |= int(bytes[v]) << i
		v--
	}
	return float32(out)
}

func WriteDouble(buffer *[]byte, double float64) {
	var i uint
	len2 := 4
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(double) >> i))
	}
}

func ReadDouble(buffer *[]byte, offset *int) (float64) {
	var v int
	var i uint
	var out int
	bytes := Read(buffer, offset, 8)
	len2 := len(bytes)
	v = len2
	for i = 0; i < uint(len2) * 8; i += 8 {
		if i == 0 {
			out = int(bytes[v])
			continue
		}
		out |= int(bytes[v]) << i
		v--
	}
	return float64(out)
}

func WriteString(buffer *[]byte, string string) {
	len2 := len(string)
	WriteVarInt(buffer, int32(len2))
	for i := 0; i < len2; i++ {
		WriteByte(buffer, byte(string[i]))
	}
}

func ReadString(buffer *[]byte, offset *int) (string) {
	bytes := Read(buffer, offset, int(ReadVarInt(buffer, offset)))
	return string(bytes)
}

func WriteVarInt(buffer *[]byte, int int32) {
	var int2 uint32
	for int2 != 0 {
		out := int & 0x7F
		int2 = uint32(int) >> 7
		if int2 != 0 {
			out |= 0x7F
		}
		WriteByte(buffer, byte(out))
	}
}

func ReadVarInt(buffer *[]byte, offset *int) (int32) {
	var out int32
	var next byte
	var bytesRead int32

	for (next & 0x7F) != 0 {
		next = ReadByte(buffer, offset)
		out |= int32(next & 0x7F) << 7 * bytesRead
		bytesRead++
		if bytesRead > 5 {
			fmt.Printf("An error occurred: var int is too big")
			panic("Aborting...")
		}
	}

	return out
}

func WriteVarLong(buffer *[]byte, int int64) {
	var int2 uint64
	for int2 != 0 {
		out := int & 0x7F
		int2 = uint64(int) >> 7
		if int2 != 0 {
			out |= 0x7F
		}
		WriteByte(buffer, byte(out))
	}
}

func ReadVarLong(buffer *[]byte, offset *int) (int64) {
	var out int64
	var next byte
	var bytesRead int64

	for (next & 0x7F) != 0 {
		next = ReadByte(buffer, offset)
		out |= int64(next & 0x7F) << 7 * bytesRead
		bytesRead++
		if bytesRead > 10 {
			fmt.Printf("An error occurred: var long is too big")
			panic("Aborting...")
		}
	}

	return out
}

func WritePosition(buffer *[]byte, x, y, z int) {
	var v int
	v = (x & 0x3FFFFFF) << 38
	v |= (y & 0xFFF) << 26
	v |= (z & 0x3FFFFFF) << 38
	WriteVarLong(buffer, int64(v))
}

func ReadPosition(buffer *[]byte, offset *int) (x, y, z int) {
	long := ReadVarLong(buffer, offset)
	x = int(long >> 38)
	y = int(long >> 26) & 0xFFF
	z = int(long << 38 >> 38)
	return x, y, z
}

func WriteEId(buffer *[]byte, eid int32) {
	WriteVarInt(buffer, eid)
}

func ReadEId(buffer *[]byte, offset *int) (eid int32) {
	eid = ReadVarInt(buffer, offset)
	return eid
}

func WriteId(buffer *[]byte, id string) {//NO IDEA WHAT THIS DOES
	WriteString(buffer, id)
}

func ReadId(buffer *[]byte, offset *int) (id string) {//OR THIS
	id = ReadString(buffer, offset)
	return id
}

func WriteUUID(buffer *[]byte, parts []int) {
	//todo
}

func ReadUUID(buffer *[]byte, offset *int) ([]int) {
	//todo
	return []int{}
}