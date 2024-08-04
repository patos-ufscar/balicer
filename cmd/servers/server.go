package servers

import "net"

type Server interface {
	// Bind(port uint16) 					(*net.Listener, error)
	ValidHost(host string) bool
	Serve(lis net.Listener)
	HandleConnection(conn net.Conn)
}
