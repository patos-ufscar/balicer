package servers

import "net"

type Server interface {
	// Bind(port uint16) 					(*net.Listener, error)
	Serve(lis net.Listener)
	HandleConnection(conn net.Conn)
}
