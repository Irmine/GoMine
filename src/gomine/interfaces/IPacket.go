package interfaces

type IPacket interface {
	SetBuffer([]byte)
	GetBuffer() []byte
	GetId() int
	EncodeHeader()
	Encode()
	Decode()
}
