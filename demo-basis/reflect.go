package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("float:", v.Float())
	fmt.Println("interface:", v.Interface())

	y := v.Interface().(float64)
	fmt.Println("float64:", y)

	// 修改值，某些值需要通过地址去改变
	fmt.Println()
	fmt.Println("settable:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type:", v.Type())
	fmt.Println("settable:", v.CanSet())

	fmt.Println()
	v = v.Elem()
	fmt.Println("elem:", v)
	fmt.Println("settable:", v.CanSet())
	// this works
	v.SetFloat(6.8)
	fmt.Println("value:", x)

	// 反射字段
	fmt.Println()
	var s interface{} = unknownType{"java", "go", "python"}
	v = reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("field %s: %s\n", t.Field(i).Name, v.Field(i))
	}
	fmt.Println("string:", v.Method(0).Call(nil))
}

type unknownType struct {
	s1, s2, s3 string
}

func (n unknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}
