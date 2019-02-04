package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
	"github.com/irmine/worlds/blocks"
)

type NetworkChunkPublisherUpdatePacket struct {
	*packets.Packet
	Position blocks.Position
	Radius uint32
}

func NewNetworkChunkPublisherUpdatePacket() *NetworkChunkPublisherUpdatePacket {
	return &NetworkChunkPublisherUpdatePacket{packets.NewPacket(info.PacketIds[info.NetworkChunkPublisherUpdatePacket]), blocks.NewPosition(0, 0, 0), 0}
}

func (pk *NetworkChunkPublisherUpdatePacket) Encode() {
	pk.PutBlockPosition(pk.Position)
	pk.PutUnsignedVarInt(pk.Radius)
}

func (pk *NetworkChunkPublisherUpdatePacket) Decode() {
}
