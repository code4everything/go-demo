package main

import "fmt"

func main() {
	// 分配内存
	lover := new(person)
	lover.name = "fairy"
	lover.age = 22
	lover.gender = "girl"

	fmt.Println(*lover)

	// 初始化
	var temp person
	temp = person{"girl", 23, "elf"}
	fmt.Println(temp)

	temp = person{name: "coder", gender: "boy", age: 24}
	fmt.Println(temp)
}

// 包内私有
type person struct {
	gender string
	age    int
	name   string
}
