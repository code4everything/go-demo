package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	hash := sha1.New()
	_, _ = io.WriteString(hash, "hash test")
	b := make([]byte, 0, 1024)
	fmt.Printf("result: %x\n", hash.Sum(b))
	fmt.Printf("result: %b\n", hash.Sum(b))

	hash.Reset()
	data := []byte("we shall overcome")
	_, err := hash.Write(data)
	if err == nil {
		fmt.Printf("result: %x\n", hash.Sum(b))
	} else {
		fmt.Println("has write error!")
	}
}
