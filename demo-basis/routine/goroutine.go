package main

import (
	"fmt"
	"time"
)

// 协程会随着程序的结束而消亡
func main() {
	fmt.Println("in main")
	go longWait()
	go shortWait()
	fmt.Println("about to sleep in main")
	time.Sleep(10 * 1e9)
	fmt.Println("at the end of main()")
}

func longWait() {
	fmt.Println("beginning longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("end of longWait()")
}

func shortWait() {
	fmt.Println("beginning shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("end of shortWait()")
}
