package gonbt

type BinaryStream struct {
	Offset int
	Buffer []byte
	Network bool
	EndianType byte
}

func NewStream(buffer []byte, network bool, endian byte) *BinaryStream {
	return &BinaryStream{0, buffer, network, endian}
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

func (stream *BinaryStream) PutShort(v int16) {
	WriteShort(&stream.Buffer, v, stream.EndianType)
}

func (stream *BinaryStream) GetShort() int16 {
	return ReadShort(&stream.Buffer, &stream.Offset, stream.EndianType)
}

func (stream *BinaryStream) PutInt(v int32) {
	WriteInt(&stream.Buffer, v, stream.Network, stream.EndianType)
}

func (stream *BinaryStream) GetInt() int32 {
	return ReadInt(&stream.Buffer, &stream.Offset, stream.Network, stream.EndianType)
}

func (stream *BinaryStream) PutLong(v int64) {
	WriteLong(&stream.Buffer, v, stream.Network, stream.EndianType)
}

func (stream *BinaryStream) GetLong() int64 {
	return ReadLong(&stream.Buffer, &stream.Offset, stream.Network, stream.EndianType)
}

func (stream *BinaryStream) PutFloat(v float32) {
	WriteFloat(&stream.Buffer, v, stream.EndianType)
}

func (stream *BinaryStream) GetFloat() float32 {
	return ReadFloat(&stream.Buffer, &stream.Offset, stream.EndianType)
}

func (stream *BinaryStream) PutDouble(v float64) {
	WriteDouble(&stream.Buffer, v, stream.EndianType)
}

func (stream *BinaryStream) GetDouble() float64 {
	return ReadDouble(&stream.Buffer, &stream.Offset, stream.EndianType)
}

func (stream *BinaryStream) PutString(v string) {
	WriteString(&stream.Buffer, v, stream.Network, stream.EndianType)
}

func (stream *BinaryStream) GetString() string {
	return ReadString(&stream.Buffer, &stream.Offset, stream.Network, stream.EndianType)
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
