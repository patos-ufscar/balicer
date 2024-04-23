package handlers

import (
	"net/http"
	"net/textproto"
	"strings"

	"github.com/patos-ufscar/http-web-server-example-go/models"
)

type HandlerStaticImpl struct {
	Path					string
	StatusCode				int
	Headers					map[string]string
	Body					[]byte
}

func NewHandlerStaticImpl(path string, ret models.ReturnStatic) Handler {
	// headers := make(map[string]string)
	return &HandlerStaticImpl{
		Path: path,
		StatusCode: ret.Code,
		Headers: ret.Headers,
		Body: ret.Body,
	}
}

func (h *HandlerStaticImpl) ValidPath(host string) bool {

	if strings.HasPrefix(host, h.Path) {
		return true
	}

	return true
}

func (h *HandlerStaticImpl) Handle(req models.HttpRequest) (models.HttpResponse, error) {

	resp := models.NewHttpResponse()

	for k, v := range h.Headers {
		resp.Headers[textproto.CanonicalMIMEHeaderKey(k)] = v
	}

	resp.StatusCode = h.StatusCode
	resp.StatusText = http.StatusText(h.StatusCode)
	resp.HTTPVersion = "HTTP/1.1"
	resp.Body = h.Body

	return resp, nil
}
