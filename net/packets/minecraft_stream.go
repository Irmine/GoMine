package packets

import (
	"fmt"
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/binutils"
	"github.com/irmine/gomine/items"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gonbt"
	"github.com/irmine/worlds/blocks"
	"github.com/irmine/worlds/entities/data"
)

// MinecraftStream extends the binutils stream,
// and implements methods for writing types specific
// to the Minecraft bedrock.
type MinecraftStream struct {
	// MinecraftStream embeds binutils.Stream.
	// Usual binary encoding/decoding functions can
	// be called on a MinecraftStream.
	*binutils.Stream
}

// NewMinecraftStream reads a new MinecraftStream.
// This stream is pre-initialized and ready for usage.
func NewMinecraftStream() *MinecraftStream {
	return &MinecraftStream{binutils.NewStream()}
}

// PutEntityRuntimeId writes the runtime ID of an entity.
// Entity runtime IDs are an uint64.
func (stream *MinecraftStream) PutEntityRuntimeId(id uint64) {
	stream.PutUnsignedVarLong(id)
}

// GetEntityRuntimeId reads the runtime ID of an entity.
// Entity runtime IDs are an uint64, and can be looked up
// in the level they belong to.
func (stream *MinecraftStream) GetEntityRuntimeId() uint64 {
	return stream.GetUnsignedVarLong()
}

// PutEntityUniqueId writes the unique ID of an entity.
// Entity unique IDs are an int64, and remain the same through sessions.
func (stream *MinecraftStream) PutEntityUniqueId(id int64) {
	stream.PutVarLong(id)
}

// GetEntityUniqueId reads the unique ID of an entity.
// Unique IDs will currently always be identical to runtime IDs,
// and will therefore have the same result.
func (stream *MinecraftStream) GetEntityUniqueId() int64 {
	return stream.GetVarLong()
}

// PutVector writes a float64 r3.Vector.
// Vector values are first converted to a float32,
// after which they are written little endian.
func (stream *MinecraftStream) PutVector(vector r3.Vector) {
	stream.PutLittleFloat(float32(vector.X))
	stream.PutLittleFloat(float32(vector.Y))
	stream.PutLittleFloat(float32(vector.Z))
}

// GetVector reads a float64 r3.Vector.
// Values read are actually float32, but converted to float64.
func (stream *MinecraftStream) GetVector() r3.Vector {
	return r3.Vector{X: float64(stream.GetLittleFloat()), Y: float64(stream.GetLittleFloat()), Z: float64(stream.GetLittleFloat())}
}

// PutBlockPosition writes a position of a block.
// Block positions are always rounded numbers,
// and the Y value is always positive.
func (stream *MinecraftStream) PutBlockPosition(position blocks.Position) {
	stream.PutVarInt(position.X)
	stream.PutUnsignedVarInt(position.Y)
	stream.PutVarInt(position.Z)
}

// GetBlockPosition reads a position of a block.
// Block positions are always rounded numbers,
// and the Y value is always positive.
func (stream *MinecraftStream) GetBlockPosition() blocks.Position {
	return blocks.NewPosition(stream.GetVarInt(), stream.GetUnsignedVarInt(), stream.GetVarInt())
}

// PutEntityRotation writes the rotation of an entity in bytes.
// The rotation of an entity will only contain yaw and pitch.
func (stream *MinecraftStream) PutEntityRotationBytes(rotation data.Rotation) {
	stream.PutRotationByte(byte(rotation.Pitch))
	stream.PutRotationByte(byte(rotation.Yaw))
	stream.PutRotationByte(byte(rotation.Yaw))
}

// GetEntityRotation reads the rotation of an entity in bytes.
// The rotation of an entity has no different head yaw,
// which will therefore always be the same as the yaw when returned.
func (stream *MinecraftStream) GetEntityRotationBytes() data.Rotation {
	return data.Rotation{Pitch: float64(stream.getRotationByte()), Yaw: float64(stream.getRotationByte()), HeadYaw: float64(stream.getRotationByte())}
}

// PutEntityRotation writes the rotation of an entity.
// The rotation of an entity will only contain yaw and pitch.
func (stream *MinecraftStream) PutEntityRotation(rotation data.Rotation) {
	stream.PutLittleFloat(float32(rotation.Pitch))
	stream.PutLittleFloat(float32(rotation.Yaw))
	stream.PutLittleFloat(float32(rotation.Yaw))
}

// GetEntityRotation reads the rotation of an entity.
// The rotation of an entity has no different head yaw,
// which will therefore always be the same as the yaw when returned.
func (stream *MinecraftStream) GetEntityRotation() data.Rotation {
	return data.Rotation{Pitch: float64(stream.GetLittleFloat()), Yaw: float64(stream.GetLittleFloat()), HeadYaw: float64(stream.GetLittleFloat())}
}

// PutPlayerRotation writes the rotation of a player.
// Players have a head yaw too, which gets written.
func (stream *MinecraftStream) PutPlayerRotation(rot data.Rotation) {
	stream.PutLittleFloat(float32(rot.Pitch))
	stream.PutLittleFloat(float32(rot.Yaw))
	stream.PutLittleFloat(float32(rot.Yaw))
}

// GetPlayerRotation reads the rotation of a player.
// Players are supposed to have a different head yaw than normal yaw,
// but since recent updates the head yaw and yaw are always the same.
func (stream *MinecraftStream) GetPlayerRotation() data.Rotation {
	return data.Rotation{Pitch: float64(stream.GetLittleFloat()), Yaw: float64(stream.GetLittleFloat()), HeadYaw: float64(stream.GetLittleFloat())}
}

func (stream *MinecraftStream) PutRotationByte(rot byte){
	stream.PutByte(rot / (360 / 256))
}

func (stream *MinecraftStream) getRotationByte() byte {
	return stream.GetByte() * (360 / 256)
}

// PutAttributeMap writes the attribute map of an entity.
// The amount of attributes of the map is written,
// after which the attribute properties follow.
func (stream *MinecraftStream) PutAttributeMap(m data.AttributeMap) {
	stream.PutUnsignedVarInt(uint32(len(m)))
	for _, v := range m {
		stream.PutLittleFloat(v.MinValue)
		stream.PutLittleFloat(v.MaxValue)
		stream.PutLittleFloat(v.Value)
		stream.PutLittleFloat(v.DefaultValue)
		stream.PutString(string(v.GetName()))
	}
}

// GetAttributeMap reads an attribute map of an entity.
// There may be attributes in this attribute map that are
// not set in the default attribute map, or missing attributes.
func (stream *MinecraftStream) GetAttributeMap() data.AttributeMap {
	m := data.NewAttributeMap()
	c := stream.GetUnsignedVarInt()
	for i := uint32(0); i < c; i++ {
		min := stream.GetLittleFloat()
		max := stream.GetLittleFloat()
		value := stream.GetLittleFloat()
		defaultValue := stream.GetLittleFloat()
		name := data.AttributeName(stream.GetString())
		att := data.NewAttribute(name, value, max)
		att.DefaultValue = defaultValue
		att.MinValue = min
		m.SetAttribute(att)
	}
	return m
}

// PutItem writes an item stack.
// Item stacks also get their NBT written to network,
// through the call of Stack.EmitNBT().
func (stream *MinecraftStream) PutItem(item *items.Stack) {
	id, v := items.FromKey(items.TypeToId[fmt.Sprint(item.Type)])
	stream.PutVarInt(int32(id))
	stream.PutVarInt(item.GetAuxValue(item, v))

	writer := gonbt.NewWriter(true, binutils.LittleEndian)
	compound := gonbt.NewCompound("", make(map[string]gonbt.INamedTag))
	item.NBTEmitFunction(compound, item)
	writer.WriteUncompressedCompound(compound)
	d := writer.GetBuffer()
	stream.PutLittleShort(int16(len(d)))
	stream.PutBytes(d)

	// Fields for canPlaceOn and canBreak are not implemented.
	// TODO
	stream.PutVarInt(0)
	stream.PutVarInt(0)
}

// GetItem reads a new item stack.
// The item stack returned may have NBT properties.
// If the item ID was unknown, an air item gets returned.
func (stream *MinecraftStream) GetItem() *items.Stack {
	id := stream.GetVarInt()
	if id == 0 {
		i, _ := items.DefaultManager.Get("minecraft:air", 0)
		return i
	}
	aux := stream.GetVarInt()
	itemData := aux >> 8

	t := items.IdToType[items.GetKey(int16(id), int16(itemData))]
	count := aux & 0xff

	item, _ := items.DefaultManager.Get(t.GetId(), int(count))
	nbtData := stream.Get(int(stream.GetLittleShort()))
	reader := gonbt.NewReader(nbtData, true, binutils.LittleEndian)
	item.NBTParseFunction(reader.ReadUncompressedIntoCompound(), item)

	// Fields for canPlaceOn and canBreak are not implemented.
	// TODO
	stream.GetVarInt()
	stream.GetVarInt()

	return item
}

// PutEntityData writes the data properties of an entity.
// TODO: Make a proper implementation.
func (stream *MinecraftStream) PutEntityData(entityData map[uint32][]interface{}) {
	var count= uint32(len(entityData))
	stream.PutUnsignedVarInt(count)
	for key, dataValues := range entityData {
		stream.PutUnsignedVarInt(key)
		var flagId, ok = dataValues[0].(uint32)
		if !ok {
			stream.PutUnsignedVarInt(999999) // invalid flag id
			continue
		}

		stream.PutUnsignedVarInt(flagId)

		switch flagId {
		case data.EntityDataByte:
			if value, ok := dataValues[1].(byte); ok {
				stream.PutByte(value)
			}
			break
		case data.EntityDataShort:
			if value, ok := dataValues[1].(int16); ok {
				stream.PutLittleShort(value)
			}
			break
		case data.EntityDataInt:
			if value, ok := dataValues[1].(int32); ok {
				stream.PutVarInt(value)
			}
			break
		case data.EntityDataFloat:
			if value, ok := dataValues[1].(float32); ok {
				stream.PutLittleFloat(value)
			}
			break
		case data.EntityDataString:
			if value, ok := dataValues[1].(string); ok {
				stream.PutString(value)
			}
			break
		case data.EntityDataItem:
			if value, ok := dataValues[1].(*items.Stack); ok {
				stream.PutItem(value)
			}
			break
		case data.EntityDataPos:
			if value, ok := dataValues[1].(blocks.Position); ok {
				stream.PutBlockPosition(value)
			}
			break
		case data.EntityDataLong:
			if value, ok := dataValues[1].(int64); ok {
				stream.PutVarLong(value)
			}
			break
		case data.EntityDataVector:
			if value, ok := dataValues[1].(r3.Vector); ok {
				stream.PutVector(value)
			}
			break
		}
	}
}

// GetEntityData reads an entity data property map from an entity.
func (stream *MinecraftStream) GetEntityData() map[uint32][]interface{} {
	entityData := make(map[uint32][]interface{})
	count := stream.GetUnsignedVarInt()
	if count > 0 {
		for i := uint32(0); i < count; i++ {
			var key = stream.GetUnsignedVarInt()
			var flagId = stream.GetUnsignedVarInt()
			switch flagId {
			case data.EntityDataByte:
				entityData[key] = []interface{}{flagId, stream.GetByte()}
				break
			case data.EntityDataShort:
				entityData[key] = []interface{}{flagId, stream.GetLittleShort()}
				break
			case data.EntityDataInt:
				entityData[key] = []interface{}{flagId, stream.GetVarInt()}
				break
			case data.EntityDataFloat:
				entityData[key] = []interface{}{flagId, stream.GetLittleFloat()}
				break
			case data.EntityDataString:
				entityData[key] = []interface{}{flagId, stream.GetString()}
				break
			case data.EntityDataItem:
				entityData[key] = []interface{}{flagId, stream.GetItem()}
				break
			case data.EntityDataPos:
				entityData[key] = []interface{}{flagId, stream.GetBlockPosition()}
				break
			case data.EntityDataLong:
				entityData[key] = []interface{}{flagId, stream.GetVarLong()}
				break
			case data.EntityDataVector:
				entityData[key] = []interface{}{flagId, stream.GetVector()}
				break
			}
		}
	}
	return entityData
}

// PutGameRules writes a map of game rules.
// Game rules get prefixed by the type of the game rule,
// 1 being bool, 2 being uint32, 3 being float32.
func (stream *MinecraftStream) PutGameRules(gameRules map[string]types.GameRuleEntry) {
	stream.PutUnsignedVarInt(uint32(len(gameRules)))
	for _, gameRule := range gameRules {
		stream.PutString(gameRule.Name)
		switch value := gameRule.Value.(type) {
		case bool:
			stream.PutByte(1)
			stream.PutBool(value)
		case uint32:
			stream.PutByte(2)
			stream.PutUnsignedVarInt(value)
		case float32:
			stream.PutByte(3)
			stream.PutLittleFloat(value)
		}
	}
}

// PutPackInfo writes the info of an array of resource pack entries.
// The UUID, version and pack size gets written.
func (stream *MinecraftStream) PutPackInfo(packs []types.ResourcePackInfoEntry) {
	stream.PutLittleShort(int16(len(packs)))

	for _, pack := range packs {
		stream.PutString(pack.UUID)
		stream.PutString(pack.Version)
		stream.PutLittleLong(pack.PackSize)
		stream.PutString("")
		stream.PutString("")
	}
}

// PutPackStack writes an array of resource pack entries.
// The order of this array specifies the order the client should apply those,
// with index 0 meaning highest priority.
func (stream *MinecraftStream) PutPackStack(packs []types.ResourcePackStackEntry) {
	stream.PutUnsignedVarInt(uint32(len(packs)))
	for _, pack := range packs {
		stream.PutString(pack.UUID)
		stream.PutString(pack.Version)
		stream.PutString("")
	}
}

// PutUUID writes a UUID.
// UUIDs are first re-ordered for little endian byte order,
// after which they get written.
func (stream *MinecraftStream) PutUUID(uuid uuid.UUID) {
	b, err := uuid.MarshalBinary()
	if err != nil {
		panic(err)
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	stream.PutBytes(b[8:])
	stream.PutBytes(b[:8])
}

// GetUUID reads a UUID.
// TODO: Re-order for little endian byte order. Order gets messed up.
func (stream *MinecraftStream) GetUUID() uuid.UUID {
	return uuid.Must(uuid.FromBytes(stream.Get(16)))
}
