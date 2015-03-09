package main

import (
	"fmt"
	"net"
)

func main() {
	l, _ := net.Listen("tcp", ":5252")
	for {
		conn, _ := l.Accept()
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Printf("local:  %v\n", conn.LocalAddr())
	fmt.Printf("remote: %v\n\n", conn.RemoteAddr())
}
