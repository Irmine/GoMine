package packets

import (
	"github.com/irmine/gomine/entities/data"
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/gomine/vectors"
	"github.com/irmine/binutils"
)

type Packet struct {
	*binutils.Stream
	PacketId           int
	SenderIdentifier   byte
	ReceiverIdentifier byte
	discarded          bool
}

func NewPacket(id int) *Packet {
	return &Packet{binutils.NewStream(), id, 0, 0, false}
}

func (pk *Packet) GetId() int {
	return pk.PacketId
}

func (pk *Packet) SetId(id int) {
	pk.PacketId = id
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

func (pk *Packet) Discard() {
	pk.discarded = true
}

func (pk *Packet) IsDiscarded() bool {
	return pk.discarded
}

func (pk *Packet) EncodeHeader() {
	pk.ResetStream()
	pk.EncodeId()
	pk.PutByte(pk.SenderIdentifier)
	pk.PutByte(pk.ReceiverIdentifier)
}

func (pk *Packet) EncodeId() {
	pk.PutUnsignedVarInt(uint32(pk.GetId()))
}

func (pk *Packet) DecodeId() {
	var pid = int(pk.GetUnsignedVarInt())
	if pid != pk.PacketId {
		panic("Packet IDs do not match")
	}
}

func (pk *Packet) DecodeHeader() {
	pk.DecodeId()
	pk.SenderIdentifier = pk.GetByte()
	pk.ReceiverIdentifier = pk.GetByte()
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

func (pk *Packet) PutEntityAttributeMap(attr *data.AttributeMap) {
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

func (pk *Packet) GetEntityAttributeMap() *data.AttributeMap {
	attributes := data.NewAttributeMap()
	c := pk.GetUnsignedVarInt()

	for i := uint32(0); i < c; i++ {
		pk.GetLittleFloat()
		max := pk.GetLittleFloat()
		value := pk.GetLittleFloat()
		pk.GetLittleFloat()
		name := pk.GetString()

		if data.AttributeExists(name) {
			attributes.SetAttribute(data.NewAttribute(name, value, max))
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
		case data.Byte:
			pk.PutByte(v[1].(byte))
		case data.Short:
			pk.PutLittleShort(v[1].(int16))
		case data.Int:
			pk.PutVarInt(v[1].(int32))
		case data.Float:
			pk.PutLittleFloat(v[1].(float32))
		case data.String:
			pk.PutString(v[1].(string))
		case data.Slot:
			//todo
		case data.Pos:
			//todo
		case data.Long:
			pk.PutVarLong(v[1].(int64))
		case data.TripleFloat:
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
		case data.Byte:
			v = pk.GetByte()
		case data.Short:
			v = pk.GetLittleShort()
		case data.Int:
			v = pk.GetVarInt()
		case data.Float:
			v = pk.GetLittleFloat()
		case data.String:
			v = pk.GetString()
		case data.Slot:
			//todo
		case data.Pos:
			//todo
		case data.Long:
			v = pk.GetVarLong()
		case data.TripleFloat:
			//todo
		}
		dat[k][0] = t
		dat[k][1] = v
	}
	return dat
}

func (pk *Packet) PutGameRules(gameRules map[string]types.GameRuleEntry) {
	pk.PutUnsignedVarInt(uint32(len(gameRules)))
	for _, gameRule := range gameRules {
		pk.PutString(gameRule.Name)
		switch value := gameRule.Value.(type) {
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

func (pk *Packet) PutPackInfo(packs []types.ResourcePackInfoEntry) {
	pk.PutLittleShort(int16(len(packs)))

	for _, pack := range packs {
		pk.PutString(pack.UUID)
		pk.PutString(pack.Version)
		pk.PutLittleLong(pack.PackSize)
		pk.PutString("")
		pk.PutString("")
	}
}

func (pk *Packet) PutPackStack(packs []types.ResourcePackStackEntry) {
	pk.PutUnsignedVarInt(uint32(len(packs)))
	for _, pack := range packs {
		pk.PutString(pack.UUID)
		pk.PutString(pack.Version)
		pk.PutString("")
	}
}

func (pk *Packet) PutUUID(uuid utils.UUID) {
	pk.PutLittleInt(uuid.GetParts()[1])
	pk.PutLittleInt(uuid.GetParts()[0])
	pk.PutLittleInt(uuid.GetParts()[3])
	pk.PutLittleInt(uuid.GetParts()[2])
}

func (pk *Packet) GetUUID() utils.UUID {
	var unorderedParts = [4]int32{pk.GetLittleInt(), pk.GetLittleInt(), pk.GetLittleInt(), pk.GetLittleInt()}
	var parts = [4]int32{unorderedParts[1], unorderedParts[0], unorderedParts[3], unorderedParts[2]}
	return utils.NewUUID(parts)
}
