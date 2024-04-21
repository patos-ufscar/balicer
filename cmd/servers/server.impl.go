package servers

import (
	"fmt"
	"log/slog"
	"net"
	"regexp"

	"github.com/patos-ufscar/http-web-server-example-go/handlers"
	"github.com/patos-ufscar/http-web-server-example-go/utils"
)

type ServerImpl struct {
	port					uint16
	hostsRegs				[]regexp.Regexp
	handlers				[]handlers.Handler
}

func NewServer(port uint16, hostsRegs []regexp.Regexp, handlers []handlers.Handler) Server {
	return &ServerImpl{
		port: port,
		hostsRegs: hostsRegs,
		handlers: handlers,
	}
}

func (s *ServerImpl) Bind(port uint16) (*net.Listener, error) {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to bind to port %d: %s", port, err.Error()))
		return nil, fmt.Errorf("could not bind to port %d", port)
	}

	slog.Info(fmt.Sprintf("Listening on port %d", port))

	return &l, nil
}

func (s *ServerImpl) Serve(lis net.Listener) {

	for {
		conn, err := lis.Accept()
		if err != nil {
			slog.Error(fmt.Sprintf("Error accepting connection: %s", err.Error()))
			continue
		}

		go func (conn net.Conn) {
			// Recover Func
			defer func(conn net.Conn) {
				// we re-reply in case of error (reply missing)
				r := recover()
				if r != nil {
					slog.Error(fmt.Sprint("Recovered from: ", r))
					err := utils.Reply502(conn)
					if err != nil {
						slog.Error(fmt.Sprintf("Could not reply: %s", err.Error()))
					}
				}
			}(conn)
			defer conn.Close()

			utils.ReplyString(conn, "OK")

			// iterate over handlers

			// Getting the handler to handle
			// rep, err := handlers.HandleGlobal(conn, config)
			// if err != nil {
			// 	panic(err)
			// } else {
			// 	slog.Info(fmt.Sprintf("handled: %s", string(rep)))
			// }

		}(conn)
	}
}
