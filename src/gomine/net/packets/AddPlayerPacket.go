package packets

import (
	"gomine/utils"
	"gomine/vectors"
	"gomine/entities/math"
	"gomine/net/info"
)

type AddPlayerPacket struct {
	*Packet
	UUID utils.UUID
	Username string
	DisplayName string
	Platform int32
	UnknownString string

	EntityUniqueId int64
	EntityRuntimeId uint64
	Position vectors.TripleVector
	Motion vectors.TripleVector
	Rotation math.Rotation
	// HandItem TODO: Items.
	Metadata map[uint32][]interface{}

	Flags uint32
	CommandPermission uint32
	Flags2 uint32
	PlayerPermission uint32
	CustomFlags uint32

	Long1 int64
	// EntityLinks TODO
}

func NewAddPlayerPacket() *AddPlayerPacket {
	return &AddPlayerPacket{Packet: NewPacket(info.AddPlayerPacket), Metadata: make(map[uint32][]interface{}), Motion: vectors.TripleVector{}}
}

func (pk *AddPlayerPacket) Encode()  {
	pk.PutUUID(pk.UUID)
	pk.PutString(pk.Username)
	pk.PutString(pk.DisplayName)
	pk.PutVarInt(pk.Platform)

	pk.PutUniqueId(pk.EntityUniqueId)
	pk.PutRuntimeId(pk.EntityRuntimeId)
	pk.PutString(pk.UnknownString)

	pk.PutTripleVectorObject(pk.Position)
	pk.PutTripleVectorObject(pk.Motion)
	pk.PutRotationObject(pk.Rotation, true)

	pk.PutVarInt(0) // TODO
	pk.PutEntityData(pk.Metadata)

	pk.PutUnsignedVarInt(pk.Flags)
	pk.PutUnsignedVarInt(pk.CommandPermission)
	pk.PutUnsignedVarInt(pk.Flags2)
	pk.PutUnsignedVarInt(pk.PlayerPermission)
	pk.PutUnsignedVarInt(pk.CustomFlags)

	pk.PutVarLong(pk.Long1)

	pk.PutUnsignedVarInt(0) // TODO
}

func (pk *AddPlayerPacket) Decode()  {

}
