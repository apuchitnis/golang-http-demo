package main

import (
	"bufio"
	"net"
	"os"
)

func netcatish(conn net.Conn) {
	for {
		// First, read from conn and write to STDOUT.
		text, _ := bufio.NewReader(conn).ReadString(byte('\n'))
		println("server received:", text)

		// Second, read from STDIN and write to conn.
		text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		println("server sending:", text)
		conn.Write([]byte(text))
	}
}

func main() {
	ln, _ := net.Listen("tcp", "localhost:1234")
	conn, _ := ln.Accept()
	netcatish(conn)
}
