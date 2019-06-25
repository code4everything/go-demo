package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		read = true
		data = make([]uint8, 4096)
	)
	conn, err := net.Dial("tcp", "www.apache.org:80")
	if err == nil {
		io.WriteString(conn, "GET / \n")
		for read {
			count, err := conn.Read(data)
			read = err == nil
			fmt.Println(string(data[:count]))
		}
		conn.Close()
	}
}
