// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func httpRequest() (string, error) {
	resp, err := http.Get("http://localhost:1234/brexitDate")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("brexit date is:" + string(body))

	request, err := http.NewRequest(http.MethodPut, "http://localhost:1234/brexitDate", strings.NewReader("29rd April"))
	if err != nil {
		return "", err
	}
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	fmt.Println("put succeeded")
	return "", nil
}

func netcatish() error {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		return err
	}

	for {
		// first write to conn
		inputReader := bufio.NewReader(os.Stdin)
		text, err := inputReader.ReadString('\n')
		if err != nil {
			return err
		}

		fmt.Print("client sending: " + text)
		_, err = conn.Write([]byte(text))
		if err != nil {
			return err
		}

		// second read from conn
		connReader := bufio.NewReader(conn)
		connText, err := connReader.ReadString(byte('\n'))
		if err != nil {
			return err
		}

		fmt.Print("client received: " + connText)
	}
}

func rawRequest() (string, error) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		return "", err
	}

	lineAndHeaders := `GET /brexitDate HTTP/1.1
Host: localhost:1234
Accept: */*`
	crlf := "\r\n\r\n"
	request := lineAndHeaders + crlf
	_, err = conn.Write([]byte(request))
	if err != nil {
		return "", err
	}

	response, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}

	body := strings.Split(string(response), crlf)[1]
	fmt.Print("brexit date is: " + body)
	conn.Close()

	conn, err = net.Dial("tcp", "localhost:1234")
	if err != nil {
		return "", err
	}

	lineAndHeaders = `PUT /brexitDate HTTP/1.1
Host: localhost:1234
Accept: */*`
	body = "neverrr!!!"
	request = lineAndHeaders + crlf + body
	_, err = conn.Write([]byte(request))
	if err != nil {
		return "", err
	}

	response, err = ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}
	fmt.Print("PUT succeeded")
	conn.Close()
	return "", nil
}

func main() {
	response, err := httpRequest()
	//err := netcatish()
	//response, err := rawRequest()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
