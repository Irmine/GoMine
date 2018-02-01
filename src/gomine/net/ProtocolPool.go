package net

import (
	"gomine/interfaces"
)

type ProtocolPool struct {
	protocols map[int32]interfaces.IProtocol
}

func NewProtocolPool() *ProtocolPool {
	return &ProtocolPool{make(map[int32]interfaces.IProtocol)}
}

/**
 * Returns a protocol by its protocol number.
 */
func (pool *ProtocolPool) GetProtocol(protocolNumber int32) interfaces.IProtocol {
	return pool.protocols[protocolNumber]
}

/**
 * Registers the given protocol.
 */
func (pool *ProtocolPool) RegisterProtocol(protocol interfaces.IProtocol) {
	pool.protocols[protocol.GetProtocolNumber()] = protocol
}

/**
 * Checks if a protocol with the given protocol number is registered.
 */
func (pool *ProtocolPool) IsProtocolRegistered(protocolNumber int32) bool {
	var _, ok = pool.protocols[protocolNumber]
	return ok
}