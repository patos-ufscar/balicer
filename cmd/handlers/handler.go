package handlers

import (
	"fmt"

	"github.com/patos-ufscar/http-web-server-example-go/cli"
	"github.com/patos-ufscar/http-web-server-example-go/models"
)

type Handler interface {
	ValidPath(host string)						bool
	Handle(req models.HttpRequest)				(models.HttpResponse, error)
}

func HandlerFactory(locConf models.HandlerConfig) (Handler, error) {
	switch locConf.ReturnType {
	case "static":
		ret, err := cli.ParseStaticReturn(locConf.Return)
		if err != nil {
			return nil, err
		}
		return NewHandlerStaticImpl(
			locConf.Path,
			*ret,
		), nil
	}

	// return nil, errors.New("not yet implemented")
	return nil, fmt.Errorf("not yet implemented handler: %s", locConf.ReturnType)
}
