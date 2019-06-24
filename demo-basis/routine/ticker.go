package main

import (
	"fmt"
	"time"
)

func main() {
	// 限制处理频率
	rateChan := time.Tick(1e9)
	for i := 0; i < 10; i++ {
		<-rateChan
		fmt.Println("time:", i)
	}

	// 与上面等价
	ticker := time.NewTicker(1e9)
	for i := 0; i < 10; i++ {
		<-ticker.C
		fmt.Println("time:", i)
	}

	// 只会写一次
	boom := time.After(5e9)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom!")
			return
		}
	}
}
