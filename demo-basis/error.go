package main

import (
	"errors"
	"fmt"
)

func main() {
	var err = errors.New("test go error")
	fmt.Println(err.Error())

	defer func() {
		// recover 将终止错误的报告
		if err := recover(); err != nil {
			fmt.Println("recovered:", err)
		}
	}()

	panic("awesome " + fmt.Errorf("fmt error").Error())
}
