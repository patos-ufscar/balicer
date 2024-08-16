package models

import (
	"bytes"
	"log/slog"
	"net/textproto"
	"strconv"
	"strings"
)

const (
	CRLF        string = "\r\n"
	DOUBLE_CRLF string = CRLF + CRLF
)

type HttpRequest struct {
	Method      string
	RequestURI  string
	HTTPVersion string
	Headers     map[string]string
	Body        []byte
}

func NewEmptyHttpRequest() HttpRequest {
	var httpReq HttpRequest
	httpReq.Headers = make(map[string]string)
	return httpReq
}

func ParseHttpRequest(requestBytes []byte) HttpRequest {
	req := NewEmptyHttpRequest()

	// seperate in reqHeader and reqBody
	blocks := bytes.SplitN(requestBytes, []byte(DOUBLE_CRLF), 2)
	// if len(blocks) >= 1 {
	// 	finalBlock := blocks[len(blocks)-1]
	// 	finalBlockLen := len(finalBlock)
	// 	if finalBlockLen > 2 {
	// 		if slices.Compare(finalBlock[finalBlockLen-2:], []byte(CRLF)) == 0 {
	// 			blocks[len(blocks)-1] = blocks[len(blocks)-1][:finalBlockLen-1-2]
	// 		}
	// 	}
	// }

	if len(blocks) >= 1 {
		linesStr := string(blocks[0])
		lines := strings.Split(linesStr, CRLF)

		if len(lines) == 0 {
			return req
		}

		reqLine := lines[0]
		words := strings.Split(reqLine, " ")
		switch len(words) {
		case 1:
			req.Method = words[0]
		case 2:
			req.Method = words[0]
			req.RequestURI = words[1]
		case 3:
			req.Method = words[0]
			req.RequestURI = words[1]
			req.HTTPVersion = words[2]
		}

		if len(lines) > 1 { // header is more than just request line
			for _, headerLine := range lines[1:] {
				words := strings.SplitN(headerLine, ": ", 2)
				if len(words) <= 1 {
					continue
				}
				req.Headers[textproto.CanonicalMIMEHeaderKey(words[0])] = words[1]
			}
		}
	}

	if len(blocks) >= 2 {
		req.Body = bytes.Join(blocks[1:], []byte(""))
	}

	contentLen, ok := req.Headers["Content-Length"]
	if ok {
		contentLenInt, err := strconv.Atoi(contentLen)
		if err != nil {
			slog.Warn("invalid Content-Length header")
		} else {
			req.Body = req.Body[:contentLenInt]
		}
	}

	return req
}
