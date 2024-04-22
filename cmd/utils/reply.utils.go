package utils

import (
	"fmt"
	"net"
	"strconv"
)

func ReplyHTTP(conn net.Conn, rep []byte) error {
	// httpResponse := "HTTP/1.1 200 OK\r\n\r\n"
	nBytesSent, err := conn.Write(rep)
	if err != nil {
		return err
	}

	fmt.Printf("Sent %d bytes to client (expected: %d)\n", nBytesSent, len(rep))

	return nil
}

func Reply502(conn net.Conn) error {
	return ReplyHTTP(conn, []byte("HTTP/1.1 502 BAD GATEWAY\r\n"))
}

func ReplyString(conn net.Conn, msg string) error {
	response := "HTTP/1.1 200 OK\r\n"
	response += "Content-Type: text/plain\r\n"
	response += "Content-Length:" + strconv.Itoa(len(msg)) + "\r\n"

	response += "\r\n"
	response += msg + "\r\n"
	response += "\r\n"

	return ReplyHTTP(conn, []byte(response))
}

// func ReplyResponse(conn net.Conn, resp models.HttpResponse) error {
	
// }

func ReplyOctetStream(conn net.Conn, msg string) error {
	response := "HTTP/1.1 200 OK\r\n"
	response += "Content-Type: application/octet-stream\r\n"
	response += "Content-Length:" + strconv.Itoa(len(msg)) + "\r\n"

	response += "\r\n"
	response += msg + "\r\n"
	response += "\r\n"

	return ReplyHTTP(conn, []byte(response))
}