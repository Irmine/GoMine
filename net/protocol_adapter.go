package net

import (
	"github.com/golang/geo/r3"
	"github.com/google/uuid"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/gomine/net/protocol"
	"github.com/irmine/gomine/packs"
	"github.com/irmine/worlds/blocks"
	"github.com/irmine/worlds/chunks"
	"github.com/irmine/worlds/entities/data"
)

func (session *MinecraftSession) SendAddEntity(entity protocol.AddEntityEntry) {
	session.SendPacket(session.adapter.packetManager.GetAddEntity(entity))
}

func (session *MinecraftSession) SendAddPlayer(uuid uuid.UUID, player protocol.AddPlayerEntry) {
	session.SendPacket(session.adapter.packetManager.GetAddPlayer(uuid, player))
}

func (session *MinecraftSession) SendChunkRadiusUpdated(radius int32) {
	session.SendPacket(session.adapter.packetManager.GetChunkRadiusUpdated(radius))
}

func (session *MinecraftSession) SendCraftingData() {
	session.SendPacket(session.adapter.packetManager.GetCraftingData())
}

func (session *MinecraftSession) SendDisconnect(message string, hideDisconnect bool) {
	session.SendPacket(session.adapter.packetManager.GetDisconnect(message, hideDisconnect))
}

func (session *MinecraftSession) SendFullChunkData(chunk *chunks.Chunk) {
	session.SendPacket(session.adapter.packetManager.GetFullChunkData(chunk))
}

func (session *MinecraftSession) SendMovePlayer(runtimeId uint64, position r3.Vector, rotation data.Rotation, mode byte, onGround bool, ridingRuntimeId uint64) {
	session.SendPacket(session.adapter.packetManager.GetMovePlayer(runtimeId, position, rotation, mode, onGround, ridingRuntimeId))
}

func (session *MinecraftSession) SendPlayerList(listType byte, players map[string]protocol.PlayerListEntry) {
	session.SendPacket(session.adapter.packetManager.GetPlayerList(listType, players))
}

func (session *MinecraftSession) SendPlayStatus(status int32) {
	session.SendPacket(session.adapter.packetManager.GetPlayStatus(status))
}

func (session *MinecraftSession) SendRemoveEntity(uniqueId int64) {
	session.SendPacket(session.adapter.packetManager.GetRemoveEntity(uniqueId))
}

func (session *MinecraftSession) SendResourcePackChunkData(packUUID string, chunkIndex int32, progress int64, data []byte) {
	session.SendPacket(session.adapter.packetManager.GetResourcePackChunkData(packUUID, chunkIndex, progress, data))
}

func (session *MinecraftSession) SendResourcePackDataInfo(pack packs.Pack) {
	session.SendPacket(session.adapter.packetManager.GetResourcePackDataInfo(pack))
}

func (session *MinecraftSession) SendResourcePackInfo(mustAccept bool, resourcePacks *packs.Stack, behaviorPacks *packs.Stack) {
	session.SendPacket(session.adapter.packetManager.GetResourcePackInfo(mustAccept, resourcePacks, behaviorPacks))
}

func (session *MinecraftSession) SendResourcePackStack(mustAccept bool, resourcePacks *packs.Stack, behaviorPacks *packs.Stack) {
	session.SendPacket(session.adapter.packetManager.GetResourcePackStack(mustAccept, resourcePacks, behaviorPacks))
}

func (session *MinecraftSession) SendServerHandshake(encryptionJwt string) {
	session.SendPacket(session.adapter.packetManager.GetServerHandshake(encryptionJwt))
}

func (session *MinecraftSession) SendSetEntityData(runtimeId uint64, data map[uint32][]interface{}) {
	session.SendPacket(session.adapter.packetManager.GetSetEntityData(runtimeId, data))
}

func (session *MinecraftSession) SendStartGame(player protocol.StartGameEntry, runtimeIdsTable []byte) {
	session.SendPacket(session.adapter.packetManager.GetStartGame(player, runtimeIdsTable))
}

func (session *MinecraftSession) SendText(text types.Text) {
	session.SendPacket(session.adapter.packetManager.GetText(text))
}

func (session *MinecraftSession) Transfer(address string, port uint16) {
	session.SendPacket(session.adapter.packetManager.GetTransfer(address, port))
}

func (session *MinecraftSession) SendUpdateAttributes(runtimeId uint64, attributes data.AttributeMap) {
	session.SendPacket(session.adapter.packetManager.GetUpdateAttributes(runtimeId, attributes))
}

func (session *MinecraftSession) SendNetworkChunkPublisherUpdate(position blocks.Position, radius uint32) {
	session.SendPacket(session.adapter.packetManager.GetNetworkChunkPublisherUpdatePacket(position, radius))
}

func (session *MinecraftSession) SendMoveEntity(runtimeId uint64, position r3.Vector, rot data.Rotation, flags byte, teleport bool) {
	session.SendPacket(session.adapter.packetManager.GetMoveEntity(runtimeId, position, rot, flags, teleport))
}

func (session *MinecraftSession) SendPlayerSkin(uuid2 uuid.UUID, skinId, geometryName, geometryData string, skinData, capeData []byte) {
	session.SendPacket(session.adapter.packetManager.GetPlayerSkin(uuid2, skinId, geometryName, geometryData, skinData, capeData))
}

func (session *MinecraftSession) SendPlayerAction(runtimeId uint64, action int32, position blocks.Position, face int32) {
	session.SendPacket(session.adapter.packetManager.GetPlayerAction(runtimeId, action, position, face))
}

func (session *MinecraftSession) SendAnimate(action int32, runtimeId uint64, float float32) {
	session.SendPacket(session.adapter.packetManager.GetAnimate(action, runtimeId, float))
}