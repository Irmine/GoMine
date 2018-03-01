package packets

import (
	"github.com/golang/geo/r3"
	"github.com/irmine/binutils"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/utils"
	"github.com/irmine/worlds/entities/data"
)

type IPacket interface {
	SetBuffer([]byte)
	GetBuffer() []byte
	GetId() int
	SetId(int)
	EncodeHeader()
	Encode()
	DecodeHeader()
	Decode()
	ResetStream()
	GetOffset() int
	SetOffset(int)
	Discard()
	IsDiscarded() bool
	EncodeId()
	DecodeId()
}

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

func (pk *Packet) PutVector(obj r3.Vector) {
	pk.PutLittleFloat(float32(obj.X))
	pk.PutLittleFloat(float32(obj.Y))
	pk.PutLittleFloat(float32(obj.Z))
}

func (pk *Packet) GetVector() r3.Vector {
	return r3.Vector{X: float64(pk.GetLittleFloat()), Y: float64(pk.GetLittleFloat()), Z: float64(pk.GetLittleFloat())}
}

func (pk *Packet) PutRotation(rot data.Rotation, isPlayer bool) {
	pk.PutLittleFloat(float32(rot.Pitch))
	pk.PutLittleFloat(float32(rot.Yaw))
	if isPlayer {
		pk.PutLittleFloat(float32(rot.HeadYaw))
	}
}

func (pk *Packet) GetRotation(isPlayer bool) data.Rotation {
	var yaw = float64(pk.GetLittleFloat())
	var pitch = float64(pk.GetLittleFloat())
	var headYaw float64 = 0
	if isPlayer {
		headYaw = float64(pk.GetLittleFloat())
	}
	return data.Rotation{yaw, pitch, headYaw}
}

func (pk *Packet) PutAttributeMap(attMap data.AttributeMap) {
	pk.PutUnsignedVarInt(uint32(len(attMap)))
	for _, v := range attMap {
		pk.PutLittleFloat(v.GetMinValue())
		pk.PutLittleFloat(v.GetMaxValue())
		pk.PutLittleFloat(v.GetValue())
		pk.PutLittleFloat(v.GetDefaultValue())
		pk.PutString(string(v.GetName()))
	}
}

func (pk *Packet) GetAttributeMap() data.AttributeMap {
	attributes := data.NewAttributeMap()
	c := pk.GetUnsignedVarInt()

	for i := uint32(0); i < c; i++ {
		pk.GetLittleFloat()
		max := pk.GetLittleFloat()
		value := pk.GetLittleFloat()
		pk.GetLittleFloat()
		name := data.AttributeName(pk.GetString())

		if attributes.Exists(name) {
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
		case 0:
			pk.PutByte(v[1].(byte))
		case 1:
			pk.PutLittleShort(v[1].(int16))
		case 2:
			pk.PutVarInt(v[1].(int32))
		case 3:
			pk.PutLittleFloat(v[1].(float32))
		case 4:
			pk.PutString(v[1].(string))
		case 5:
			//todo
		case 6:
			//todo
		case 7:
			pk.PutVarLong(v[1].(int64))
		case 8:
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
		case 0:
			v = pk.GetByte()
		case 1:
			v = pk.GetLittleShort()
		case 2:
			v = pk.GetVarInt()
		case 3:
			v = pk.GetLittleFloat()
		case 4:
			v = pk.GetString()
		case 5:
			//todo
		case 6:
			//todo
		case 7:
			v = pk.GetVarLong()
		case 8:
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

func (pk *Packet) PutBlockPos(vector r3.Vector) {
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
