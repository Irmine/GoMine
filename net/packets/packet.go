package packets

import (
	"github.com/irmine/gomine/text"
)

// IPacket gets implemented by every packet.
// Every packet can be encoded and decoded.
type IPacket interface {
	SetBuffer([]byte)
	GetBuffer() []byte
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
	GetId() int
}

// Packet is a Minecraft mcpe packet.
// Packets have a given ID and contain two prefix
// bytes, which are used for split screen.
// Packets can be discarded during handling
// of the packets to stop other handlers from
// handling those packets.
type Packet struct {
	*MinecraftStream
	// PacketId is the ID of the packet.
	// Packet IDs may differ for different protocols.
	PacketId int
	// SenderIdentifier is used for split screen.
	// It specifies the sender sub ID.
	SenderIdentifier byte
	// ReceiverIdentifier is used for split screen.
	// It specifies the receiver sub ID.
	ReceiverIdentifier byte
	discarded          bool
}

// NewPacket returns a new packet with packet ID.
// The packet's stream gets pre-initialized.
func NewPacket(id int) *Packet {
	return &Packet{NewMinecraftStream(), id, 0, 0, false}
}

// GetId returns the packet ID of the packet.
func (pk *Packet) GetId() int {
	return pk.PacketId
}

// Discard discards the packet.
// Once discarded, handlers will no longer
// handle this packet.
func (pk *Packet) Discard() {
	pk.discarded = true
}

// IsDiscard checks if a packet has been discarded.
// Discarded packets are no longer processed,
// and get disposed immediately.
func (pk *Packet) IsDiscarded() bool {
	return pk.discarded
}

// EncodeId encodes the ID of the packet.
func (pk *Packet) EncodeId() {
	pk.PutUnsignedVarInt(uint32(pk.PacketId))
}

// DecodeId decodes the packet ID of the packet.
// The function panics if the packet ID
// and read ID do not match.
func (pk *Packet) DecodeId() {
	id := int(pk.GetUnsignedVarInt())
	if id != pk.PacketId {
		text.DefaultLogger.Debug("Packet IDs do not match. Expected:", pk.PacketId, "Got:", id)
	}
}

// EncodeHeader encodes the header of a packet,
// with mcpe >= 200.
// First the packet ID gets encoded,
// after which the sender and receiver ID bytes get written.
func (pk *Packet) EncodeHeader() {
	pk.EncodeId()
}

// DecodeHeader decodes a header of a packet,
// with mcpe >= 200.
// First the packet ID gets decoded,
// after which the sender and receiver ID bytes.
func (pk *Packet) DecodeHeader() {
	pk.DecodeId()
}

func (pk *Packet) Encode() {}

func (pk *Packet) Decode() {}
