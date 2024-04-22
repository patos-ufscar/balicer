package common

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/patos-ufscar/http-web-server-example-go/models"
)

func Bind(port uint16) (*net.Listener, error) {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to bind to port %d: %s", port, err.Error()))
		return nil, fmt.Errorf("could not bind to port %d", port)
	}

	slog.Info(fmt.Sprintf("Listening on port %d", port))

	return &l, nil
}

func ParseHeader(header string) (*http.Request, error) {

	return nil, nil
}

func ParseHttpRequest(requestBytes []byte) (*models.HttpRequest, error) {
	httpReqFrame := models.NewHttpRequest()

	fmt.Println(string(requestBytes))

	lines := strings.Split(string(requestBytes), "\r\n")
	// request := [][]string{}
	for i, v := range lines {
		if i == 0 {
			words := strings.Split(v, " ")
			httpReqFrame.Method = words[0]
			httpReqFrame.RequestURI = words[1]
			httpReqFrame.HTTPVersion = words[2]
		} else if i == 1 {
			words := strings.Split(v, " ")
			httpReqFrame.Host = words[1]
		} else {
			words := strings.SplitN(v, ": ", 2)
			if len(words[0]) == len(v) {
				continue
			}

			httpReqFrame.Headers[textproto.CanonicalMIMEHeaderKey(words[0])] = words[1]
			// parseRequestLine(&httpReqFrame, words)
		}
	}

	return &httpReqFrame, nil
}

// func parseRequestLine(frame *models.HttpRequest, words []string) error {

// 	// textproto.CanonicalMIMEHeaderKey(words[0])
// 	frame.Headers[textproto.CanonicalMIMEHeaderKey(words[0])] = words[1]

// 	// switch words[0] {
// 	// case "Content-Type":
// 	// 	frame.RequestHeaders.ContentType = words[1]
// 	// case "Content-Length":
// 	// 	val, err := strconv.ParseUint(words[1], 10, 64)
// 	// 	if err != nil {
// 	// 		return errors.New("could not convert to uint64")
// 	// 	}
// 	// 	frame.RequestHeaders.ContentLength = val
// 	// case "Content-Encoding":
// 	// 	frame.RequestHeaders.ContentEncoding = words[1]
// 	// case "Content-Language":
// 	// 	frame.RequestHeaders.ContentLanguage = words[1]

// 	// default:
// 	// 	return errors.New("could not find match")
// 	// }

// 	return nil
// }