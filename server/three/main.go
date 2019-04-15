// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

)

var (
	brexitDate = "29th March"
)

const (
	crlf = "\r\n\r\n"
)

func handleConnection(conn net.Conn) {
	requestMethod, URL, body := getRequestParameters(conn)

	println("server handling request: ", requestMethod+" "+URL)

	switch requestMethod {
	case "GET":
		lineAndHeaders := fmt.Sprintf(`HTTP/1.1 200 OK
			Content-Type: text/plain
			Content-Length: %v`, len(brexitDate))
		body := brexitDate + "\n"
		response := lineAndHeaders + crlf + body
		_, _ = conn.Write([]byte(response))
	case "PUT":
		brexitDate = body
		response := "HTTP/1.1 200 OK" + crlf
		_, _ = conn.Write([]byte(response))
	default:
		_, _ = conn.Write([]byte("HTTP/1.1 501 Unimplemented\n"))
	}
	conn.Close()
}

func getRequestParameters(conn net.Conn) (method string, URL string, body string) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(
		func(data []byte, atEOF bool) (advance int, token []byte, err error) {

			if strings.Contains(string(data), crlf) {
				return 0, data, nil
			}
			if atEOF {
				log.Print(string(data))
				return 0, data, errors.New("unable to get more data, failed to find doubleNewLine, eg if Read call on the reader fails")
			}
			return 0, nil, nil
		})
	_ = scanner.Scan()
	request := scanner.Text()
	statusLine := strings.Split(request, "\n")[0]
	method = strings.Split(statusLine, " ")[0]
	URL = strings.Split(statusLine, " ")[1]
	body = strings.Split(request, crlf)[1]
	return method, URL, body
}

func main() {
	ln, _ := net.Listen("tcp", "localhost:1234")

	for {
		conn, _ := ln.Accept()
		defer conn.Close()
		handleConnection(conn)
	}
}
