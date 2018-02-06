package main

import (
	"fmt"
)

type Strategy func()

type Context struct {
	strategy Strategy
}

func (c Context) Exec() {
	c.strategy()
}

func main() {
	var st1, st2 Strategy
	st1 = func() {
		fmt.Println("Strategy 1")
	}
	st2 = func() {
		fmt.Println("Strategy 2")
	}

	Context{st1}.Exec()
	Context{st2}.Exec()
}
