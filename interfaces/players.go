package interfaces

import (
	"github.com/irmine/gomine/entities/math"
	"github.com/irmine/gomine/net/packets/types"
	"github.com/irmine/goraklib/server"
	"github.com/irmine/gomine/permissions"
	"github.com/golang/geo/r3"
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
	GetPermissionGroup() *permissions.Group
	SetPermissionGroup(*permissions.Group)
	HasPermission(string) bool
	AddPermission(*permissions.Permission) bool
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
	SyncMove(float64, float64, float64, float32, float32, float32, bool)
	SendMessage(...interface{})
	PlaceInWorld(r3.Vector, *math.Rotation, ILevel, IDimension)
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
