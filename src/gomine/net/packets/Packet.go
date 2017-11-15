package packets

import (
	"gomine/utils"
)

type Packet struct {
	packetId int
	*utils.BinaryStream
	ExtraBytes [2]byte
}

type IPacket interface {
	SetBuffer([]byte)
	GetBuffer() []byte
	GetId() int
	Encode()
	Decode()
}

func NewPacket(id int) *Packet {
	return &Packet{id, utils.NewStream(), [2]byte{}}
}

func (pk *Packet) SetBuffer(b []byte) {
	pk.SetBuffer(b)
}

func (pk *Packet) GetBuffer() []byte {
	return pk.GetBuffer()
}

func (pk *Packet) GetId() int {
	return pk.packetId
}

func (pk *Packet) Encode() {

}

func (pk *Packet) Decode() {

}

func (pk *Packet) EncodeHeader() {
	pk.ResetStream()
	pk.PutUnsignedVarInt(uint32(pk.GetId()))
	pk.PutByte(pk.ExtraBytes[0])
	pk.PutByte(pk.ExtraBytes[1])
}

func (pk *Packet) DecodeHeader() {
	pid := int(pk.GetUnsignedVarInt())
	if pid != pk.packetId {
		panic("Packet IDs do not match")
	}

	pk.ExtraBytes[0] = pk.GetByte()
	pk.ExtraBytes[1] = pk.GetByte()

	if pk.ExtraBytes[0] != 0 && pk.ExtraBytes[1] != 0 {
		panic("extra bytes are not zero")
	}
}