package main

import (
	"fmt"
	"github.com/su-kun1899/go-sandbox/file"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	writer := os.Stdout
	errWriter := os.Stderr

	// TODO パスはどうにかしないと
	fileName, err := filepath.Abs("file/testdata/append.txt")
	if err != nil {
		fmt.Fprintf(errWriter, "%v\n", err)
		return
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Fprintf(errWriter, "%v\n", err)
		return
	}
	fmt.Printf("fileName: %v, fileSize: %v\n", fileInfo.Name(), fileInfo.Size())


	for {
		time.Sleep(1 * time.Second)

		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(errWriter, "%v\n", err)
			return
		}

		fileInfo, err := f.Stat()
		if err != nil {
			fmt.Fprintf(errWriter, "%v\n", err)
			return
		}
		//fmt.Printf("fileName: %v, fileSize: %v\n", fileInfo.Name(), fileInfo.Size())

		//fileInfo, err := os.Stat(fileName)
		//if err != nil {
		//	fmt.Fprintf(errWriter, "%v\n", err)
		//	return
		//}
		//
		size := fileInfo.Size()
		if size > 5 {
			fmt.Fprintf(writer, "size: %v\n", size)
			return
		}
	}

	read(writer, errWriter, fileName)
}

func read(writer, errWriter io.Writer, fileName string) {
	strings, err := file.ReadFromLast(fileName)
	if err != nil {
		fmt.Fprintf(errWriter, "%v\n", err)
		return
	}
	fmt.Fprintf(writer, "%v\n", strings[0])
}
