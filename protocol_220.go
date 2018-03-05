package gomine

type P220 struct {
	*P200
}

func NewP220(server *Server) *P220 {
	var proto = &P220{NewP200(server)}
	proto.ProtocolNumber = 220

	return proto
}
