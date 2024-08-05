package servers

import (
	"fmt"
	"log/slog"
	"net"
	"regexp"

	"github.com/patos-ufscar/http-web-server-example-go/common"
	"github.com/patos-ufscar/http-web-server-example-go/handlers"
	"github.com/patos-ufscar/http-web-server-example-go/models"
	"github.com/patos-ufscar/http-web-server-example-go/utils"
)

const (
	READ_BUFFER_SIZE int32 = 32 * 1 << 10
	READ_DEADLINE_MS int32 = 100
)

type ServerImpl struct {
	port      uint16
	hostsRegs []regexp.Regexp
	handlers  []handlers.Handler
}

func NewServer(port uint16, hostsRegs []regexp.Regexp, handlers []handlers.Handler) Server {
	return &ServerImpl{
		port:      port,
		hostsRegs: hostsRegs,
		handlers:  handlers,
	}
}

func (s *ServerImpl) ValidHost(host string) bool {

	for _, v := range s.hostsRegs {
		// fmt.Println(v.MatchString(host))
		if v.MatchString(host) {
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

		go s.HandleConnection(conn)

		// go func(conn net.Conn) {
		// 	// Recover Func
		// 	defer func(conn net.Conn) {
		// 		// we re-reply in case of error (reply missing)
		// 		r := recover()
		// 		if r != nil {
		// 			slog.Error(fmt.Sprint("Recovered from: ", r))
		// 			err := utils.Reply502(conn)
		// 			if err != nil {
		// 				slog.Error(fmt.Sprintf("Could not reply: %s", err.Error()))
		// 			}
		// 		}
		// 	}(conn)

		// 	go s.HandleConnection(conn)

		// }(conn)
	}
}

func (s *ServerImpl) HandleConnection(conn net.Conn) {
	defer conn.Close()
	// conn.SetReadDeadline(time.Now().Add(time.Duration(READ_DEADLINE_MS) * time.Millisecond))

	readBytes, err := common.ReadBytesFromConn(conn)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	req := models.ParseHttpRequest(readBytes)

	host, ok := req.Headers["Host"]
	if !ok || !s.ValidHost(host) {
		return
	}

	// TODO: Here would be TLS

	slog.Debug(fmt.Sprintf("req: %+v", req))

	var rep models.HttpResponse
	for _, v := range s.handlers {
		if v.ValidPath(req.RequestURI) {
			rep, err = v.Handle(req)
			if err != nil {
				slog.Error(err.Error())
				return
			}
		}
	}

	// fazer o try except aqui (recover)

	slog.Info(fmt.Sprintf("%s %s %s %d", req.Method, req.RequestURI, req.HTTPVersion, rep.StatusCode))

	resp := rep.DumpResponse()
	fmt.Printf("resp: %v\n", string(resp))
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
