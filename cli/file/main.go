package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"time"
	"path/filepath"
)

func main() {
	// TODO パスはどうにかしないと
	fileName, err := filepath.Abs("file/testdata/append.txt")
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}
	//detectFileSizeChange(os.Stdout, fileName)
	seek(os.Stdout, fileName)

	var chars = []byte{97, 98, 99, 10, 100, 101, 102, 103, 10, 104, 105, 106, 107, 108, 109, 110, 111, 112, 10, 113, 114, 10, 0, 0, 0, 0, 0, 0, 0, 0}
	extractLine(os.Stdout, chars)
}

func extractLine(w io.Writer, chars []byte) {
	var extracted = make([]byte, 0, len(chars))
	for _, char := range chars {
		extracted = append(extracted, char)
	}
	fmt.Fprintf(w, "extracted: %v\n", extracted)
}

func seek(w io.Writer, fileName string) error {

	fp, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()
	fileInfo, err := fp.Stat()
	if err != nil {
		return err
	}

	var offset, limit int64 = 0, 30

	cursor, err := fp.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "fileSize: %v\n", fileInfo.Size())
	fmt.Fprintf(w, "cursor1: %v\n", cursor)

	chars := make([]byte, limit-offset)
	fp.Read(chars)

	fmt.Fprintf(w, "chars:%s \n", string(chars))

	cursor, err = fp.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "cursor2: %v\n", cursor)
	fmt.Fprintf(w, "chars: %v\n", chars)

	return nil
}

func detectFileSizeChange(w io.Writer, fileName string) error {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		preSize := fileInfo.Size()
		fileInfo, err = os.Stat(fileName)
		if err != nil {
			return err
		}
		currentSize := fileInfo.Size()

		fmt.Fprintf(w, "preSize: %v, currentSize:%v\n", preSize, currentSize)

		time.Sleep(1 * time.Second)
	}

	return nil
}

//func main() {
//	writer := os.Stdout
//	errWriter := os.Stderr
//
//	// TODO パスはどうにかしないと
//	fileName, err := filepath.Abs("file/testdata/append.txt")
//	if err != nil {
//		fmt.Fprintf(errWriter, "%v\n", err)
//		return
//	}
//
//	fileInfo, err := os.Stat(fileName)
//	if err != nil {
//		fmt.Fprintf(errWriter, "%v\n", err)
//		return
//	}
//	fmt.Printf("fileName: %v, fileSize: %v\n", fileInfo.Name(), fileInfo.Size())
//
//	currentLast, err := lastCursor(fileName)
//	if err != nil {
//		fmt.Fprintf(errWriter, "%v\n", err)
//		return
//	}
//	fmt.Fprintf(writer, "currentLast: %v\n", currentLast)
//
//	currentSize, err := fileSize(fileName)
//	if err != nil {
//		fmt.Fprintf(errWriter, "%v\n", err)
//		return
//	}
//	for {
//		time.Sleep(1 * time.Second)
//
//		newSize, err := fileSize(fileName)
//		if err != nil {
//			fmt.Fprintf(errWriter, "%v\n", err)
//			return
//		}
//		//log.Printf("currentSize: %v, newSize: %v\n", currentSize, newSize)
//		if err != nil {
//			fmt.Fprintf(errWriter, "%v\n", err)
//			return
//		}
//
//		if currentSize != newSize {
//			currentSize = newSize
//			currentLast, err = read(fileName, currentLast, writer)
//			if err != nil {
//				fmt.Fprintf(errWriter, "%v\n", err)
//				return
//			}
//		}
//	}
//}

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
