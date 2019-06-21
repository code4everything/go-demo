package main

import (
	"fmt"
)

func main() {
	out := make(chan int)
	// why deadlock?
	// while you write a value to the chanel(no buffer), it will block current routine,
	// so the following codes will not execute until the chanel become available,
	// but in this case, the chanel has no chance to read,obviously deadlock happened.
	out <- 1
	go read(out)
}

func read(ch chan int) {
	fmt.Println(<-ch)
}
