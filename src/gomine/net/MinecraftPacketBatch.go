package net

import (
	"gomine/utils"
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"gomine/interfaces"
	"fmt"
)

const McpeFlag = 0xFE

type MinecraftPacketBatch struct {
	stream *utils.BinaryStream

	raw []byte

	packets []interfaces.IPacket
}

/**
 * Returns a new Minecraft Packet Batch used to decode/encode batches from Encapsulated Packets.
 */
func NewMinecraftPacketBatch() *MinecraftPacketBatch {
	var batch = &MinecraftPacketBatch{}
	batch.stream = utils.NewStream()

	return batch
}

/**
 * Returns the Binary stream of this batch.
 */
func (batch *MinecraftPacketBatch) GetStream() *utils.BinaryStream {
	return batch.stream
}

/**
 * Decodes the batch and separates packets. This does not decode the packets.
 */
func (batch *MinecraftPacketBatch) Decode() {
	var mcpeFlag = batch.stream.GetByte()
	if mcpeFlag != McpeFlag {
		return
	}

	var reader = bytes.NewReader(batch.stream.Buffer[batch.stream.Offset:])
	var zlibReader, _ = zlib.NewReader(reader)
	defer zlibReader.Close()

	batch.raw, _ = ioutil.ReadAll(zlibReader)

	batch.stream.ResetStream()
	batch.stream.SetBuffer(batch.raw)

	var packetData [][]byte

	for !batch.stream.Feof() {
		packetData = append(packetData, []byte(batch.stream.GetString()))
	}

	for _, data := range packetData {
		packetId := int(data[0])

		if !IsPacketRegistered(packetId) {
			fmt.Println("Unhandled Minecraft packet with ID:", packetId)
			continue
		}
		packet := GetPacket(packetId)

		packet.ResetStream()

		packet.SetBuffer(data)
		batch.packets = append(batch.packets, packet)
	}

	return
}

/**
 * Encodes all packets in the batch.
 */
func (batch *MinecraftPacketBatch) Encode() {
	batch.stream.ResetStream()
	batch.stream.PutByte(McpeFlag)

	var stream = utils.NewStream()
	for _, packet := range batch.GetPackets() {
		stream.PutUnsignedVarInt(uint32(len(packet.GetBuffer())))
		stream.PutBytes(packet.GetBuffer())
	}

	var buff = bytes.Buffer{}
	var writer = zlib.NewWriter(&buff)
	writer.Write(stream.Buffer)
	writer.Close()

	batch.stream.PutBytes(buff.Bytes())
}

/**
 * Adds a packet to the batch when encoding.
 */
func (batch *MinecraftPacketBatch) AddPacket(packet interfaces.IPacket) {
	batch.packets = append(batch.packets, packet)
}

/**
 * Returns all packets inside of the batch.
 * This only returns correctly when done after decoding, or before encoding.
 */
func (batch *MinecraftPacketBatch) GetPackets() []interfaces.IPacket {
	return batch.packets
}
