package handlers

import (
	"net"

	"github.com/patos-ufscar/http-web-server-example-go/models"
)

type Handler interface {
	ValidHost(host string)									bool
	Handle(conn net.Conn, req models.HttpRequestFrame)		error
}
