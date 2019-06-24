package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("dialing error", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("first, what is your name?")
	name, err := inputReader.ReadString('\n')
	name = strings.Trim(name, "\r\n")
	for {
		fmt.Println("what to send to the server? type q to quit.")
		input, _ := inputReader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		if input == "q" || input == "Q" {
			return
		}
		_, err = conn.Write([]byte(name + " says: " + input))
	}
}
