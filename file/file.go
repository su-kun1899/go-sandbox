package file

import (
	"os"
	"fmt"
	"io"
)

// TODO 複数行読み込みたい
func ReadFromLast(fileName string) ([]string, error) {
	// TODO 空ファイルの時どうなる？

	file, err := os.Open(fileName)
	if err != nil {
		// TODO untested
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		// TODO untested
		return nil, err
	}
	fmt.Printf("fileName: %v, fileSize: %v\n", fileInfo.Name(), fileInfo.Size())

	// TODO 改行はLF/CR/CRLFに対応したい
	cursor := int64(0)
	line := make([]byte, 0)
	for i := int64(0); i < fileInfo.Size(); i++ {
		cursor--

		// 一文字を読む
		file.Seek(cursor, io.SeekEnd)
		char := make([]byte, 1)
		file.Read(char)

		if char[0] == 10 {
			fmt.Println("改行だ")
			if i == 0 {
				continue
			}
			break
		}

		fmt.Printf("char[0]: %v\n", char[0])
		line = append(line, char[0])
	}

	// 逆向きソートする
	for left, right := 0, len(line)-1; left < right; left, right = left+1, right-1 {
		line[left], line[right] = line[right], line[left]
	}

	fmt.Printf("line: %v\n", string(line))

	return []string{string(line)}, nil
}
