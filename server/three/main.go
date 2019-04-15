// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package three

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

func handleConnection(conn net.Conn) error {
	scanner := bufio.NewScanner(conn)
	crlf := "\r\n\r\n"
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

	success := scanner.Scan()
	if !success {
		return scanner.Err()
	}

	request := scanner.Text()
	getRequest := "GET /brexitDate HTTP/1.1"
	putRequest := "PUT /brexitDate HTTP/1.1"
	switch {
	case strings.HasPrefix(request, getRequest):
		fmt.Println("GET request")
		lineAndHeaders := fmt.Sprintf(`HTTP/1.1 200 OK
		Content-Type: text/plain
		Content-Length: %v`, len(brexitDate))
		body := brexitDate + "\n"
		response := lineAndHeaders + crlf + body
		_, err := conn.Write([]byte(response))
		if err != nil {
			return err
		}
	case strings.HasPrefix(request, putRequest):
		fmt.Println("PUT request")
		body := strings.Split(request, crlf)[1]
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
