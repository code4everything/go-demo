package main

import "fmt"

var arr = make([]int, 100)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go routine1(ch1)
	go routine2(ch1, ch2)
	// waiting for complete
	<-ch2
	for i := 0; i < 100; i++ {
		fmt.Println(arr[i])
	}
}

func routine1(ch chan int) {
	for i := 1; i < 100; i += 2 {
		arr[i-1] = i
		ch <- i
	}
	close(ch)
}

func routine2(ch1, ch2 chan int) {
	for v := range ch1 {
		arr[v] = v + 1
	}
	// completed flag
	ch2 <- 1
}
