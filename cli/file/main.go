package main

import (
	"fmt"
	"os"
)

func main()  {
	writer := os.Stdout

	fmt.Fprint(writer,"hello\n")
}
