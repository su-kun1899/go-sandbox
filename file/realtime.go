package file

import (
	"io"
	"os"
	"fmt"
)

func Seek(w io.Writer, fileName string) error {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "fileSize: %v\n", fileInfo.Size())

	return nil
}
