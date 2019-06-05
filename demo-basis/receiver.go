package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	tn := twoInts{23, 99}
	fmt.Println(fmt.Sprint(tn))
	n := none{}
	fmt.Println(tn.addThem())
	n.addThem()

	// 查看内存状态
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("%d Kb\n", ms.Alloc/1024)

	// 如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过函数回调方法实现
	runtime.SetFinalizer(&tn, func(t *twoInts) { fmt.Println("run time gc", t) })
	runtime.GC()
	fmt.Println("end")
}

// 接受者方法只能作用于本包内的类型
func (tn twoInts) addThem() int {
	return tn.a + tn.b
}

func (_ none) addThem() {
	fmt.Println("基于接受者的方法重载")
}

// 但是可以通过别名的方式实现
// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.6.md
func (t myTime) toString() string {
	return t.toString()
}

// alias for time.Time
// meanwhile, myTime will obtain serial methods from time.Time
type myTime struct {
	// anonymous field
	time.Time
}

type none struct{}

type twoInts struct {
	a int
	b int
}

// 类似面向对象中的setter、getter
func (tn *twoInts) B() int {
	return tn.b
}

func (tn *twoInts) SetB(b int) {
	tn.b = b
}
