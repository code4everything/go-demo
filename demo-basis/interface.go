package main

import "fmt"

func main() {
	var l learner = employer{"trump"}
	l.learn()
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
