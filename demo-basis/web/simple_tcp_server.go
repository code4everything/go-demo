package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting server...")
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("listening error", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accepting error", err.Error())
			return
		}
		doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println("reading error", err.Error())
			return
		}
		fmt.Printf("received data:%v\n", string(buf[:size]))
	}
}
