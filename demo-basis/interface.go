package main

import "fmt"

func main() {
	var l learner = employer{"trump"}
	l.learn()
	// 转型（实际返回两个参数，第二个为是否转换成功的布尔值）
	var e = l.(employer)
	fmt.Println(e.name)

	fmt.Println(typeOf(true))
}

// type switch
// 空接口可以看成所有类型的基类
func typeOf(item interface{}) (t string) {
	switch item.(type) {
	case int, int8, int16, int32, int64:
		t = "integer"
	case float32, float64:
		t = "float"
	case bool:
		t = "boolean"
	case string:
		t = "string"
	default:
		t = "unknown"
	}
	return
}

type learner interface {
	learn()
}

type employer struct {
	name string
}

func (e employer) learn() {
	fmt.Println(e.name, "is learning")
}

// 接口的继承
type locker interface {
	lock()
	unlock()
}

type file interface {
	locker
	read()
	write()
}
