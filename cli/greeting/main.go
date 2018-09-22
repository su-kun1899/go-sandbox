package main

import (
	"github.com/su-kun1899/go-sandbox/greeting"
	"os"
)

func main() {
	var g greeting.Greeting
	g.Do(os.Stdout)
}
