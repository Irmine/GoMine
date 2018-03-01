package protocol

// Manager is a map managing several protocols, and supplies functions to register and deregister protocols.
type Manager map[int32]Protocol

// NewManager returns a new protocol manager.
func NewManager() Manager {
	return Manager{}
}

// GetProtocol returns a protocol by its protocol number.
func (pool Manager) GetProtocol(protocolNumber int32) Protocol {
	if !pool.IsProtocolRegistered(protocolNumber) {
		return pool[200]
	}
	return pool[protocolNumber]
}

// RegisterProtocol registers the given protocol.
func (pool Manager) RegisterProtocol(protocol Protocol) {
	pool[protocol.GetProtocolNumber()] = protocol
}

// IsProtocolRegistered checks if a protocol with the given protocol number is registered.
func (pool Manager) IsProtocolRegistered(protocolNumber int32) bool {
	var _, ok = pool[protocolNumber]
	return ok
}

// DeregisterProtocol deregisters a protocol from the pool.
func (pool Manager) DeregisterProtocol(protocolNumber int32) {
	delete(pool, protocolNumber)
}
