package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

var (
	usernames = "leo,nev,sam"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//log.Println(fmt.Sprintf("server handling request: %v", r))
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET request")
		w.Write([]byte(fmt.Sprintf("usernames are: %v\n", usernames)))
	case http.MethodPut:
		fmt.Println("PUT request")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		usernames = string(body)
	default:
		w.WriteHeader(501)
	}
}

func main() {
	// Complex golang server example
	// http.HandleFunc("/username", handler)
	// log.Fatal(http.ListenAndServe(":1234", nil))

	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("err")
		}
		go func() {
			defer conn.Close()
			err := handleConnection(conn)
			if err != nil {
				log.Println(err)
			}
		}()
	}
}

func handleConnection(conn net.Conn) error {
	scanner := bufio.NewScanner(conn)
	crlf := "\r\n\r\n"
	scanner.Split(
		func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if strings.Contains(string(data), crlf) {
				log.Println("found end of HTTP request!")
				return 0, data, nil
			}
			if atEOF {
				log.Print(string(data))
				return 0, data, errors.New("unable to get more data, failed to find doubleNewLine, eg if Read call on the reader fails")
			}
			// https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4			return 0, nil, nil
			// token not found yet, keep scanning until we find a token. this return asks for larger input
			return 0, nil, nil
		})

	success := scanner.Scan()
	if !success {
		return scanner.Err()
	}
	// simple line by line scanner
	//		connReader := bufio.NewReader(conn)
	//		connText, _ := connReader.ReadString(byte('\r\n'))
	//		log.Println("received: " + connText)
	//fmt.Println("http request:")
	//fmt.Println(scanner.Text())
	//inputReader := bufio.NewReader(os.Stdin)
	//text, _ := inputReader.ReadString('\n')

	//rd := bufio.NewReader(strings.NewReader(scanner.Text()))
	//request, err := http.ReadRequest(rd)
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Println(request.Method)

	getRequest := "GET /username HTTP/1.1"
	putRequest := "PUT /username HTTP/1.1"
	switch {
	case strings.HasPrefix(scanner.Text(), getRequest):
		fmt.Println("GET request")
		lineAndHeaders := fmt.Sprintf(`HTTP/1.1 200 OK
Content-Type: text/plain
Content-Length: %v`, len(usernames)+15)
		body := "usernames are " + usernames + "\n"
		response := lineAndHeaders + crlf + body
		_, err := conn.Write([]byte(response))
		if err != nil {
			return err
		}
	case strings.HasPrefix(scanner.Text(), putRequest):
		fmt.Println("PUT request")
		body := strings.Split(scanner.Text(), crlf)[1]
		usernames = body
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
	return nil
}
