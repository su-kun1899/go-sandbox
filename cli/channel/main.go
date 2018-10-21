package main

import (
	"fmt"
	"time"
	"log"
)

var foo int = 0

type diff struct {
	oldVal int
	newVal int
}

func watchFoo(c chan diff) {
	currentVal := foo
	log.Printf("currentVal: %d\n", currentVal)

	for {
		if currentVal != foo {
			log.Printf("newVal: %d\n", foo)
			c <- diff{oldVal: currentVal, newVal: foo}
			return
		}
	}
}

func main() {
	c := make(chan diff)

	go watchFoo(c)

	time.Sleep(3 * time.Second)
	foo = 3

	diff := <-c

	fmt.Printf("diff: %v\n", diff)
}
