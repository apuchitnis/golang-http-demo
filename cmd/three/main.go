// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package three

import (
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func rawRequest() error {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		return err
	}

	lineAndHeaders := `GET /brexitDate HTTP/1.1
Host: localhost:1234
Accept: */*`
	crlf := "\r\n\r\n"
	request := lineAndHeaders + crlf
	_, err = conn.Write([]byte(request))
	if err != nil {
		return err
	}

	response, err := ioutil.ReadAll(conn)
	if err != nil {
		return err
	}

	body := strings.Split(string(response), crlf)[1]
	println("brexit date is: " + body)
	conn.Close()

	conn, err = net.Dial("tcp", "localhost:1234")
	if err != nil {
		return err
	}

	lineAndHeaders = `PUT /brexitDate HTTP/1.1
Host: localhost:1234
Accept: */*`
	body = "neverrr!!!"
	request = lineAndHeaders + crlf + body
	_, err = conn.Write([]byte(request))
	if err != nil {
		return err
	}

	response, err = ioutil.ReadAll(conn)
	if err != nil {
		return err
	}
	println("PUT succeeded")
	conn.Close()
	return nil
}

func main() {
	if err := rawRequest(); err != nil {
		log.Fatal(err)
	}
}
