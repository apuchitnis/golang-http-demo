// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func netcatish(conn net.Conn) error {
	for {
		//simple line by line scanner

		//first read from conn
		connReader := bufio.NewReader(conn)
		text, err := connReader.ReadString(byte('\n'))
		if err != nil {
			return err
		}
		println("server received: " + text)

		//second write to conn
		inputReader := bufio.NewReader(os.Stdin)
		text, err = inputReader.ReadString('\n')
		if err != nil {
			return err
		}

		println("server sending: " + text)
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
