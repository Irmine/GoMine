package net

import (
	"gomine/utils"
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"gomine/interfaces"
	"crypto/cipher"
	"errors"
	"encoding/hex"
)

const McpeFlag = 0xFE

type MinecraftPacketBatch struct {
	*utils.BinaryStream

	raw []byte

	packets []interfaces.IPacket

	session interfaces.IMinecraftSession
	needsEncryption bool
	logger interfaces.ILogger
}

/**
 * Returns a new Minecraft Packet Batch used to decode/encode batches from Encapsulated Packets.
 */
func NewMinecraftPacketBatch(session interfaces.IMinecraftSession, logger interfaces.ILogger) *MinecraftPacketBatch {
	var batch = &MinecraftPacketBatch{}
	batch.BinaryStream = utils.NewStream()
	batch.session = session

	if session == nil {
		batch.needsEncryption = false
	} else {
		if session.IsInitialized() {
			batch.needsEncryption = session.UsesEncryption()
		} else {
			batch.needsEncryption = false
		}
	}

	batch.logger = logger

	return batch
}

/**
 * Decodes the batch and separates packets. This does not decode the packets.
 */
func (batch *MinecraftPacketBatch) Decode() {
	var mcpeFlag = batch.GetByte()
	if mcpeFlag != McpeFlag {
		return
	}
	batch.raw = batch.Buffer[batch.Offset:]

	if batch.needsEncryption {
		batch.decrypt()
	}
	var err = batch.decompress()
	if err != nil {
		batch.logger.LogError(err)
		return
	}

	batch.ResetStream()
	batch.SetBuffer(batch.raw)

	var packetData [][]byte

	for !batch.Feof() {
		packetData = append(packetData, batch.GetLengthPrefixedBytes())
	}

	batch.fetchPackets(packetData)
}

/**
 * Encodes all packets in the batch and zlib encodes them.
 */
func (batch *MinecraftPacketBatch) Encode() {
	batch.ResetStream()
	batch.PutByte(McpeFlag)

	var stream = utils.NewStream()
	batch.putPackets(stream)

	var zlibData = batch.compress(stream)
	var data = zlibData
	if batch.needsEncryption {
		data = batch.encrypt(data)
	}

	batch.PutBytes(data)
}

/**
 * Fetches all packets from the raw packet buffers.
 */
func (batch *MinecraftPacketBatch) fetchPackets(packetData [][]byte) {
	for _, data := range packetData {
		if len(data) == 0 {
			continue
		}
		packetId := int(data[0])

		if !IsPacketRegistered(packetId) {
			batch.logger.Debug("Unknown Minecraft packet with ID:", packetId)
			continue
		}
		packet := GetPacket(packetId)

		packet.SetBuffer(data)
		batch.packets = append(batch.packets, packet)
	}
}

/**
 * Encrypts the data passed to the function.
 */
func (batch *MinecraftPacketBatch) encrypt(d []byte) []byte {
	var data = batch.session.GetEncryptionHandler().Data
	d = append(d, batch.session.GetEncryptionHandler().ComputeSendChecksum(d)...)

	for i := range d {
		var cfb = cipher.NewCFBEncrypter(data.EncryptCipher, data.EncryptIV)
		cfb.XORKeyStream(d[i:i + 1], d[i:i + 1])
		data.EncryptIV = append(data.EncryptIV[1:], d[i])
	}
	
	return d
}

/**
 * Decrypts the buffer of the packet.
 */
func (batch *MinecraftPacketBatch) decrypt() {
	var data = batch.session.GetEncryptionHandler().Data
	for i, b := range batch.raw {
		var cfb = cipher.NewCFBDecrypter(data.DecryptCipher, data.DecryptIV)
		cfb.XORKeyStream(batch.raw[i:i + 1], batch.raw[i:i + 1])
		data.DecryptIV = append(data.DecryptIV[1:], b)
	}
}

/**
 * Puts all packets of the batch inside of the stream.
 */
func (batch *MinecraftPacketBatch) putPackets(stream *utils.BinaryStream) {
	for _, packet := range batch.GetPackets() {
		packet.EncodeHeader()
		packet.Encode()
		stream.PutLengthPrefixedBytes(packet.GetBuffer())
	}
}

/**
 * Zlib compresses the data in the stream and returns it.
 */
func (batch *MinecraftPacketBatch) compress(stream *utils.BinaryStream) []byte {
	var buff = bytes.Buffer{}
	var writer = zlib.NewWriter(&buff)
	writer.Write(stream.Buffer)
	writer.Close()

	return buff.Bytes()
}

/**
 * Decompresses the zlib compressed buffer.
 */
func (batch *MinecraftPacketBatch) decompress() error {
	var reader = bytes.NewReader(batch.raw)
	zlibReader, err := zlib.NewReader(reader)
	batch.logger.LogError(err)

	if err != nil {
		println(hex.EncodeToString(batch.raw))
		return err
	}
	if zlibReader == nil {
		return errors.New("an error occurred when decompressing zlib")
	}
	zlibReader.Close()

	batch.raw, err = ioutil.ReadAll(zlibReader)

	return err
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