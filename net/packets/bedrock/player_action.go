package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/worlds/blocks"
)

const (
	PlayerStartBreak = iota
	PlayerAbortBreak
	PlayerStopBreak
	PlayerGetUpdatedBlock
	PlayerDropItem
	playerStartSleeping
	PlayerStopSleeping
	PlayerRespawn
	PlayerJump
	PlayerStartSprint
	PlayerStopSprint
	PlayerStartSneak
	PlayerStopSneak
	PlayerDimensionChangeRequest
	PlayerDimensionChangeAck
	PlayerStartGlide
	PlayerStopGlide
	PlayerBuildDenied
	PlayerContinueBreak

	//TODO: add rest
)

type PlayerActionPacket struct {
	*packets.Packet
	RuntimeId uint64
	Action    int32
	Position  blocks.Position
	Face      int32
}

func NewPlayerActionPacket() *PlayerActionPacket {
	return &PlayerActionPacket{ packets.NewPacket(info.PacketIds[info.PlayerActionPacket]), 0, 0, blocks.Position{}, 0}
}

func (pk *PlayerActionPacket) Encode() {
	pk.PutEntityRuntimeId(pk.RuntimeId)
	pk.PutVarInt(pk.Action)
	pk.PutBlockPosition(pk.Position)
	pk.PutVarInt(pk.Face)
}

func (pk *PlayerActionPacket) Decode() {
	pk.RuntimeId = pk.GetEntityRuntimeId()
	pk.Action = pk.GetVarInt()
	pk.Position = pk.GetBlockPosition()
	pk.Face = pk.GetVarInt()
}