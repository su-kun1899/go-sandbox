package main

import (
	"fmt"
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
		//log.Printf("currentSize: %v, newSize: %v\n", currentSize, newSize)
		if err != nil {
			fmt.Fprintf(errWriter, "%v\n", err)
			return
		}

		if currentSize != newSize {
			currentSize = newSize
			currentLast, err = read(fileName, currentLast, writer)
			if err != nil {
				fmt.Fprintf(errWriter, "%v\n", err)
				return
			}
		}
	}
}

func read(fileName string, cursor int64, writer io.Writer) (int64, error) {
	lastCursor, err := lastCursor(fileName)
	if err != nil {
		return 0, err
	}

	newText := make([]byte, lastCursor-cursor)
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	f.Seek(cursor, io.SeekStart)
	f.Read(newText)
	fmt.Fprintf(writer, "%s", string(newText))

	return f.Seek(-1, io.SeekCurrent)
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
