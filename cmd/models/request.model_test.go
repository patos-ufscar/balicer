package models_test

import (
	"reflect"
	"testing"

	"github.com/patos-ufscar/http-web-server-example-go/models"
)

func TestParseHttpRequest(t *testing.T) {
	reqLineOnly := models.NewEmptyHttpRequest()
	reqLineOnly.HTTPVersion = "HTTP/1.1"
	reqLineOnly.RequestURI = "/index.html"
	reqLineOnly.Method = "GET"

	withHeaders := models.NewEmptyHttpRequest()
	withHeaders.HTTPVersion = "HTTP/1.1"
	withHeaders.RequestURI = "/index.html"
	withHeaders.Method = "GET"
	withHeaders.Headers["Host"] = "localhost:8080"

	withBody := models.NewEmptyHttpRequest()
	withBody.HTTPVersion = "HTTP/1.1"
	withBody.RequestURI = "/index.html"
	withBody.Method = "GET"
	withBody.Headers["Host"] = "localhost:8080"
	withBody.Body = []byte("example-body")

	withContentLen := models.NewEmptyHttpRequest()
	withContentLen.HTTPVersion = "HTTP/1.1"
	withContentLen.RequestURI = "/index.html"
	withContentLen.Method = "GET"
	withContentLen.Headers["Host"] = "localhost:8080"
	withContentLen.Headers["Content-Length"] = "16"
	withContentLen.Body = []byte("example-body-len")

	type args struct {
		requestBytes []byte
	}
	tests := []struct {
		name string
		args args
		want models.HttpRequest
	}{
		{
			name: "empty request",
			args: args{requestBytes: []byte{}},
			want: models.NewEmptyHttpRequest(),
		},
		{
			name: "request line only",
			args: args{requestBytes: []byte("GET /index.html HTTP/1.1\r\n\r\n")},
			want: reqLineOnly,
		},
		{
			name: "with headers",
			args: args{requestBytes: []byte("GET /index.html HTTP/1.1\r\nHost: localhost:8080\r\n\r\n")},
			want: withHeaders,
		},
		{
			name: "with body",
			args: args{requestBytes: []byte("GET /index.html HTTP/1.1\r\nHost: localhost:8080\r\n\r\nexample-body\r\n\r\n")},
			want: withBody,
		},
		{
			name: "with contentLen",
			args: args{requestBytes: []byte("GET /index.html HTTP/1.1\r\nHost: localhost:8080\r\nContent-Length: 16\r\n\r\nexample-body-lenCUT\r\n\r\n")},
			want: withContentLen,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := models.ParseHttpRequest(tt.args.requestBytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHttpRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
