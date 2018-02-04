package interfaces

import (
	"goraklib/server"
	"gomine/entities/math"
	"gomine/vectors"
	"gomine/net/packets/types"
)

type IPlayerFactory interface {
	AddPlayer(IPlayer, *server.Session)
	GetPlayers() map[string]IPlayer
	GetPlayerByName(string) (IPlayer, error)
	GetPlayerBySession(*server.Session) (IPlayer, error)
	GetPlayerCount() uint
	RemovePlayer(player IPlayer)
	PlayerExistsBySession(session *server.Session) bool
	PlayerExists(name string) bool
}

type IPlayer interface {
	IEntity
	IMinecraftSession
	GetName() string
	SetName(name string)
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
	NewMinecraftSession(IServer, *server.Session, types.SessionData) IMinecraftSession
	New(IServer, IMinecraftSession, string) IPlayer
	SyncMove(float32, float32, float32, float32, float32, float32, bool)
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
	SetMinecraftSession(IMinecraftSession)
}