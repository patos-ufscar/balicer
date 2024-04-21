package handlers

import (
	"net"

	"github.com/patos-ufscar/http-web-server-example-go/models"
	"github.com/patos-ufscar/http-web-server-example-go/utils"
)

type HandlerStaticImpl struct {
	// config				models.HandlerConfig
}

// func NewHandlerStaticImpl(config models.LocationConfig) Handler {
func NewHandlerStaticImpl() Handler {
	return &HandlerStaticImpl{
		// config: config,
	}
}

func (h *HandlerStaticImpl) ValidHost(host string) bool {

	return true
	// for _, v := range h.config.HostsRegs {
	// 	match := v.FindString(host)
	// 	if match != "" {
	// 		return true
	// 	}
	// }

	// return false
}

func (h *HandlerStaticImpl) Handle(conn net.Conn, req models.HttpRequestFrame) error {

	return utils.ReplyString(conn, "OK")

	// return nil
}
