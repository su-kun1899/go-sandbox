package file

import (
	"os"
	"fmt"
	"io"
)

const LF  = 10

// TODO 複数行読み込みたい
func ReadFromLast(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		// untested: openできたのにinfoが取れないことはないはず？
		return nil, err
	}
	fmt.Printf("fileName: %v, fileSize: %v\n", fileInfo.Name(), fileInfo.Size())
	if fileInfo.Size() == 0 {
		return nil, nil
	}

	cursor := int64(0)
	line := ""
	for i := int64(0); i < fileInfo.Size(); i++ {
		cursor--

		// 一文字を読む
		file.Seek(cursor, io.SeekEnd)
		char := make([]byte, 1)
		file.Read(char)

		if i != 0 && char[0] == LF {
			break
		}
		line = fmt.Sprintf("%s%s", string(char), line)
		fmt.Printf("line: %v", line)
	}

	return []string{line}, nil
}
