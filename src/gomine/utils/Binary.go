package utils

import (
	"fmt"
	"math"
)

func Read(buffer *[]byte, offset *int, length int) ([]byte) {
	bytes := make([]byte, 0)
	if length == 0 {
		return bytes
	}
	if *offset >= len(*buffer) {
		fmt.Printf("An error occurred: %v", "no bytes left to read")
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
	var v uint
	len2 := uint(2)
	v = uint(len2*8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(signed >> v))
		v -= 8
	}
}

func ReadShort(buffer *[]byte, offset *int) (int16) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 2)
	len2 := uint(len(bytes))
	v = uint(len2*8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v -= 8
			continue
		}
		out |= int(bytes[i]) << v
		v -= 8
	}
	return int16(out)
}

func WriteUnsignedShort(buffer *[]byte, int uint16) {
	var i uint
	var v uint
	len2 := uint(2)
	v = uint(len2*8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(int >> v))
		v -= 8
	}
}

func ReadUnsignedShort(buffer *[]byte, offset *int) (uint16) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 2)
	len2 := uint(len(bytes))
	v = uint(len2*8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v -= 8
			continue
		}
		out |= int(bytes[i]) << v
		v -= 8
	}
	return uint16(out)
}

func WriteInt(buffer *[]byte, int int32) {
	var i uint
	var v uint
	len2 := uint(4)
	v = uint(len2*8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(int >> v))
		v -= 8
	}
}

func ReadInt(buffer *[]byte, offset *int) (int32) {
	var v uint
	var i uint
	var out int

	bytes := Read(buffer, offset, 4)
	len2 := uint(len(bytes))
	v = uint(len2 * 8) - 8

	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v -= 8
			continue
		}
		out |= int(bytes[i]) << v
		v -= 8
	}
	return int32(out)
}

func WriteLong(buffer *[]byte, int int64) {
	var i uint
	var v uint
	len2 := uint(8)
	v = uint(len2*8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(int >> v))
		v -= 8
	}
}

func ReadLong(buffer *[]byte, offset *int) (int64) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 8)
	len2 := uint(len(bytes))
	v = uint(len2*8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v -= 8
			continue
		}
		out |= int(bytes[i]) << v
		v -= 8
	}
	return int64(out)
}

func WriteUnsignedLong(buffer *[]byte, int uint64) {
	var i uint
	var v uint
	len2 := uint(8)
	v = uint(len2*8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(int >> v))
		v -= 8
	}
}

func ReadUnsignedLong(buffer *[]byte, offset *int) (uint64) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 8)
	len2 := uint(len(bytes))
	v = uint(len2*8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v -= 8
			continue
		}
		out |= int(bytes[i]) << v
		v -= 8
	}
	return uint64(out)
}

func WriteFloat(buffer *[]byte, float float32) {
	var i uint
	var v uint
	x := math.Float32bits(float)
	len2 := uint(4)
	v = uint(len2*8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(x >> v))
		v -= 8
	}
}

func ReadFloat(buffer *[]byte, offset *int) (float32) {
	var v uint
	var i uint
	var out uint32
	bytes := Read(buffer, offset, 4)
	len2 := uint(len(bytes))
	v = uint(len2*8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = uint32(bytes[i]) << v
			v -= 8
			continue
		}
		out |= uint32(bytes[i]) << v
		v -= 8
	}
	return math.Float32frombits(out)
}

func WriteDouble(buffer *[]byte, double float64) {
	var i uint
	var v uint
	x := math.Float64bits(double)
	len2 := uint(8)
	v = uint(len2*8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(x >> v))
		v -= 8
	}
}

func ReadDouble(buffer *[]byte, offset *int) (float64) {
	var v uint
	var i uint
	var out uint64
	bytes := Read(buffer, offset, 8)
	len2 := uint(len(bytes))
	v = uint(len2*8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = uint64(bytes[i]) << v
			v -= 8
			continue
		}
		out |= uint64(bytes[i]) << v
		v -= 8
	}
	return math.Float64frombits(out)
}

//little

func WriteLittleShort(buffer *[]byte, short int16) {
	var i uint
	len2 := uint(2)
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(uint(short) >> i))
	}
}

func ReadLittleShort(buffer *[]byte, offset *int) (int16) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 2)
	len2 := uint(len(bytes))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v += 8
			continue
		}
		out |= int(bytes[i]) << v
		v += 8
	}
	return int16(out)
}

func WriteLittleUnsignedShort(buffer *[]byte, short uint16) {
	var i uint
	len2 := 2
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(short) >> i))
	}
}

func ReadLittleUnsignedShort(buffer *[]byte, offset *int) (uint16) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 2)
	len2 := uint(len(bytes))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v += 8
			continue
		}
		out |= int(bytes[i]) << v
		v += 8
	}
	return uint16(out)
}

func WriteLittleInt(buffer *[]byte, int int32) {
	var i uint
	len2 := 4
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(int) >> i))
	}
}

func ReadLittleInt(buffer *[]byte, offset *int) (int32) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 4)
	len2 := uint(len(bytes))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v += 8
			continue
		}
		out |= int(bytes[i]) << v
		v += 8
	}
	return int32(out)
}

func WriteLittleLong(buffer *[]byte, int int64) {
	var i uint
	len2 := 8
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(int) >> i))
	}
}

func ReadLittleLong(buffer *[]byte, offset *int) (int64) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 8)
	len2 := uint(len(bytes))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v += 8
			continue
		}
		out |= int(bytes[i]) << v
		v += 8
	}
	return int64(out)
}

func WriteLittleUnsignedLong(buffer *[]byte, int uint64) {
	var i uint
	len2 := 8
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(int) >> i))
	}
}

func ReadLittleUnsignedLong(buffer *[]byte, offset *int) (uint64) {
	var v uint
	var i uint
	var out int
	bytes := Read(buffer, offset, 8)
	len2 := uint(len(bytes))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(bytes[i]) << v
			v += 8
			continue
		}
		out |= int(bytes[i]) << v
		v += 8
	}
	return uint64(out)
}

func WriteLittleFloat(buffer *[]byte, f float32) {
	var i uint
	x := math.Float32bits(f)
	len2 := 4
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(x >> i))
	}
}

func ReadLittleFloat(buffer *[]byte, offset *int) (float32) {
	var v uint
	var i uint
	var out uint32
	bytes := Read(buffer, offset, 4)
	len2 := uint(len(bytes))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = uint32(bytes[i]) << v
			v += 8
			continue
		}
		out |= uint32(bytes[i]) << v
		v += 8
	}
	return math.Float32frombits(out)
}

func WriteLittleDouble(buffer *[]byte, double float64) {
	var i uint
	x := math.Float64bits(double)
	len2 := 8
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(x >> i))
	}
}

func ReadLittleDouble(buffer *[]byte, offset *int) (float64) {
	var v uint
	var i uint
	var out uint64
	bytes := Read(buffer, offset, 8)
	len2 := uint(len(bytes))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = uint64(bytes[i]) << v
			v += 8
			continue
		}
		out |= uint64(bytes[i]) << v
		v += 8
	}
	return math.Float64frombits(out)
}

func WriteString(buffer *[]byte, string string) {
	len := len(string)
	WriteUnsignedVarInt(buffer, uint32(len))
	for i := 0; i < len; i++ {
		WriteByte(buffer, byte(string[i]))
	}
}

func ReadString(buffer *[]byte, offset *int) (string) {
	bytes := Read(buffer, offset, int(ReadUnsignedVarInt(buffer, offset)))
	return string(bytes)
}

func WriteVarInt(buffer *[]byte, int int32) {
	int = int << 1
	for u := 0; u < 5; u++ {
		if (int >> 7) != 0 {
			Write(buffer, byte(int | 0x80))
		} else {
			Write(buffer, byte(int & 0x7f))
			break
		}
		int >>= 7
	}
}

func ReadVarInt(buffer *[]byte, offset *int) int32 {
	var out int32 = 0
	for v := uint(0); v < 35; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= int32(b << v)

		if (b & 0x80) == 0 {
			return out >> 1
		}
	}

	return 0
}

func WriteVarLong(buffer *[]byte, int int64) {
	int = int << 1
	for u := 0; u < 10; u++ {
		if (int >> 7) != 0 {
			Write(buffer, byte(int | 0x80))
		} else {
			Write(buffer, byte(int & 0x7f))
			break
		}
		int >>= 7
	}
}

func ReadVarLong(buffer *[]byte, offset *int) (int64) {
	var out int64 = 0
	for v := uint(0); v <= 63; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= int64(b << v)

		if (b & 0x80) == 0 {
			return out >> 1
		}
	}

	return 0
}

func WriteUnsignedVarInt(buffer *[]byte, int uint32) {
	for u := 0; u < 5; u++ {
		if (int >> 7) != 0 {
			Write(buffer, byte(int | 0x80))
		} else {
			Write(buffer, byte(int & 0x7f))
			break
		}
		int >>= 7
	}
}

func ReadUnsignedVarInt(buffer *[]byte, offset *int) (uint32) {
	var out uint32 = 0
	for v := uint(0); v < 35; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= uint32(b << v)

		if (b & 0x80) == 0 {
			return out
		}
	}

	return 0
}

func WriteUnsignedVarLong(buffer *[]byte, int uint64) {
	for u := 0; u < 10; u++ {
		if (int >> 7) != 0 {
			Write(buffer, byte(int | 0x80))
		} else {
			Write(buffer, byte(int & 0x7f))
			break
		}
		int >>= 7
	}
}

func ReadUnsignedVarLong(buffer *[]byte, offset *int) (uint64) {
	var out uint64 = 0
	for v := uint(0); v <= 63; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= uint64(b << v)

		if (b & 0x80) == 0 {
			return out
		}
	}

	return 0
}

func ReadBigEndianTriad(buffer *[]byte, offset *int) uint32 {
	var out uint32
	var bytes = Read(buffer, offset, 3)
	out = uint32(bytes[0] | (bytes[1] << 8) | (bytes[2] << 16))

	return out
}

func WriteBigEndianTriad(buffer *[]byte, uint uint32) {
	Write(buffer, byte(uint & 0xFF))
	Write(buffer, byte(uint >> 8 & 0xFF))
	Write(buffer, byte(uint >> 16))
}

func ReadLittleEndianTriad(buffer *[]byte, offset *int) uint32 {
	var out uint32
	var bytes = Read(buffer, offset, 3)
	out = uint32(bytes[2] | (bytes[1] << 8) | (bytes[0] << 16))

	return out
}

func WriteLittleEndianTriad(buffer *[]byte, uint uint32) {
	Write(buffer, byte(uint >> 16))
	Write(buffer, byte(uint >> 8 & 0xFF))
	Write(buffer, byte(uint & 0xFF))
}