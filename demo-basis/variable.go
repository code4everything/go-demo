package main

import "fmt"

// Go 中不允许不同类型之间的混合使用，但对常量的限制较少，因此常量可以混合使用
// 因为 Go 不能够隐式转换，所以我们必须显示的转换变量的类型

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
}
