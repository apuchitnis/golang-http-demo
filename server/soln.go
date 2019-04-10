// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

// 1

/*func brexitDateHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Println("server handling request: ", r.Method)
		switch r.Method {
		case http.MethodGet:
				w.Write([]byte(brexitDate))
		case http.MethodPut:
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						return
				}
				brexitDate = string(body)
		default:
				w.WriteHeader(http.StatusNotImplemented)
		}
}

func main() {
		// net/http golang server example
		http.HandleFunc("/brexitDate", brexitDateHandler)
		log.Fatal(http.ListenAndServe("localhost:1234", nil))
}*/

// 2

/*func netcatish(conn net.Conn) error {
		for {
				//simple line by line scanner

				//first read from conn
				connReader := bufio.NewReader(conn)
				text, err := connReader.ReadString(byte('\n'))
				if err != nil {
						return err
				}
				fmt.Print("server received: " + text)

				//second write to conn
				inputReader := bufio.NewReader(os.Stdin)
				text, err = inputReader.ReadString('\n')
				if err != nil {
						return err
				}

				fmt.Print("server sending: " + text)
				_, err = conn.Write([]byte(text))
				if err != nil {
						return err
				}
		}
}


func main() {
	ln, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Println("err")
	}
	err = netcatish(conn)
	if err != nil {
		log.Println(err)
	}
}
*/

// 3

/*func handleConnection(conn net.Conn) error {
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
		go func() {
			defer conn.Close()
			err := handleConnection(conn)
			if err != nil {
				log.Println(err)
			}
		}()
	}
}


*/
