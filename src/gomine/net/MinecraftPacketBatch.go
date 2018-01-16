package net

import (
	"gomine/utils"
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"gomine/interfaces"
	"crypto/cipher"
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
func (batch *MinecraftPacketBatch) Decode(player interfaces.IPlayer, logger interfaces.ILogger) {
	var mcpeFlag = batch.GetByte()
	if mcpeFlag != McpeFlag {
		return
	}
	var err error

	batch.raw = batch.Buffer[batch.Offset:]

	if player.UsesEncryption() {
		for i, b := range batch.raw {
			var cfb = cipher.NewCFBDecrypter(player.GetEncryptionHandler().Data.Cipher, player.GetEncryptionHandler().Data.IV)
			cfb.XORKeyStream(batch.raw[i:i + 1], batch.raw[i:i + 1])
			player.GetEncryptionHandler().Data.IV = append(player.GetEncryptionHandler().Data.IV[1:], b)
		}
	}

	var reader = bytes.NewReader(batch.raw)
	zlibReader, err := zlib.NewReader(reader)
	logger.LogError(err)
	if zlibReader == nil {
		return
	}
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