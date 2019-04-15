// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"io/ioutil"
	"net"
	"strings"
)

const (
	crlf = "\r\n\r\n"
)

func rawRequest() {
	conn, _ := net.Dial("tcp", "localhost:1234")

	lineAndHeaders := `GET /brexitDate HTTP/1.1
Host: localhost:1234
Accept: */*`

	request := lineAndHeaders + crlf
	_, _ = conn.Write([]byte(request))

	response, _ := ioutil.ReadAll(conn)

	body := strings.Split(string(response), crlf)[1]
	println("brexit date is: " + body)
	conn.Close()

	conn, _ = net.Dial("tcp", "localhost:1234")

	lineAndHeaders = `PUT /brexitDate HTTP/1.1
Host: localhost:1234
Accept: */*`
	body = "NEVERRR!!! 🙅"
	request = lineAndHeaders + crlf + body
	_, _ = conn.Write([]byte(request))

	response, _ = ioutil.ReadAll(conn)
	println("PUT succeeded")
	conn.Close()
}

func main() {
	rawRequest()
}
