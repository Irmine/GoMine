package mcpe

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/worlds/entities/data"
)

type AddPlayerPacket struct {
	*packets.Packet
	UUID            uuid.UUID
	Username        string
	DisplayName     string
	Platform        int32
	UnknownString   string
	EntityUniqueId  int64
	EntityRuntimeId uint64
	Position        r3.Vector
	Motion          r3.Vector
	Rotation        data.Rotation
	// HandItem TODO: Items.
	Metadata          map[uint32][]interface{}
	Flags             uint32
	CommandPermission uint32
	Flags2            uint32
	PlayerPermission  uint32
	CustomFlags       uint32
	Long1             int64
	// EntityLinks TODO
	DeviceID          string
}

func NewAddPlayerPacket() *AddPlayerPacket {
	return &AddPlayerPacket{Packet: packets.NewPacket(info.PacketIds[info.AddPlayerPacket]), Metadata: make(map[uint32][]interface{}), Motion: r3.Vector{}}
}

func (pk *AddPlayerPacket) Encode() {
	pk.PutUUID(pk.UUID)
	pk.PutString(pk.Username)
	pk.PutString(pk.DisplayName)
	pk.PutVarInt(pk.Platform)

	pk.PutEntityUniqueId(pk.EntityUniqueId)
	pk.PutEntityRuntimeId(pk.EntityRuntimeId)
	pk.PutString(pk.UnknownString)

	pk.PutVector(pk.Position)
	pk.PutVector(pk.Motion)
	pk.PutPlayerRotation(pk.Rotation)

	pk.PutVarInt(0) // TODO
	pk.PutEntityData(pk.Metadata)

	pk.PutUnsignedVarInt(pk.Flags)
	pk.PutUnsignedVarInt(pk.CommandPermission)
	pk.PutUnsignedVarInt(pk.Flags2)
	pk.PutUnsignedVarInt(pk.PlayerPermission)
	pk.PutUnsignedVarInt(pk.CustomFlags)

	pk.PutVarLong(pk.Long1)

	pk.PutUnsignedVarInt(0) // TODO
	pk.PutString(pk.DeviceID)
}

func (pk *AddPlayerPacket) Decode() {

}
