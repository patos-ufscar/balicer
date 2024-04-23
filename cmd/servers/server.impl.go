package servers

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"regexp"

	"github.com/patos-ufscar/http-web-server-example-go/common"
	"github.com/patos-ufscar/http-web-server-example-go/handlers"
	"github.com/patos-ufscar/http-web-server-example-go/models"
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

func (s *ServerImpl) ValidHost(host string) bool {

	for _, v := range s.hostsRegs {
		if v.Match([]byte(host)) {
			return true
		}
	}

	return false
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

			go s.HandleConnection(conn)

		}(conn)
	}
}

func (s *ServerImpl) HandleConnection(conn net.Conn) {
	defer conn.Close()

	readBuffer := make([]byte, 8 * 1<<10)
	_, err := conn.Read(readBuffer)
	if err != nil {
		if err == io.EOF {
			slog.Warn("Connection closed by the server")
		} else {
			slog.Error(err.Error())
		}
		return
	}

	req, err := common.ParseBaseRequest(readBuffer)
	if err != nil {
		return
	}

	if !s.ValidHost(req.Host) {
		return
	}

	// TODO: Here would be TLS

	req, err = common.ParseHttpRequest(readBuffer)
	if err != nil {
		return
	}

	// slog.Debug(fmt.Sprintf("req: %+v", req))

	var rep models.HttpResponse
	for _, v := range s.handlers {
		if v.ValidPath(req.RequestURI) {
			rep, err = v.Handle(*req)
			if err != nil {
				slog.Error(err.Error())
				return
			}
		}
	}

	// fazer o try except aqui (recover)

	slog.Info(fmt.Sprintf("%s %s %s %d", req.Method, req.RequestURI, req.HTTPVersion, rep.StatusCode))

	err = utils.ReplyHTTP(conn, rep.DumpResponse())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	// OLD CODE: using more http package funcs
	// reader := bufio.NewReader(conn)
	// req, err := http.ReadRequest(reader)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	return
	// }
}
