package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	age := rand.Intn(100)
	fmt.Println("age:", age)

	if age > 60 {
		fmt.Println("people retired")
	} else if age > 22 {
		fmt.Println("people working")
	} else if age > 6 {
		fmt.Println("people learning")
	} else {
		fmt.Println("people bear and growing")
	}

	// 带初始化语句的 if
	if r := rand.Intn(2); r == 1 {
		fmt.Println("light:", r)
	} else {
		fmt.Println("dark:", r)
	}

	switch runtime.GOOS {
	case "windows", "linux":
		fmt.Println(runtime.GOOS)
	case "darwin":
		fmt.Println("mac osx")
	default:
		fmt.Println("unknown system")
	}

	sum := 0
	var i uint32 = 0
	for ; i < 5; i++ {
		sum += 1 << i
	}
	fmt.Println(sum)

	// while style
	str := "▄"
	for sum > 0 {
		sum >>= 1
		fmt.Println(strings.Repeat(str, rand.Intn(100)))
	}

	// range style: foreach list, map etc.
	str = "hello go"
	for _, c := range str {
		fmt.Print(string(c))
	}

	// about goto, we strongly recommend you don't use it, detail see：
	// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/05.6.md
}
