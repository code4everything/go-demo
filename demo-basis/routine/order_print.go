package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)
	go routine1(ch1)
	go routine2(ch1, ch2)
	// waiting for complete
	<-ch2
}

func routine1(ch chan int) {
	for i := 1; i < 100; i += 2 {
		ch <- i
		fmt.Println("one:", i)
	}
	close(ch)
}

func routine2(ch1, ch2 chan int) {
	for v := range ch1 {
		fmt.Println("two:", v+1)
	}
	// completed flag
	ch2 <- 1
}
