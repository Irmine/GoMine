package utils

type BinaryStream struct {
	Offset int
	Buffer []byte
}

func NewStream() *BinaryStream {
	return &BinaryStream{0, make([]byte, 4096)}
}

func (stream *BinaryStream) GetOffset() int {
	return stream.Offset
}

func (stream *BinaryStream) SetOffset(offset int) {
	stream.Offset = offset
}

func (stream *BinaryStream) SetBuffer(Buffer []byte) {
	stream.Buffer = Buffer
}

func (stream *BinaryStream) GetBuffer() []byte {
	return stream.Buffer
}

func (stream *BinaryStream) Feof() bool {
	return stream.Offset >= len(stream.Buffer) - 1
}

func (stream *BinaryStream) Get(length int) []byte {
	return Read(&stream.Buffer, &stream.Offset, length)
}

//big

func (stream *BinaryStream) PutBool(v bool) {
	WriteBool(&stream.Buffer, v)
}

func (stream *BinaryStream) GetBool() bool {
	return ReadBool(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutByte(v byte) {
	WriteByte(&stream.Buffer, v)
}

func (stream *BinaryStream) GetByte() byte {
	return ReadByte(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutUnsignedByte(v byte) {
	WriteUnsignedByte(&stream.Buffer, v)
}

func (stream *BinaryStream) GetUnsignedByte() byte {
	return ReadUnsignedByte(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutShort(v int16) {
	WriteShort(&stream.Buffer, v)
}

func (stream *BinaryStream) GetShort() int16 {
	return ReadShort(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutUnsignedShort(v uint16) {
	WriteUnsignedShort(&stream.Buffer, v)
}

func (stream *BinaryStream) GetUnsignedShort() uint16 {
	return ReadUnsignedShort(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutInt(v int32) {
	WriteInt(&stream.Buffer, v)
}

func (stream *BinaryStream) GetInt() int32 {
	return ReadInt(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLong(v int64) {
	WriteLong(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLong() int64 {
	return ReadLong(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutUnsignedLong(v uint64) {
	WriteUnsignedLong(&stream.Buffer, v)
}

func (stream *BinaryStream) GetUnsignedLong() uint64 {
	return ReadUnsignedLong(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutFloat(v float32) {
	WriteFloat(&stream.Buffer, v)
}

func (stream *BinaryStream) GetFloat() float32 {
	return ReadFloat(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutDouble(v float64) {
	WriteDouble(&stream.Buffer, v)
}

func (stream *BinaryStream) GetDouble() float64 {
	return ReadDouble(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutVarInt(v int32) {
	WriteVarInt(&stream.Buffer, v)
}

func (stream *BinaryStream) GetVarInt() int32 {
	return ReadVarInt(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutVarLong(v int64) {
	WriteVarLong(&stream.Buffer, v)
}

func (stream *BinaryStream) GetVarLong() int64 {
	return ReadVarLong(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutUnsignedVarInt(v uint32) {
	WriteUnsignedVarInt(&stream.Buffer, v)
}

func (stream *BinaryStream) GetUnsignedVarInt() uint32 {
	return ReadUnsignedVarInt(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutUnsignedVarLong(v uint64) {
	WriteUnsignedVarLong(&stream.Buffer, v)
}

func (stream *BinaryStream) GetUnsignedVarLong() uint64 {
	return ReadUnsignedVarLong(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutString(v string) {
	WriteString(&stream.Buffer, v)
}

func (stream *BinaryStream) GetString() string {
	return ReadString(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLittleShort(v int16) {
	WriteLittleShort(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLittleShort() int16 {
	return ReadLittleShort(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLittleUnsignedShort(v uint16) {
	WriteLittleUnsignedShort(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLittleUnsignedShort() uint16 {
	return ReadLittleUnsignedShort(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLittleInt(v int32) {
	WriteLittleInt(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLittleInt() int32 {
	return ReadLittleInt(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLittleLong(v int64) {
	WriteLittleLong(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLittleLong() int64 {
	return ReadLittleLong(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLittleUnsignedLong(v uint64) {
	WriteLittleUnsignedLong(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLittleUnsignedLong() uint64 {
	return ReadLittleUnsignedLong(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLittleFloat(v float32) {
	WriteLittleFloat(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLittleFloat() float32 {
	return ReadLittleFloat(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLittleDouble(v float64) {
	WriteDouble(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLittleDouble() float64 {
	return ReadDouble(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutPosition(x, y, z int) {
	WritePosition(&stream.Buffer, x, y, z)
}

func (stream *BinaryStream) GetPosition() (int, int, int) {
	return ReadPosition(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutTriad(v uint32) {
	WriteBigEndianTriad(&stream.Buffer, v)
}

func (stream *BinaryStream) GetTriad() uint32 {
	return ReadBigEndianTriad(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutLittleTriad(v uint32) {
	WriteLittleEndianTriad(&stream.Buffer, v)
}

func (stream *BinaryStream) GetLittleTriad() uint32 {
	return ReadLittleEndianTriad(&stream.Buffer, &stream.Offset)
}

func (stream *BinaryStream) PutBytes(bytes []byte) {
	for _, byte2 := range bytes {
		stream.PutByte(byte2)
	}
}

func (stream *BinaryStream) ResetStream() {
	stream.Offset = 0
	stream.Buffer = []byte{}
}