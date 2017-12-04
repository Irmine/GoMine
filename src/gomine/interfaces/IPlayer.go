package interfaces

import (
	"goraklib/server"
	"gomine/vectors"
)

type IPlayer interface {
	GetSession() *server.Session
	GetName() string
	GetDisplayName() string
	SetDisplayName(string)
	GetPermissionGroup() IPermissionGroup
	SetPermissionGroup(IPermissionGroup)
	HasPermission(string) bool
	AddPermission(IPermission) bool
	RemovePermission(string) bool
	GetServer() IServer
	SetViewDistance(int32)
	GetViewDistance() int32
	GetUUID() string
	GetXUID() string
	SetLanguage(string)
	GetLanguage() string
	GetClientId() int
	SetSkinId(id string)
	GetSkinId() string
	GetSkinData() []byte
	SetSkinData([]byte)
	GetCapeData() []byte
	SetCapeData([]byte)
	GetGeometryName() string
	SetGeometryName(string)
	GetGeometryData() string
	SetGeometryData(string)
	SetPosition(vector *vectors.TripleVector)
	GetPosition() *vectors.TripleVector
	//SetDimension(worlds.Dimension)
	GetDimension() IDimension
	SendChunk(IChunk)
	New(IServer, *server.Session, string, string, string, int) IPlayer
	Tick()
}
