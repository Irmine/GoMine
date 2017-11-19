package packets

import (
	"gomine/utils"
	"gomine/vectorMath"
	"gomine/entities"
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

func (pk *Packet) GetId() int {
	return pk.packetId
}

func (pk *Packet) Encode() {

}

func (pk *Packet) SkipId() {
	pk.Offset++
}

func (pk *Packet) SkipSplitBytes() {
	pk.Offset += 2
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

func (pk *Packet) PutEId(eid uint64) {
	pk.PutUnsignedVarLong(eid)
}

func (pk *Packet) GetEId() uint64 {
	return pk.GetUnsignedVarLong()
}

func (pk *Packet) PutRotation(rot float32) {
	pk.PutLittleFloat(rot)
}

func (pk *Packet) GetRotation() float32 {
	return pk.GetLittleFloat()
}

func (pk *Packet) PutTripleVectorObject(obj vectorMath.TripleVector) {
	pk.PutLittleFloat(obj.GetX())
	pk.PutLittleFloat(obj.GetY())
	pk.PutLittleFloat(obj.GetZ())
}

func (pk *Packet) GetTripleVectorObject() *vectorMath.TripleVector {
	return vectorMath.NewTripleVector(pk.GetLittleFloat(), pk.GetLittleFloat(), pk.GetLittleFloat())
}

func (pk *Packet) PutEntityAttributes(attr map[int]entities.Attribute) {
	for _, v := range attr {
		pk.PutLittleFloat(v.GetMinValue())
		pk.PutLittleFloat(v.GetMaxValue())
		pk.PutLittleFloat(v.GetValue())
		pk.PutLittleFloat(v.GetDefaultValue())
		pk.PutString(v.GetName())
	}
}

func (pk *Packet) GetEntityAttributes() map[int]entities.Attribute {
	//todo
	return map[int]entities.Attribute{}
}