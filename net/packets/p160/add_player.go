package p160

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
}

func NewAddPlayerPacket() *AddPlayerPacket {
	return &AddPlayerPacket{Packet: packets.NewPacket(info.PacketIds200[info.AddPlayerPacket]), Metadata: make(map[uint32][]interface{}), Motion: r3.Vector{}}
}

func (pk *AddPlayerPacket) Encode() {
	pk.PutUUID(pk.UUID)
	pk.PutString(pk.Username)

	pk.PutUniqueId(pk.EntityUniqueId)
	pk.PutRuntimeId(pk.EntityRuntimeId)

	pk.PutVector(pk.Position)
	pk.PutVector(pk.Motion)
	pk.PutRotation(pk.Rotation, true)

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

func (pk *AddPlayerPacket) Decode() {

}
