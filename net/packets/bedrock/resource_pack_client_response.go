package bedrock

import (
	"github.com/irmine/gomine/net/info"
	"github.com/irmine/gomine/net/packets"
)

type ResourcePackClientResponsePacket struct {
	*packets.Packet
	Status    byte
	PackUUIDs []string
}

func NewResourcePackClientResponsePacket() *ResourcePackClientResponsePacket {
	return &ResourcePackClientResponsePacket{packets.NewPacket(info.PacketIds[info.ResourcePackClientResponsePacket]), 0, []string{}}
}

func (pk *ResourcePackClientResponsePacket) Encode() {

}

func (pk *ResourcePackClientResponsePacket) Decode() {
	pk.Status = pk.GetByte()
	var idCount = pk.GetLittleShort()
	for i := int16(0); i < idCount; i++ {
		pk.PackUUIDs = append(pk.PackUUIDs, pk.GetString())
	}
}
