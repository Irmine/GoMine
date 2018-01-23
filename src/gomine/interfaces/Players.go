package interfaces

import (
	"goraklib/server"
	"gomine/entities/math"
	"gomine/vectors"
)

type IPlayerFactory interface {
	AddPlayer(IPlayer, *server.Session)
	GetPlayers() map[string]IPlayer
	GetPlayerByName(string) (IPlayer, error)
	GetPlayerBySession(*server.Session) (IPlayer, error)
	GetPlayerCount() uint
	RemovePlayer(player IPlayer)
}

type IPlayer interface {
	IEntity
	IMinecraftSession
	GetName() string
	GetDisplayName() string
	SetDisplayName(string)
	GetPermissionGroup() IPermissionGroup
	SetPermissionGroup(IPermissionGroup)
	HasPermission(string) bool
	AddPermission(IPermission) bool
	RemovePermission(string) bool
	SetViewDistance(int32)
	GetViewDistance() int32
	SetSkinId(string)
	GetSkinId() string
	GetSkinData() []byte
	SetSkinData([]byte)
	GetCapeData() []byte
	SetCapeData([]byte)
	GetGeometryName() string
	SetGeometryName(string)
	GetGeometryData() string
	SetGeometryData(string)
	SendChunk(IChunk, int)
	NewMinecraftSession(server IServer, session *server.Session, loginPacket IPacket) IMinecraftSession
	New(IServer, IMinecraftSession, string) IPlayer
	SyncMove(x, y, z, pitch, yaw, headYaw float32, onGround bool)
	SendMessage(string)
	PlaceInWorld(*vectors.TripleVector, *math.Rotation, ILevel, IDimension)
	HasChunkInUse(int) bool
	HasAnyChunkInUse() bool
	SpawnPlayerTo(IPlayer)
	SpawnPlayerToAll()
	IsFinalized() bool
	SetFinalized()
	UpdateAttributes()
	HasSpawned() bool
	SetSpawned(bool)
	Transfer(string, uint16)
}