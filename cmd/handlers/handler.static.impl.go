package handlers

import (
	"net/http"
	"net/textproto"
	"strings"

	"github.com/patos-ufscar/http-web-server-example-go/models"
)

type HandlerStaticImpl struct {
	config				models.HandlerConfig
}

func NewHandlerStaticImpl(config models.HandlerConfig) Handler {
	headers := make(map[string]string)
	headers["Content-Type"] = "text/html"
	headers["Server"] = "balicer"

	return &HandlerStaticImpl{
		config: models.HandlerConfig{
			Path: "/",
			ReturnType: "static",
			Return: models.ReturnConfig{
				Code: 200,
				Headers: headers,
				Body: []byte("<h1>quack!</h1>"),
			},
		},
	}
}

func (h *HandlerStaticImpl) ValidPath(host string) bool {

	if strings.HasPrefix(host, h.config.Path) {
		return true
	}

	return true
}

func (h *HandlerStaticImpl) Handle(req models.HttpRequest) (models.HttpResponse, error) {

	resp := models.NewHttpResponse()

	for k, v := range h.config.Return.Headers {
		resp.Headers[textproto.CanonicalMIMEHeaderKey(k)] = v
	}

	resp.StatusCode = h.config.Return.Code
	resp.StatusText = http.StatusText(h.config.Return.Code)
	resp.HTTPVersion = "HTTP/1.1"
	resp.Body = h.config.Return.Body

	return resp, nil
}
