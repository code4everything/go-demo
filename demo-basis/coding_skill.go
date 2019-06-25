package main

import (
	"fmt"
	"strings"
)

func main() {
	// =======================================字符串=====================================================================

	// 1.修改字符串中的一个字符
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	fmt.Println(string(c))
	// 2. 获取子串
	fmt.Println(str[1:])
	// 3. 连接字符串
	fmt.Println(strings.Join([]string{"china", "england", "america"}, ", "))

	// ===============================================协程与通道=========================================================

	// 1. 检测通道是否已关闭
	ch := make(chan int)
	waiting := make(chan interface{})
	go func() {
		if _, open := <-ch; !open {
			fmt.Println("channel is closed")
		}
		waiting <- 0
	}()
	close(ch)
	<-waiting
}
