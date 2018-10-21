package main

import (
	"fmt"
	"time"
)

var foo int = 0

type diff struct {
	oldVal int
	newVal int
}

func watchFoo(c chan int) {
	currentVal := foo

	for {
		if currentVal != foo {
			c <- foo
		}
	}
}

func main() {
	c := make(chan int)

	go watchFoo(c)

	time.Sleep(3 * time.Second)
	foo = 3

	newVal := <-c

	fmt.Printf("newVal: %d", newVal)

	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // receive from c
	//
	//fmt.Println(x, y, x+y) // -5 17 12
}
