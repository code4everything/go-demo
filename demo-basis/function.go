package main

import "fmt"

// 在go语言中是没有方法重载的，因为会消耗过多的性能
// 函数是不能嵌套的，但是匿名函数可以打破这个规则
func main() {
	x := 13
	callByValue(x)
	fmt.Println(x)

	callByReference(&x)
	fmt.Println(x)

	fmt.Println(fib(5))

	fmt.Println(sum(5, 5))

	testDefer()

	var max int = 1e3
	// 想使用闭包函数调试？
	// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.10.md
	continueSum(max, func(sum int) {
		fmt.Printf("the sum of 1 to %d is %d", max, sum)
	})
}

// 将函数作为参数，常用于回调，当然我们也可将函数作为返回值
func continueSum(max int, call func(int)) {
	sum := 0
	for i := 1; i <= max; i++ {
		sum += i
	}
	call(sum)
}

// defer 推迟到函数返回之前执行
func testDefer() {
	fmt.Println("starting")
	defer fmt.Println("finished")
	fmt.Println("started")
}

// 变长参数的函数，不带变量名的返回值
// 接收不同类型的变长参数？
// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.3.md
func sum(nums ...int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

// 带变量名的返回值
func fib(len uint32) (s uint32) {
	s = 1
	var p uint32 = 1
	for ; len > 2; len-- {
		p, s = s, s+p
	}
	return
	// also works: return s
}

// package private
func callByValue(x int) {
	fmt.Println(x)
	*&x = 0
	fmt.Println(x)
}

func callByReference(x *int) {
	fmt.Println(*x)
	*x = 0
	fmt.Println(*x)
}
