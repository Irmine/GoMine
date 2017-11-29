package packets

import (
	"gomine/utils"
	"gomine/vectorMath"
	"gomine/entities"
)

type Packet struct {
	*utils.BinaryStream
	PacketId int
	ExtraBytes [2]byte
}

func NewPacket(id int) *Packet {
	return &Packet{utils.NewStream(), id, [2]byte{}}
}

func (pk *Packet) GetId() int {
	return pk.PacketId
}

func (pk *Packet) Encode() {

}

func (pk *Packet) Decode() {

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
	if pid != pk.PacketId {
		panic("Packet IDs do not match")
	}

	pk.ExtraBytes[0] = pk.GetByte()
	pk.ExtraBytes[1] = pk.GetByte()

	if pk.ExtraBytes[0] != 0 && pk.ExtraBytes[1] != 0 {
		panic("extra bytes are not zero")
	}
}

func (pk *Packet) PutRuntimeId(eid uint64) {
	pk.PutUnsignedVarLong(eid)
}

func (pk *Packet) GetRuntimeId() uint64 {
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

func (pk *Packet) PutEntityData(dat map[uint32][]interface{}) {
	pk.PutUnsignedVarInt(uint32(len(dat)))
	for k, v := range dat {
		pk.PutUnsignedVarInt(k)
		pk.PutUnsignedVarInt(v[0].(uint32))
		switch v[1] {
		case entities.Byte:
			pk.PutByte(v[1].(byte))
		case entities.Short:
			pk.PutLittleShort(v[1].(int16))
		case entities.Int:
			pk.PutVarInt(v[1].(int32))
		case entities.Float:
			pk.PutLittleFloat(v[1].(float32))
		case entities.String:
			pk.PutString(v[1].(string))
		case entities.Slot:
			//todo
		case entities.Pos:
			//todo
		case entities.Long:
			pk.PutVarLong(v[1].(int64))
		case entities.Vector3f:
			//todo
		}
	}
}

func (pk *Packet) GetEntityData() map[uint32][]interface{} {
	var dat = make(map[uint32][]interface{})
	len2 := pk.GetUnsignedVarInt()
	for i := uint32(0); i < len2; i++ {
		k := pk.GetUnsignedVarInt()
		t := pk.GetUnsignedVarInt()
		var v interface{}
		switch t {
		case entities.Byte:
			v = pk.GetByte()
		case entities.Short:
			v = pk.GetLittleShort()
		case entities.Int:
			v = pk.GetVarInt()
		case entities.Float:
			v = pk.GetLittleFloat()
		case entities.String:
			v = pk.GetString()
		case entities.Slot:
			//todo
		case entities.Pos:
			//todo
		case entities.Long:
			v = pk.GetVarLong()
		case entities.Vector3f:
			//todo
		}
		dat[k][0] = t
		dat[k][1] = v
	}
	return dat
}

func (pk *Packet) PutBlockPos(x int32, y uint32, z int32) {
	pk.PutVarInt(x)
	pk.PutUnsignedVarInt(y)
	pk.PutVarInt(x)
}