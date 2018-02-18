package net

import (
	"github.com/irmine/gomine/entities/data"
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/interfaces"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/packs"
	"github.com/golang/geo/r3"
)

func (session *MinecraftSession) SendAddEntity(entity interfaces.IEntity) {
	session.SendPacket(session.protocol.GetAddEntity(entity))
}

func (session *MinecraftSession) SendAddPlayer(player interfaces.IPlayer) {
	session.SendPacket(session.protocol.GetAddPlayer(player))
}

func (session *MinecraftSession) SendChunkRadiusUpdated(radius int32) {
	session.SendPacket(session.protocol.GetChunkRadiusUpdated(radius))
}

func (session *MinecraftSession) SendCraftingData() {
	session.SendPacket(session.protocol.GetCraftingData())
}

func (session *MinecraftSession) SendDisconnect(message string, hideDisconnect bool) {
	session.SendPacket(session.protocol.GetDisconnect(message, hideDisconnect))
}

func (session *MinecraftSession) SendFullChunkData(chunk interfaces.IChunk) {
	session.SendPacket(session.protocol.GetFullChunkData(chunk))
}

func (session *MinecraftSession) SendMovePlayer(player interfaces.IPlayer, position r3.Vector, rotation math.Rotation, mode byte, onGround bool, ridingRuntimeId uint64) {
	session.SendPacket(session.protocol.GetMovePlayer(player.GetRuntimeId(), position, rotation, mode, onGround, ridingRuntimeId))
}

func (session *MinecraftSession) SendPlayerList(listType byte, players map[string]interfaces.IPlayer) {
	session.SendPacket(session.protocol.GetPlayerList(listType, players))
}

func (session *MinecraftSession) SendPlayStatus(status int32) {
	session.SendPacket(session.protocol.GetPlayStatus(status))
}

func (session *MinecraftSession) SendRemoveEntity(entity interfaces.IEntity) {
	session.SendPacket(session.protocol.GetRemoveEntity(entity.GetUniqueId()))
}

func (session *MinecraftSession) SendResourcePackChunkData(packUUID string, chunkIndex int32, progress int64, data []byte) {
	session.SendPacket(session.protocol.GetResourcePackChunkData(packUUID, chunkIndex, progress, data))
}

func (session *MinecraftSession) SendResourcePackDataInfo(pack packs.Pack) {
	session.SendPacket(session.protocol.GetResourcePackDataInfo(pack))
}

func (session *MinecraftSession) SendResourcePackInfo(mustAccept bool, resourcePacks []packs.Pack, behaviorPacks []packs.Pack) {
	session.SendPacket(session.protocol.GetResourcePackInfo(mustAccept, resourcePacks, behaviorPacks))
}

func (session *MinecraftSession) SendResourcePackStack(mustAccept bool, resourcePacks []packs.Pack, behaviorPacks []packs.Pack) {
	session.SendPacket(session.protocol.GetResourcePackStack(mustAccept, resourcePacks, behaviorPacks))
}

func (session *MinecraftSession) SendServerHandshake(encryptionJwt string) {
	session.SendPacket(session.protocol.GetServerHandshake(encryptionJwt))
}

func (session *MinecraftSession) SendSetEntityData(entity interfaces.IEntity, data map[uint32][]interface{}) {
	session.SendPacket(session.protocol.GetSetEntityData(entity, data))
}

func (session *MinecraftSession) SendStartGame(player interfaces.IPlayer) {
	session.SendPacket(session.protocol.GetStartGame(player))
}

func (session *MinecraftSession) SendText(text types.Text) {
	session.SendPacket(session.protocol.GetText(text))
}

func (session *MinecraftSession) Transfer(address string, port uint16) {
	session.SendPacket(session.protocol.GetTransfer(address, port))
}

func (session *MinecraftSession) SendUpdateAttributes(entity interfaces.IEntity, attributes *data.AttributeMap) {
	session.SendPacket(session.protocol.GetUpdateAttributes(entity, attributes))
}
