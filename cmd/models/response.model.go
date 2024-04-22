package models

import (
	"net/textproto"
	"strconv"
	"strings"

	"github.com/patos-ufscar/http-web-server-example-go/utils"
)

type HttpResponse struct {
	HTTPVersion								string
	StatusCode								int
	StatusText								string
	Headers									map[string]string
	Body									[]byte
}

func NewHttpResponse() HttpResponse {
	var httpResp HttpResponse
	httpResp.Headers = make(map[string]string)
	return httpResp
}

func (r HttpResponse) DumpResponse() []byte {
	respLines := []string{}

	codeStr := strconv.Itoa(int(r.StatusCode))
	respLines = append(
		respLines, 
		r.HTTPVersion + utils.SP + codeStr + utils.SP + r.StatusText,
	)
	for k, v := range r.Headers {
		respLines = append(
			respLines, 
			textproto.CanonicalMIMEHeaderKey(k) + ":" + utils.SP + v,
		)
	}

	resp := []byte(strings.Join(respLines, utils.CLRF))
	resp = append(resp, []byte(utils.CLRF)...)

	contentLen := len(r.Body)
	if contentLen > 0 {
		r.Headers["Content-Lenght"] = strconv.Itoa(contentLen)
		resp = append(resp, []byte(utils.CLRF)...)
		resp = append(resp, r.Body...)
		resp = append(resp, []byte(utils.CLRF)...)
	}

	return resp
}
