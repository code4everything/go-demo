package _package

import (
	"./pack1"
	// 当使用英文句号作为包别名时，你可以不通过包名来使用其中的项目
	. "./pack2"
	// 当使用下划线时，只执行包内的初始化函数
	_ "./pack3"
	"fmt"
	"testing"
)

func TestPack(t *testing.T) {
	pack1.PrintNow()
	fmt.Println(GetRandomString(7))
}
