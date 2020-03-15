package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello world"

	// 是否包含
	strings.Contains(str, "w")

	// 索引
	strings.Index(str, "d")

	str = "1#2#345"

	// 切割字符串
	separatedStr := strings.Split(str, "#")

	// 合并字符串
	strings.Join(separatedStr, "#")

	// 是否包含前缀
	strings.HasPrefix(str, "@")

	// 整型转字符串
	strconv.Itoa(32)

	// 字符串转整型
	i, _ := strconv.Atoi("23")
	fmt.Print(i)

}
