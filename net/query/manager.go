package query

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// Manager handles the sequence of incoming queries.
type Manager struct {
	token  []byte
	result Result
}

// NewManager returns a new query manager with the given server.
func NewManager() Manager {
	var b = make([]byte, 4)
	rand.Read(b)
	return Manager{b, Result{}}
}

// SetQueryResult sets the query result data a query will receive when querying us.
func (manager *Manager) SetQueryResult(result Result) {
	manager.result = result
}

// HandleQuery handles an incoming query.
func (manager *Manager) HandleQuery(query *Query) {
	switch query.Header {
	case Challenge:
		var q = New(query.Address, query.Port)
		q.Header = Challenge
		q.QueryId = query.QueryId
		q.Token = manager.token

		manager.sendQuery(q)

	case Statistics:
		if string(manager.token) != string(query.Token) {
			return
		}

		var q = New(query.Address, query.Port)
		q.Header = Statistics
		q.QueryId = query.QueryId

		if query.IsShort {
			q.Statistics = manager.result.GetShort()
		} else {
			q.Statistics = manager.result.GetLong()
		}

		manager.sendQuery(q)
	}
}

// sendQuery sends a query to the address and port set in it.
func (manager *Manager) sendQuery(query *Query) error {
	query.EncodeServer()

	var conn, err = net.Dial("udp", query.Address+":"+strconv.Itoa(int(query.Port)))
	if err != nil {
		return err
	}
	conn.Write(query.Buffer)

	conn.Close()
	return nil
}

// Send queries a server with the given address and port.
// The call times out after the given timeout duration if no response is given.
//
// NOTE: This function is time consuming and should be used one a different goroutine where adequate.
func Send(address string, port uint16, timeout time.Duration) (Result, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var realAddresses, err = net.LookupIP(address)
	if err == nil {
		address = realAddresses[0].String()
	}

	var connection, err2 = net.Dial("udp", address+":"+strconv.Itoa(int(port)))
	connection.SetReadDeadline(time.Now().Add(timeout))

	if err2 != nil {
		return Result{}, err2
	}

	var q = New(address, port)
	q.Header = Challenge
	q.QueryId = int32(time.Now().Unix())
	q.EncodeClient()

	connection.Write(q.Buffer)

	var buf = make([]byte, 128)
	var bytesRead, err3 = connection.Read(buf)
	if err3 != nil {
		return Result{}, err3
	}

	buf = buf[:bytesRead]

	q = New(address, port)
	q.Buffer = buf
	q.DecodeClient()

	var statQuery = New(address, port)
	statQuery.Header = Statistics
	statQuery.QueryId = int32(time.Now().Unix())
	statQuery.Token = q.Token
	statQuery.EncodeClient()

	connection.Write(statQuery.Buffer)

	buf = make([]byte, 4096)
	var byteCount, err4 = connection.Read(buf)
	if err4 != nil {
		return Result{}, err4
	}
	connection.Close()

	buf = buf[:byteCount]

	q = New(address, port)
	q.Buffer = buf
	q.DecodeClient()

	var res = Result{}.ParseLong(q.Data)
	res.Port = port
	res.Address = address

	return res, nil
}
