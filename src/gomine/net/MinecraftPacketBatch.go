package net

import (
	"gomine/utils"
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"gomine/net/packets"
)

const McpeFlag = 0xFE

type MinecraftPacketBatch struct {
	stream *utils.BinaryStream

	raw []byte

	packets []packets.IPacket
}

func NewMinecraftPacketBatch() MinecraftPacketBatch {
	var batch = MinecraftPacketBatch{}
	batch.stream = utils.NewStream()

	return batch
}

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
		packet := packets.GetPacket(packetId)

		packet.SetBuffer(data)
		batch.packets = append(batch.packets, packet)
	}

	return
}

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

func (batch *MinecraftPacketBatch) AddPacket(packet packets.IPacket) {
	batch.packets = append(batch.packets, packet)
}

func (batch *MinecraftPacketBatch) GetPackets() []packets.IPacket {
	return batch.packets
}
