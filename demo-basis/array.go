package main

import "fmt"

func main() {
	// 引用类型数组
	var arr1 = new([5]int)
	fmt.Println(*arr1)
	// 值类型数组，更多的数组声明方式
	// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.1.md
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// 切片是引用类型
	s := arr2[0:4]
	s[0] = 0
	fmt.Println(s, len(s), cap(s))
	fmt.Println(arr2)
	// 切片重组
	s = make([]int, 10, 20)
	s = s[0:15]
	fmt.Println(len(s))
}
