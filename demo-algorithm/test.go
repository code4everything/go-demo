package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// test convert byte to int
	bs := []byte("abcdefghijklmnopqrstuvwxyz")
	for _, b := range bs {
		fmt.Print(strconv.Itoa(int(b)) + " ")
	}
	fmt.Println()
	// test convert int to byte
	for i := 33; i < math.MaxInt8; i++ {
		fmt.Print(string(byte(i)) + " ")
	}
}
