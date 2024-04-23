package handlers

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/patos-ufscar/http-web-server-example-go/utils"
)


func HandleGlobal(conn net.Conn, directory string) ([]byte, error) {
	readBuffer := make([]byte, 8 * (1 << 10))
	// readBuffer := [8 * (1 << 10)]byte{}
	_, err := conn.Read(readBuffer)
	if err != nil {
		slog.Error(fmt.Sprintf("Error reading request: %s", err.Error()))
		os.Exit(1)
	}

	strBody := strings.Split(string(readBuffer), "\r\n")

	request := [][]string{}
	for _, v := range strBody {
		words := strings.Split(v, " ")
		request = append(request, words)
	}

	// fmt.Println(request)

	path := request[0][1]

	if path == "/" {
		err := utils.ReplyHTTP(conn, []byte("HTTP/1.1 200 OK\r\n\r\n"))
		return nil, err
	}

	if strings.HasPrefix(path, "/echo/") {
		// fmt.Printf("'%s'\n", path[len("/echo/"):])
		err := utils.ReplyString(conn, path[6:])
		return nil, err
	}

	if strings.HasPrefix(path, "/user-agent") {
		err := utils.ReplyString(conn, request[2][1])
		return nil, err
	}

	if strings.HasPrefix(path, "/files/") {
		fileName := path[len("/files/"):]
		filePath := fmt.Sprintf("%s%s", directory, fileName)

		if request[0][0] == "GET" {
			content, err := os.ReadFile(filePath)
			if err != nil {
				err = utils.ReplyHTTP(conn, []byte("HTTP/1.1 404 NOT FOUND\r\n\r\n"))
				return nil, err
			}
			err = utils.ReplyOctetStream(conn, string(content))
			return nil, err
		}

		if request[0][0] == "POST" {
			fmt.Println(strBody)
			contentStr := strBody[len(strBody)-1]
			fmt.Println(contentStr)
			contentLen, err := strconv.Atoi(request[3][1])
			if err != nil {
				fmt.Println("AAAAA", request)
				return nil, err
			}
			contentStr = contentStr[0:contentLen]
			fmt.Println(contentStr)
			err = os.WriteFile(filePath, []byte(contentStr), 0644)
			if err != nil {
				fmt.Println("Error writing file: ", err.Error())
				return nil, err
			}
			err = utils.ReplyHTTP(conn, []byte("HTTP/1.1 201 Created\r\n\r\n"))
			return nil, err
		}
	}

	err = utils.ReplyHTTP(conn, []byte("HTTP/1.1 404 NOT FOUND\r\n\r\n"))
	
	return nil, err
}