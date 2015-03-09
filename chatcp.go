package main

import "fmt"
import "net"

func main() {
	l, err := net.Listen("tcp", ":5566")
	name_map := make(map[string]net.Conn)

	check_err(err)
	for {
		conn, err := l.Accept()
		check_err(err)
		go handle_conn(conn, name_map)
	}

}

func handle_conn(conn net.Conn, name_map map[string]net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			conn.Close()
		}
	}()
	fmt.Println("Got a conn", conn)
	name_slice := make([]byte, 22)
	talk_to_slice := make([]byte, 22)

	n, err := conn.Write([]byte("Who are you?\n"))
	check_err(err)
	n, err = conn.Read(name_slice)
	check_err(err)

	name := string(name_slice[:n-2])
	name_map[name] = conn

	n, err = conn.Write([]byte("Who would you like to talk to\n?"))
	check_err(err)
	n, err = conn.Read(talk_to_slice)
	check_err(err)
	talk_to := string(talk_to_slice[:n-2])

	send_conn := name_map[talk_to]

	recv_buff := make([]byte, 22)
	for {
		_, err = conn.Read(recv_buff)
		check_err(err)
		_, err = send_conn.Write(recv_buff)
		check_err(err)
		clear_bs(recv_buff)

	}

	err = conn.Close()
	check_err(err)
}

func check_err(err error, extra_context ...string) {
	if err != nil {
		panic(fmt.Sprintf("Error: %v, Context: %v", err, extra_context))
	}
}

func clear_bs(bs []byte) {
	for i := 0; i < len(bs); i++ {
		bs[i] = 0
	}
}
