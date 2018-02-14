package utils

import (
	"math"
)

func Read(buffer *[]byte, offset *int, length int) []byte {
	var b = (*buffer)[*offset:*offset + length]
	*offset += length
	return b
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

func ReadBool(buffer *[]byte, offset *int) bool {
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

func ReadUnsignedByte(buffer *[]byte, offset *int) byte {
	out := Read(buffer, offset, 1)
	return byte(out[0])
}

func WriteShort(buffer *[]byte, signed int16) {
	var b = make([]byte, 2)
	var v = uint16(signed)
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadShort(buffer *[]byte, offset *int) int16 {
	b := Read(buffer, offset, 2)
	return int16(uint16(b[1]) | uint16(b[0]) << 8)
}

func WriteUnsignedShort(buffer *[]byte, v uint16) {
	var b = make([]byte, 2)
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadUnsignedShort(buffer *[]byte, offset *int) uint16 {
	b := Read(buffer, offset, 2)
	return uint16(b[1]) | uint16(b[0]) << 8
}

func WriteInt(buffer *[]byte, int int32) {
	var b = make([]byte, 4)
	var v = uint32(int)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadInt(buffer *[]byte, offset *int) int32 {
	b := Read(buffer, offset, 4)
	return int32(uint32(b[3]) | uint32(b[2]) << 8 | uint32(b[1]) << 16 | uint32(b[0]) << 24)
}

func WriteLong(buffer *[]byte, long int64) {
	var b = make([]byte, 8)
	var v = uint64(long)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLong(buffer *[]byte, offset *int) int64 {
	b := Read(buffer, offset, 8)
	return int64(uint64(b[7]) | uint64(b[6]) << 8 | uint64(b[5]) << 16 | uint64(b[4]) << 24 |
		uint64(b[3]) << 32 | uint64(b[2]) << 40 | uint64(b[1]) << 48 | uint64(b[0]) << 56)
}

func WriteUnsignedLong(buffer *[]byte, v uint64) {
	var b = make([]byte, 8)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadUnsignedLong(buffer *[]byte, offset *int) uint64 {
	b := Read(buffer, offset, 8)
	return uint64(b[7]) | uint64(b[6]) << 8 | uint64(b[5]) << 16 | uint64(b[4]) << 24 |
		uint64(b[3]) << 32 | uint64(b[2]) << 40 | uint64(b[1]) << 48 | uint64(b[0]) << 56
}

func WriteFloat(buffer *[]byte, float float32) {
	var b = make([]byte, 4)
	var v = math.Float32bits(float)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadFloat(buffer *[]byte, offset *int) float32 {
	b := Read(buffer, offset, 4)

	var out = uint32(b[3]) | uint32(b[2]) << 8 | uint32(b[1]) << 16 | uint32(b[0]) << 24
	return math.Float32frombits(out)
}

func WriteDouble(buffer *[]byte, double float64) {
	var b = make([]byte, 8)
	var v = math.Float64bits(double)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadDouble(buffer *[]byte, offset *int) float64 {
	b := Read(buffer, offset, 8)
	var out = uint64(b[7]) | uint64(b[6]) << 8 | uint64(b[5]) << 16 | uint64(b[4]) << 24 |
		uint64(b[3]) << 32 | uint64(b[2]) << 40 | uint64(b[1]) << 48 | uint64(b[0]) << 56
	return math.Float64frombits(out)
}

func WriteLittleShort(buffer *[]byte, signed int16) {
	var b = make([]byte, 2)
	var v = uint16(signed)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleShort(buffer *[]byte, offset *int) int16 {
	b := Read(buffer, offset, 2)
	return int16(uint16(b[0]) | uint16(b[1]) << 8)
}

func WriteLittleUnsignedShort(buffer *[]byte, v uint16) {
	var b = make([]byte, 2)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleUnsignedShort(buffer *[]byte, offset *int) uint16 {
	b := Read(buffer, offset, 2)
	return uint16(b[0]) | uint16(b[1]) << 8
}

func WriteLittleInt(buffer *[]byte, int int32) {
	var b = make([]byte, 4)
	var v = uint32(int)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleInt(buffer *[]byte, offset *int) int32 {
	b := Read(buffer, offset, 4)
	return int32(uint32(b[0]) | uint32(b[1]) << 8 | uint32(b[2]) << 16 | uint32(b[3]) << 24)
}

func WriteLittleLong(buffer *[]byte, long int64) {
	var b = make([]byte, 8)
	var v = uint64(long)
	b[7] = byte(v >> 56)
	b[6] = byte(v >> 48)
	b[5] = byte(v >> 40)
	b[4] = byte(v >> 32)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleLong(buffer *[]byte, offset *int) int64 {
	b := Read(buffer, offset, 8)
	return int64(uint64(b[0]) | uint64(b[1]) << 8 | uint64(b[2]) << 16 | uint64(b[3]) << 24 |
		uint64(b[4]) << 32 | uint64(b[5]) << 40 | uint64(b[6]) << 48 | uint64(b[7]) << 56)
}

func WriteLittleUnsignedLong(buffer *[]byte, v uint64) {
	var b = make([]byte, 8)
	b[7] = byte(v >> 56)
	b[6] = byte(v >> 48)
	b[5] = byte(v >> 40)
	b[4] = byte(v >> 32)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleUnsignedLong(buffer *[]byte, offset *int) uint64 {
	b := Read(buffer, offset, 8)
	return uint64(b[0]) | uint64(b[1]) << 8 | uint64(b[2]) << 16 | uint64(b[3]) << 24 |
		uint64(b[4]) << 32 | uint64(b[5]) << 40 | uint64(b[6]) << 48 | uint64(b[7]) << 56
}

func WriteLittleFloat(buffer *[]byte, float float32) {
	var b = make([]byte, 4)
	var v = math.Float32bits(float)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleFloat(buffer *[]byte, offset *int) float32 {
	b := Read(buffer, offset, 4)

	var out = uint32(b[0]) | uint32(b[1]) << 8 | uint32(b[2]) << 16 | uint32(b[3]) << 24
	return math.Float32frombits(out)
}

func WriteLittleDouble(buffer *[]byte, double float64) {
	var b = make([]byte, 8)
	var v = math.Float64bits(double)
	b[7] = byte(v >> 56)
	b[6] = byte(v >> 48)
	b[5] = byte(v >> 40)
	b[4] = byte(v >> 32)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	*buffer = append(*buffer, b...)
}

func ReadLittleDouble(buffer *[]byte, offset *int) float64 {
	b := Read(buffer, offset, 8)
	var out = uint64(b[0]) | uint64(b[1]) << 8 | uint64(b[2]) << 16 | uint64(b[3]) << 24 |
		uint64(b[4]) << 32 | uint64(b[5]) << 40 | uint64(b[6]) << 48 | uint64(b[7]) << 56
	return math.Float64frombits(out)
}

func ReadBigTriad(buffer *[]byte, offset *int) uint32 {
	var out uint32
	var b = Read(buffer, offset, 3)
	out = (uint32(b[2]) & 0xFF) | ((uint32(b[1]) & 0xFF) << 8) | ((uint32(b[0]) & 0x0F) << 16)

	return out
}

func WriteLittleTriad(buffer *[]byte, uint uint32) {
	Write(buffer, byte(uint & 0xFF))
	Write(buffer, byte(uint >> 8) & 0xFF)
	Write(buffer, byte(uint >> 16) & 0xFF)
}

func ReadLittleTriad(buffer *[]byte, offset *int) uint32 {
	var b = Read(buffer, offset, 3)

	return (uint32(b[0]) & 0xFF) | ((uint32(b[1]) & 0xFF) << 8) | ((uint32(b[2]) & 0x0F) << 16)
}

func WriteBigTriad(buffer *[]byte, uint uint32) {
	Write(buffer, byte(uint >> 16) & 0xFF)
	Write(buffer, byte(uint >> 8) & 0xFF)
	Write(buffer, byte(uint & 0xFF))
}

func WriteVarInt(buffer *[]byte, int int32) {
	int <<= 32 >> 32
	int = (int << 1) ^ (int >> 31)

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
	for v := uint(0); v <= 35; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= int32((b & 0x7f) << v)

		if (b & 0x80) == 0 {
			break
		}
	}

	var out2 = (((out << 32) >> 32) ^ out) >> 1
	return out2 ^ (out & (1 << 30))
}

func WriteVarLong(buffer *[]byte, int int64) {
	int <<= 64 >> 64
	int = (int << 1) ^ (int >> 63)

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

func ReadVarLong(buffer *[]byte, offset *int) int64 {
	var out int64 = 0
	for v := uint(0); v <= 70; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= int64((b & 0x7f) << v)

		if (b & 0x80) == 0 {
			break
		}
	}

	var out2 = (((out << 64) >> 64) ^ out) >> 1
	return out2 ^ (out & (1 << 62))
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

func ReadUnsignedVarInt(buffer *[]byte, offset *int) uint32 {
	var out uint32 = 0
	for v := uint(0); v <= 35; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= uint32((b & 0x7f) << v)

		if (b & 0x80) == 0 {
			break
		}
	}

	return out
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

func ReadUnsignedVarLong(buffer *[]byte, offset *int) uint64 {
	var out uint64 = 0
	for v := uint(0); v <= 70; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= uint64((b & 0x7f) << v)

		if (b & 0x80) == 0 {
			return out
		}
	}

	return 0
}