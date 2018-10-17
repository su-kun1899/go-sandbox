package main

import (
	"fmt"
	"github.com/su-kun1899/go-sandbox/file"
	"os"
	"path/filepath"
)

func main()  {
	writer := os.Stdout
	errWriter := os.Stderr

	// TODO パスはどうにかしないと
	fileName, _ := filepath.Abs("file/testdata/append.txt")

	strings, err := file.ReadFromLast(fileName)
	if err != nil {
		fmt.Fprintf(errWriter,"%v\n", err)
		return
	}

	fmt.Fprintf(writer,"%v\n", strings[0])
}
