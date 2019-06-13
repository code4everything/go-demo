package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("./demo-basis/io/filetest.txt")
	file, err := os.Open(path)
	defer file.Close()
	if err == nil {
		reader := bufio.NewReader(file)
		for {
			str, err := reader.ReadString('\n')
			if err == nil {
				fmt.Println(str)
			}
			if err == io.EOF {
				break
			}
		}
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println("using package ioutil")

	// 整个文件一次性读取
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		fmt.Println(string(bytes))
	}
	err = ioutil.WriteFile(path, []byte("write success\n"), os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	// 使用 fmt.Fscanln() 实现按列读取？
	// https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.2.md
}
