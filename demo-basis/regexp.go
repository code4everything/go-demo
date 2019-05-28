package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "oh, today is a bad day"
	regStr := "^oh"

	if ok, _ := regexp.Match(regStr, []byte(str)); ok {
		fmt.Println("sentence start with 'oh'")
	}

	if ok, _ := regexp.MatchString(regStr, str); ok {
		fmt.Println("sentence start with 'oh'")
	}

	// 通常我们会先编译正则表达式
	reg, _ := regexp.Compile("bad")
	fmt.Println(reg.ReplaceAllString(str, "nice"))
}
