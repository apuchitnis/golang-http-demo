package main

import (
	"bufio"
	"net"
	"os"
)

func netcatish() {
	conn, _ := net.Dial("tcp", "localhost:1234")

	for {
		// First, read from STDIN and write to conn.
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		println("client sending:", text)
		conn.Write([]byte(text))

		// Second, read from conn and write to STDOUT.
		text, _ = bufio.NewReader(conn).ReadString(byte('\n'))
		println("client received:", text)
	}
}

func main() {
	netcatish()
}
