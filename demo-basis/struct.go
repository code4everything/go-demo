package main

import "fmt"
import "reflect"
import "runtime"

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

	// new(person) 和 &person{} 是等价的

	temp = person{name: "coder", gender: "boy", age: 24}
	fmt.Println(temp)

	// 标签
	t := tag{1, "tag"}
	tipe := reflect.TypeOf(t)
	for i := 0; i < 2; i++ {
		fmt.Println(tipe.Field(i).Tag)
	}

	// 匿名字段与组合（可以粗略的看成继承）
	out := outer{}
	out.name = "go"
	out.int = 1
	out.version = runtime.Version()
	out.time = 123456

	fmt.Println(out)
}

// 工厂方法
func newPerson(name string, gender string, age int) *person {
	return &person{gender, age, name}
}

type tag struct {
	id   int    "tag id"
	name string "tag name"
}

// 包内私有
type person struct {
	gender string
	age    int
	name   string
}

type inner struct {
	version string
	time    int64
}

type outer struct {
	name string
	int
	inner
}
