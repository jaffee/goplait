package main

import (
	"fmt"
	"net"
	"time"
)

type command struct {
	command string
	name    string
	conn    net.Conn
	reply   chan command
}

type futureConn struct {
	messages [][]byte
}

func (f *futureConn) Write(b []byte) (n int, err error) {
	f.messages = f.messages.append(b)
	return len(b), nil
}

func (f *futureConn) Read(b []byte) (n int, err error) {
	return nil, nil
}

func (f *futureConn) Close() error {
	return nil
}

func (f *futureConn) RemoteAddr() Addr {
	return nil
}

func (f *futureConn) RemoteAddr() Addr {
	return nil
}

func (f *futureConn) SetDeadline(t time.Time) error {
	return nil
}

func (f *futureConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (f *futureConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":5566")
	registerChan := make(chan command, 20)
	go registrationHandler(registerChan)

	checkErrDie(err)
	for {
		if conn, err := l.Accept(); checkConnErr(err, conn) {
			go handleConn(conn, registerChan)
		}

	}
}

func handleConn(conn net.Conn, regChan chan command) {
	defer func() {
		recover()
	}()

	ttconn := handleWelcome(conn, regChan)
	buff := make([]byte, 100)
	for {
		read(buff, conn)
		write(buff, ttconn)
		clearBytes(buff)
	}
	fmt.Println("Connection successful")
	conn.Close()
}

func clearBytes(bs []byte) {
	for i := 0; i < len(bs); i++ {
		bs[i] = 0
	}
}

func handleWelcome(conn net.Conn, regChan chan command) net.Conn {
	name := input("What's your name?", conn)

	regChan <- command{"register", name, conn, nil}
	talk_to := input("Who do you want to talk to?", conn)
	reply_chan := make(chan command)
	regChan <- command{"get", talk_to, conn, reply_chan}
	resp := <-reply_chan
	ttconn := resp.conn

	return ttconn
}

func read(buff []byte, conn net.Conn) string {
	n, err := conn.Read(buff)
	checkErrDie(err)
	if n > 2 {
		return string(buff[:n-2])
	} else {
		return ""
	}

}

func write(buff []byte, conn net.Conn) {
	_, err := conn.Write(buff)
	checkErrDie(err)
}

func input(request string, conn net.Conn) string {
	_, err := conn.Write([]byte(request + "\n"))
	checkErrDie(err)
	resp := make([]byte, 50)
	n, err := conn.Read(resp)
	if n > 2 {
		return string(resp[:n-2])
	} else {
		return ""
	}
}

// func registerNewConn(name string, conn net.Conn, regChan chan command) {
// 	regChan <- command{name, conn}
// }

func registrationHandler(regChan chan command) {
	registeredConns := make(map[string]net.Conn)
	for {
		newreg := <-regChan
		if newreg.command == "register" {
			conn, exists = registeredConns[newreg.name]
			if exists {
				_, ok := conn.(futureConn)
				if ok {
					for b := range conn.messages {
						write(b, newreg.conn)
					}
					registeredConns[newreg.name] = newreg.conn
				} else {
					panic(nil)
				}
			} else {
				registeredConns[newreg.name] = newreg.conn
			}
		} else if newreg.command == "get" {
			reply_conn, exists := registeredConns[newreg.name]
			if !exists {
				registeredConns[newreg.name] = futureConn{make([][]byte, 5)}
			}
			newreg.reply <- command{"reply", "", reply_conn, nil}
		}
	}
}

func checkErrDie(err error) {
	if err != nil {
		fmt.Printf("Got an error: %v\n Panic and die.", err)
		panic(err)
	}
}

func checkConnErr(err error, context interface{}) bool {
	if err != nil {
		fmt.Printf("Got an error: %v, context: %v\n", err, context)
		return false
	}
	return true
}
