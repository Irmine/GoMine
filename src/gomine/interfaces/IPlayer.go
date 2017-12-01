package interfaces

import "goraklib/server"

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
	SetViewDistance(uint)
	GetViewDistance() uint
	GetUUID() string
	GetXUID() string
	SetLanguage(string)
	GetLanguage() string
	GetClientId() int
}
