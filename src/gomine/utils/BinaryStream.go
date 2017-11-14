package utils

type BinaryStream struct {
	buffer []byte
	offset int
}

func NewStream() *BinaryStream {
	return &BinaryStream{make([]byte, 4096), 0}
}

func (stream *BinaryStream) Feof() bool {
	return stream.offset >= len(stream.buffer)
}

//big

func (stream *BinaryStream) PutBool(v bool) {
	WriteBool(&stream.buffer, v)
}

func (stream *BinaryStream) GetBool() bool {
	return ReadBool(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutByte(v byte) {
	WriteByte(&stream.buffer, v)
}

func (stream *BinaryStream) GetByte() byte {
	return ReadByte(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutUnsignedByte(v byte) {
	WriteUnsignedByte(&stream.buffer, v)
}

func (stream *BinaryStream) GetUnsignedByte() byte {
	return ReadUnsignedByte(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutShort(v int16) {
	WriteShort(&stream.buffer, v)
}

func (stream *BinaryStream) GetShort() int16 {
	return ReadShort(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutUnsignedShort(v uint16) {
	WriteUnsignedShort(&stream.buffer, v)
}

func (stream *BinaryStream) GetUnsignedShort() uint16 {
	return ReadUnsignedShort(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutInt(v int32) {
	WriteInt(&stream.buffer, v)
}

func (stream *BinaryStream) GetInt() int32 {
	return ReadInt(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLong(v int64) {
	WriteLong(&stream.buffer, v)
}

func (stream *BinaryStream) GetLong() int64 {
	return ReadLong(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutUnsignedLong(v uint64) {
	WriteUnsignedLong(&stream.buffer, v)
}

func (stream *BinaryStream) GetUnsignedLong() uint64 {
	return ReadUnsignedLong(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutFloat(v float32) {
	WriteFloat(&stream.buffer, v)
}

func (stream *BinaryStream) GetFloat() float32 {
	return ReadFloat(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutDouble(v float64) {
	WriteDouble(&stream.buffer, v)
}

func (stream *BinaryStream) GetDouble() float64 {
	return ReadDouble(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutVarInt(v int32) {
	WriteVarInt(&stream.buffer, v)
}

func (stream *BinaryStream) GetVarInt() int32 {
	return ReadVarInt(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutVarLong(v int64) {
	WriteVarLong(&stream.buffer, v)
}

func (stream *BinaryStream) GetVarLong() int64 {
	return ReadVarLong(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLittleShort(v int16) {
	WriteLittleShort(&stream.buffer, v)
}

func (stream *BinaryStream) GetLittleShort() int16 {
	return ReadLittleShort(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLittleUnsignedShort(v uint16) {
	WriteLittleUnsignedShort(&stream.buffer, v)
}

func (stream *BinaryStream) GetLittleUnsignedShort() uint16 {
	return ReadLittleUnsignedShort(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLittleInt(v int32) {
	WriteLittleInt(&stream.buffer, v)
}

func (stream *BinaryStream) GetLittleInt() int32 {
	return ReadLittleInt(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLittleLong(v int64) {
	WriteLittleLong(&stream.buffer, v)
}

func (stream *BinaryStream) GetLittleLong() int64 {
	return ReadLittleLong(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLittleUnsignedLong(v uint64) {
	WriteLittleUnsignedLong(&stream.buffer, v)
}

func (stream *BinaryStream) GetLittleUnsignedLong() uint64 {
	return ReadLittleUnsignedLong(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLittleFloat(v float32) {
	WriteLittleFloat(&stream.buffer, v)
}

func (stream *BinaryStream) GetLittleFloat() float32 {
	return ReadLittleFloat(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLittleDouble(v float64) {
	WriteDouble(&stream.buffer, v)
}

func (stream *BinaryStream) GetLittleDouble() float64 {
	return ReadDouble(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutPosition(x, y, z int) {
	WritePosition(&stream.buffer, x, y, z)
}

func (stream *BinaryStream) GetUUID() (UUID) {
	return ReadUUID(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutUUID(uuid UUID) {
	WriteUUID(&stream.buffer, uuid)
}

func (stream *BinaryStream) GetPosition() (int, int, int) {
	return ReadPosition(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutTriad(v uint32) {
	WriteBigEndianTriad(&stream.buffer, v)
}

func (stream *BinaryStream) GetTriad() uint32 {
	return ReadBigEndianTriad(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutLittleTriad(v uint32) {
	WriteLittleEndianTriad(&stream.buffer, v)
}

func (stream *BinaryStream) GetLittleTriad() uint32 {
	return ReadLittleEndianTriad(&stream.buffer, &stream.offset)
}

func (stream *BinaryStream) PutBytes(bytes []byte) {
	for _, byte2 := range bytes {
		stream.PutByte(byte2)
	}
}

func (stream *BinaryStream) ResetStream() {
	stream.offset = 0
	stream.buffer = []byte{}
}