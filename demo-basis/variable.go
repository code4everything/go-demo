package main

import (
	"fmt"
	"strings"
	"time"
)

// Go 中不允许不同类型之间的混合使用，但对常量的限制较少，因此常量可以混合使用
// 因为 Go 不能够隐式转换，所以我们必须显式的转换变量类型
func main() {
	i := 1
	// 值类型
	x, y := i, i
	x++
	fmt.Println("x equals y:", x == y)

	// 引用类型
	m, n := &i, &i
	*m++
	fmt.Println("m equals n:", *m == *n)

	// 常量与 iota，每当 iota 遇到 const 时都会重置为零
	const (
		// 忽略零值
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		// ZB
		// YB
	)

	fmt.Println(KB, MB, GB, TB, PB, EB)

	// 仅仅是测试
	time.Sleep(1000)

	// 字符串，一个 ascii 码占一个字节，其他字符占 2~4 个字节
	en := "ab"
	zh := "中文"
	fmt.Println("char length equals:", strings.Count(en, "") == strings.Count(zh, ""))
	// len 函数其实获取的是 byte 数
	fmt.Println("byte length equals:", len(en) == len(zh))
}
