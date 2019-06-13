package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var NewLine = flag.Bool("n", false, "print new line")

// using: go run args.go arg1 arg2
func main() {
	who := "alice "
	fmt.Println(os.Args[0])
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("good morning", who)

	flag.PrintDefaults()
	// 扫描参数列表
	flag.Parse()
	var s = ""
	var sep = ""
	for i := 0; i < flag.NArg(); i++ {
		s += sep + flag.Arg(i)
		if *NewLine {
			sep = "\n"
		} else {
			sep = " "
		}
	}
	os.Stdout.WriteString(s)
}
