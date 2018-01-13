package net

import (
	"gomine/utils"
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"gomine/interfaces"
)

const McpeFlag = 0xFE

type MinecraftPacketBatch struct {
	*utils.BinaryStream

	raw []byte

	packets []interfaces.IPacket
}

/**
 * Returns a new Minecraft Packet Batch used to decode/encode batches from Encapsulated Packets.
 */
func NewMinecraftPacketBatch() *MinecraftPacketBatch {
	var batch = &MinecraftPacketBatch{}
	batch.BinaryStream = utils.NewStream()

	return batch
}

/**
 * Decodes the batch and separates packets. This does not decode the packets.
 */
func (batch *MinecraftPacketBatch) Decode(logger interfaces.ILogger) {
	var mcpeFlag = batch.GetByte()
	if mcpeFlag != McpeFlag {
		return
	}
	var err error

	var reader = bytes.NewReader(batch.Buffer[batch.Offset:])
	zlibReader, err := zlib.NewReader(reader)
	logger.LogError(err)
	defer zlibReader.Close()

	batch.raw, err = ioutil.ReadAll(zlibReader)
	logger.LogError(err)

	batch.ResetStream()
	batch.SetBuffer(batch.raw)

	var packetData [][]byte

	for !batch.Feof() {
		packetData = append(packetData, batch.Get(int(batch.GetUnsignedVarInt())))
	}

	for _, data := range packetData {
		if len(data) == 0 {
			continue
		}
		packetId := int(data[0])

		if !IsPacketRegistered(packetId) {
			logger.Debug("Unknown Minecraft packet with ID:", packetId)
			continue
		}
		packet := GetPacket(packetId)

		packet.SetBuffer(data)
		batch.packets = append(batch.packets, packet)
	}
}

/**
 * Encodes all packets in the batch and zlib encodes them.
 */
func (batch *MinecraftPacketBatch) Encode() {
	batch.ResetStream()
	batch.PutByte(McpeFlag)

	var stream = utils.NewStream()
	for _, packet := range batch.GetPackets() {
		packet.EncodeHeader()
		packet.Encode()
		stream.PutUnsignedVarInt(uint32(len(packet.GetBuffer())))
		stream.PutBytes(packet.GetBuffer())
	}

	var buff = bytes.Buffer{}
	var writer = zlib.NewWriter(&buff)
	writer.Write(stream.Buffer)
	writer.Close()

	batch.PutBytes(buff.Bytes())
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