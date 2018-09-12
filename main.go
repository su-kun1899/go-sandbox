package main

import (
	"fmt"
	"time"
)

func main() {
	greet()
}

func greet() {
	h := time.Now().Hour()
	switch {
	case h >= 4 && h <= 9:
		fmt.Println("おはよう")
	case h >= 10 && h <= 16:
		fmt.Println("こんにちは")
	default:
		fmt.Println("こんばんは")
	}
}
