package main

import (
	"fmt"
	"github.com/su-kun1899/go-sandbox/file"
	"io"
	"log"
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

	currentLast, err := lastCursor(fileName)
	if err != nil {
		fmt.Fprintf(errWriter, "%v\n", err)
		return
	}
	fmt.Fprintf(writer, "currentLast: %v\n", currentLast)

	currentSize, err := fileSize(fileName)
	if err != nil {
		fmt.Fprintf(errWriter, "%v\n", err)
		return
	}
	for {
		time.Sleep(1 * time.Second)

		newSize, err := fileSize(fileName)
		if err != nil {
			fmt.Fprintf(errWriter, "%v\n", err)
			return
		}
		log.Printf("currentSize: %v, newSize: %v\n", currentSize, newSize)
		if err != nil {
			fmt.Fprintf(errWriter, "%v\n", err)
			return
		}

		newLast, err := lastCursor(fileName)
		log.Printf("currentLast: %v, newLast: %v\n", currentLast, newLast)
		if currentSize != newSize {
			return
		}

		//size := fileInfo.Size()
		//if size > 5 {
		//	fmt.Fprintf(writer, "size: %v\n", size)
		//	return
		//}
	}

	read(writer, errWriter, fileName)
}

func lastCursor(fileName string) (int64, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Seek(0, io.SeekEnd)
}

func fileSize(fileName string) (int64, error) {
	// ファイルをクローズすれば最新のサイズが取れそう
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

func read(writer, errWriter io.Writer, fileName string) {
	strings, err := file.ReadFromLast(fileName)
	if err != nil {
		fmt.Fprintf(errWriter, "%v\n", err)
		return
	}
	fmt.Fprintf(writer, "%v\n", strings[0])
}
