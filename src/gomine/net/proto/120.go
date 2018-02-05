package proto

type Protocol120 struct {
	*Protocol160
}

func NewProtocol120() *Protocol120 {
	var proto = &Protocol120{NewProtocol160()}
	proto.protocolNumber = 120

	return proto
}
