package models

type HttpRequest struct {
	Method									string
	RequestURI								string
	HTTPVersion								string
	Host									string
	Headers									map[string]string
	Body									[]byte
}

func NewHttpRequest() HttpRequest {
	var httpReq HttpRequest
	httpReq.Headers = make(map[string]string)
	return httpReq
}
