package packets

import (
	"gomine/utils"
	"gomine/vectors"
	"gomine/entities"
	"gomine/interfaces"
	"gomine/entities/math"
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

func (pk *Packet) PutRuntimeId(id uint64) {
	pk.PutUnsignedVarLong(id)
}

func (pk *Packet) GetRuntimeId() uint64 {
	return pk.GetUnsignedVarLong()
}

func (pk *Packet) PutUniqueId(id int64) {
	pk.PutVarLong(id)
}

func (pk *Packet) GetUniqueId() int64 {
	return pk.GetVarLong()
}

func (pk *Packet) PutTripleVectorObject(obj vectors.TripleVector) {
	pk.PutLittleFloat(obj.GetX())
	pk.PutLittleFloat(obj.GetY())
	pk.PutLittleFloat(obj.GetZ())
}

func (pk *Packet) GetTripleVectorObject() *vectors.TripleVector {
	return &vectors.TripleVector{X: pk.GetLittleFloat(), Y: pk.GetLittleFloat(), Z: pk.GetLittleFloat()}
}

func (pk *Packet) PutRotationObject(obj math.Rotation, isPlayer bool) {
	pk.PutLittleFloat(obj.Pitch)
	pk.PutLittleFloat(obj.Yaw)
	if isPlayer {
		pk.PutLittleFloat(obj.HeadYaw)
	}
}

func (pk *Packet) GetRotationObject(isPlayer bool) math.Rotation {
	var yaw = pk.GetLittleFloat()
	var pitch = pk.GetLittleFloat()
	var headYaw float32 = 0
	if isPlayer {
		headYaw = pk.GetLittleFloat()
	}
	return *math.NewRotation(yaw, pitch, headYaw)
}

func (pk *Packet) PutEntityAttributeMap(attr *entities.AttributeMap) {
	attrList := attr.GetAttributes()
	pk.PutUnsignedVarInt(uint32(len(attrList)))
	for _, v := range attrList {
		pk.PutLittleFloat(v.GetMinValue())
		pk.PutLittleFloat(v.GetMaxValue())
		pk.PutLittleFloat(v.GetValue())
		pk.PutLittleFloat(v.GetDefaultValue())
		pk.PutString(v.GetName())
	}
}

func (pk *Packet) GetEntityAttributeMap() *entities.AttributeMap {
	attributes := entities.NewAttributeMap()
	c := pk.GetUnsignedVarInt()

	for i := uint32(0); i < c; i++ {
		pk.GetLittleFloat()
		max := pk.GetLittleFloat()
		value := pk.GetLittleFloat()
		pk.GetLittleFloat()
		name := pk.GetString()

		if entities.AttributeExists(name) {
			attributes.SetAttribute(entities.NewAttribute(name, value, max))
		}
	}

	return attributes
}

func (pk *Packet) PutEntityData(dat map[uint32][]interface{}) {
	pk.PutUnsignedVarInt(uint32(len(dat)))
	for k, v := range dat {
		pk.PutUnsignedVarInt(k)
		pk.PutUnsignedVarInt(v[0].(uint32))
		switch v[0] {
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
		case entities.TripleFloat:
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
		case entities.TripleFloat:
			//todo
		}
		dat[k][0] = t
		dat[k][1] = v
	}
	return dat
}

func (pk *Packet) PutGameRules(gameRules map[string]interfaces.IGameRule) {
	pk.PutUnsignedVarInt(uint32(len(gameRules)))
	for _, gameRule := range gameRules {
		pk.PutString(gameRule.GetName())
		switch value := gameRule.GetValue().(type) {
		case bool:
			pk.PutByte(1)
			pk.PutBool(value)
		case uint32:
			pk.PutByte(2)
			pk.PutUnsignedVarInt(value)
		case float32:
			pk.PutByte(3)
			pk.PutLittleFloat(value)
		}
	}
}

func (pk *Packet) PutBlockPos(vector vectors.TripleVector) {
	pk.PutVarInt(int32(vector.X))
	pk.PutUnsignedVarInt(uint32(vector.Y))
	pk.PutVarInt(int32(vector.Z))
}

func (pk *Packet) PutPacks(packs []interfaces.IPack, info bool) {
	if info {
		pk.PutLittleShort(int16(len(packs)))

		for _, pack := range packs {
			pk.PutString(pack.GetUUID())
			pk.PutString(pack.GetVersion())
			pk.PutLittleLong(pack.GetFileSize())
			pk.PutString("")
			pk.PutString("")
		}
	} else {
		pk.PutUnsignedVarInt(uint32(len(packs)))
		for _, pack := range packs {
			pk.PutString(pack.GetUUID())
			pk.PutString(pack.GetVersion())
			pk.PutString("")
		}
	}
}

func (pk *Packet) GetUUID() utils.UUID {
	return utils.UUIDFromBinary(&pk.Buffer, &pk.Offset)
}

func (pk *Packet) PutUUID(uuid utils.UUID) {
	pk.PutLittleInt(uuid.GetParts()[1])
	pk.PutLittleInt(uuid.GetParts()[0])
	pk.PutLittleInt(uuid.GetParts()[3])
	pk.PutLittleInt(uuid.GetParts()[2])
}