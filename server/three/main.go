// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/pkg/errors"
)

var (
	brexitDate = "29th March"
)

const (
	crlf = "\r\n\r\n"
)

func handleConnection(conn net.Conn) error {
	requestMethod, URL, body := getRequestParameters(conn)

	println("server handling request: ", requestMethod+" "+URL)

	switch requestMethod {
	case "GET":
		lineAndHeaders := fmt.Sprintf(`HTTP/1.1 200 OK
			Content-Type: text/plain
			Content-Length: %v`, len(brexitDate))
		body := brexitDate + "\n"
		response := lineAndHeaders + crlf + body
		_, err := conn.Write([]byte(response))
		if err != nil {
			return err
		}
	case "PUT":
		brexitDate = body
		response := "HTTP/1.1 200 OK" + crlf
		_, err := conn.Write([]byte(response))
		if err != nil {
			return err
		}
	default:
		_, err := conn.Write([]byte("HTTP/1.1 501 Unimplemented\n"))
		if err != nil {
			return err
		}
	}
	conn.Close()
	return nil
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
	//if !success {
	//	return scanner.Err()
	//}
	request := scanner.Text()
	statusLine := strings.Split(request, "\n")[0]
	method = strings.Split(statusLine, " ")[0]
	URL = strings.Split(statusLine, " ")[1]
	body = strings.Split(request, crlf)[1]
	return method, URL, body
}

func main() {
	ln, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("err")
		}
		defer conn.Close()
		if err = handleConnection(conn); err != nil {
			log.Println(err)
		}
	}
}
