package main

import (
	"bufio"
	"fmt"
	"os"
)

// see details:
// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.1.md
func main() {
	var (
		name   string
		result = "string scanner: 2145"
		format = "string scanner: %s"
	)
	fmt.Println("please input your name:")
	// 从控制台中输入
	_, err := fmt.Scanln(&name)
	if err == nil {
		fmt.Printf("good evening %s ", name)
	}

	// 字符串中输入
	_, err = fmt.Sscanf(result, format, &name)
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err.Error())
	}

	// 控制台缓冲读取
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("please enter something for test: ")
	// like readLine
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("your input: %s", input)
	}
	fmt.Println("programme is end")
}
