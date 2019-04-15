// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package two

import (
	"bufio"
	"log"
	"net"
	"os"
)

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

		println("client sending: " + text)
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

		println("client received: " + connText)
	}
}

func main() {
	if err := netcatish(); err != nil {
		log.Fatal(err)
	}
}
