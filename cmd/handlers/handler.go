package handlers

import (
	"fmt"
	"net"

	"github.com/patos-ufscar/http-web-server-example-go/models"
)

type Handler interface {
	ValidPath(host string)									bool
	Handle(conn net.Conn, req models.HttpRequest)			error
}

func HandlerFactory(locConf models.LocationConfig) (Handler, error) {
	switch locConf.ReturnType {
	case "static":
		return NewHandlerStaticImpl(locConf), nil
	}

	// return nil, errors.New("not yet implemented")
	return nil, fmt.Errorf("not yet implemented handler: %s", locConf.ReturnType)
}
