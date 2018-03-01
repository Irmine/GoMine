package net

import (
	"github.com/irmine/gomine/utils"
	"github.com/irmine/goraklib/server"
	"sync"
)

// SessionManager is a struct managing Minecraft sessions.
// A session manager holds multiple maps used to find sessions by given keys.
type SessionManager struct {
	mutex      sync.RWMutex
	nameMap    map[string]*MinecraftSession
	uuidMap    map[utils.UUID]*MinecraftSession
	xuidMap    map[string]*MinecraftSession
	sessionMap map[string]*MinecraftSession
}

// NewSessionManager returns a new session manager.
func NewSessionManager() *SessionManager {
	return &SessionManager{sync.RWMutex{}, make(map[string]*MinecraftSession), make(map[utils.UUID]*MinecraftSession), make(map[string]*MinecraftSession), make(map[string]*MinecraftSession)}
}

// GetSessions returns the name => session map of the manager.
func (manager *SessionManager) GetSessions() map[string]*MinecraftSession {
	return manager.nameMap
}

// AddMinecraftSession adds the given Minecraft session to the manager.
func (manager *SessionManager) AddMinecraftSession(session *MinecraftSession) {
	manager.mutex.Lock()
	manager.nameMap[session.GetPlayer().GetName()] = session
	manager.uuidMap[session.GetUUID()] = session
	manager.xuidMap[session.GetXUID()] = session
	manager.sessionMap[server.GetSessionIndex(session.GetSession())] = session
	manager.mutex.Unlock()
}

// RemoveMinecraftSession removes a Minecraft session from the manager.
func (manager *SessionManager) RemoveMinecraftSession(session *MinecraftSession) {
	manager.mutex.Lock()
	delete(manager.nameMap, session.GetPlayer().GetName())
	delete(manager.uuidMap, session.GetUUID())
	delete(manager.xuidMap, session.GetXUID())
	delete(manager.sessionMap, server.GetSessionIndex(session.GetSession()))
	manager.mutex.Unlock()
}

// GetSessionCount returns the session count of the manager.
func (manager *SessionManager) GetSessionCount() int {
	return len(manager.nameMap)
}

// HasSession checks if the session manager has a session with the given name.
func (manager *SessionManager) HasSession(name string) bool {
	manager.mutex.RLock()
	var _, ok = manager.nameMap[name]
	manager.mutex.RUnlock()
	return ok
}

// GetSession attempts to retrieve a session by its name.
// A bool is returned indicating success.
func (manager *SessionManager) GetSession(name string) (*MinecraftSession, bool) {
	manager.mutex.RLock()
	var session, ok = manager.nameMap[name]
	manager.mutex.RUnlock()
	return session, ok
}

// HasSessionWithRakNetSession checks if the session manager has a session with the given RakNet session.
func (manager *SessionManager) HasSessionWithRakNetSession(rakNetSession *server.Session) bool {
	manager.mutex.RLock()
	var _, ok = manager.sessionMap[server.GetSessionIndex(rakNetSession)]
	manager.mutex.RUnlock()
	return ok
}

// GetSessionByRakNetSession attempts to retrieve a session by its RakNet session.
// A bool is returned indicating success.
func (manager *SessionManager) GetSessionByRakNetSession(rakNetSession *server.Session) (*MinecraftSession, bool) {
	manager.mutex.RLock()
	var session, ok = manager.sessionMap[server.GetSessionIndex(rakNetSession)]
	manager.mutex.RUnlock()
	return session, ok
}

// HasSessionWithXUID checks if the session manager has a session with the given XUID.
func (manager *SessionManager) HasSessionWithXUID(xuid string) bool {
	manager.mutex.RLock()
	var _, ok = manager.xuidMap[xuid]
	manager.mutex.RUnlock()
	return ok
}

// GetSessionByXUID attempts to retrieve a session by its XUID.
// A bool is returned indicating success.
func (manager *SessionManager) GetSessionByXUID(xuid string) (*MinecraftSession, bool) {
	manager.mutex.RLock()
	var session, ok = manager.xuidMap[xuid]
	manager.mutex.RUnlock()
	return session, ok
}

// HasSessionWithUUID checks if the session manager has a session with the given UUID.
func (manager *SessionManager) HasSessionWithUUID(uuid utils.UUID) bool {
	manager.mutex.RLock()
	var _, ok = manager.uuidMap[uuid]
	manager.mutex.RUnlock()
	return ok
}

// GetSessionByUUID attempts to retrieve a session by its UUID.
// A bool is returned indicating success.
func (manager *SessionManager) GetSessionByUUID(uuid utils.UUID) (*MinecraftSession, bool) {
	manager.mutex.RLock()
	var session, ok = manager.uuidMap[uuid]
	manager.mutex.RUnlock()
	return session, ok
}
