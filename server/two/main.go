// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"bufio"
	"net"
	"os"
)

func netcatish(conn net.Conn) {
	for {
		//first read from conn
		connReader := bufio.NewReader(conn)
		text, _ := connReader.ReadString(byte('\n'))
		println("server received: " + text)

		//second write to conn
		inputReader := bufio.NewReader(os.Stdin)
		text, _ = inputReader.ReadString('\n')

		println("server sending: " + text)
		_, _ = conn.Write([]byte(text))
	}
}

func main() {
	ln, _ := net.Listen("tcp", "localhost:1234")
	conn, _ := ln.Accept()
	netcatish(conn)
}
