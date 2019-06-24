package main

import "fmt"

var yield chan int

var resume chan int

var value int

func init() {
	yield = make(chan int)
	resume = make(chan int)
	value = 0

	go func() {
		for {
			<-resume
			value += 2
			yield <- value
		}
	}()
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("lazy gen:", evaluate())
	}
}

func evaluate() int {
	resume <- 1
	return <-yield
}
