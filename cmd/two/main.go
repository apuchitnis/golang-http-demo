// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"bufio"
	"net"
	"os"
)

func netcatish() {
	conn, _ := net.Dial("tcp", "localhost:1234")

	for {
		// first write to conn
		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		println("client sending: " + text)
		conn.Write([]byte(text))

		// second read from conn
		connReader := bufio.NewReader(conn)
		connText, _ := connReader.ReadString(byte('\n'))
		println("client received: " + connText)
	}
}

func main() {
	netcatish()
}
