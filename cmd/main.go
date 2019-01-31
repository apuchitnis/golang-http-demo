package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

func httpRequest() (string, error) {
	resp, err := http.Get("http://localhost:1234/username")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))

	request, err := http.NewRequest(http.MethodPut, "http://localhost:1234/username", strings.NewReader("asdfasdf"))
	if err != nil {
		return "", err
	}
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))
	return "", nil
}

func rawRequest() (string, error) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		return "", err
	}

	lineAndHeaders := `GET /username HTTP/1.1
Host: localhost:1234
User-Agent: curl/7.63.0
Accept: */*`
	clrf := "\r\n\r\n"
	request := lineAndHeaders + clrf
	_, err = conn.Write([]byte(request))
	if err != nil {
		return "", err
	}

	response, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}

	body := strings.Split(string(response), clrf)[1]
	fmt.Print(body)
	conn.Close()

	conn, err = net.Dial("tcp", "localhost:1234")
	if err != nil {
		return "", err
	}

	lineAndHeaders = `PUT /username HTTP/1.1
Host: localhost:1234
User-Agent: curl/7.63.0
Accept: */*`
	body = "newone"
	request = lineAndHeaders + clrf + body
	_, err = conn.Write([]byte(request))
	if err != nil {
		return "", err
	}

	response, err = ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}
	conn.Close()
	return "", nil
}

func main() {
	//response, err := httpRequest()
	_, err := rawRequest()
	if err != nil {
		log.Fatal(err)
	}
}
