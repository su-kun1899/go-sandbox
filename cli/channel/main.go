package main

import (
	"fmt"
	"time"
	"log"
	"errors"
)

var foo int = 0

type diff struct {
	oldVal int
	newVal int
}

func watchFoo(d chan diff, err chan error) {
	currentVal := foo
	log.Printf("currentVal: %d\n", currentVal)

	for {
		if currentVal != foo {
			log.Printf("newVal: %d\n", foo)

			if foo > 100 {
				// 100以上の場合はエラーにする
				err <- errors.New("foo is invalid status")
				return
			}

			d <- diff{oldVal: currentVal, newVal: foo}
			return
		}
	}
}

func main() {
	d := make(chan diff)
	e := make(chan error)

	log.Println("1回目")
	go watchFoo(d, e)

	time.Sleep(3 * time.Second)
	foo = 3
	diff := <-d
	log.Printf("diff: %v\n", diff)


	log.Println("2回目")
	go watchFoo(d, e)
	time.Sleep(3 * time.Second)
	foo = 100
	select {
	case a := <-d:
		log.Printf("diff: %v\n", a)
		return
	case b := <-e:
		log.Printf("err: %v\n", b)
		return
	}
	fmt.Printf("diff: %v\n", diff)

	log.Println("3回目")
}
