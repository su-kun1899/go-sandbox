package file_test

import (
	"os"
	"testing"
	"github.com/su-kun1899/go-sandbox/file"
)

func Test_readRealTime(t *testing.T) {
	fileName := "testdata/append.txt"
	file.Seek(os.Stdout, fileName)

	t.Errorf("sample")
}


//func Test_readRealTime(t *testing.T) {
//	//type args struct {
//	//	fileName string
//	//}
//	//tests := []struct {
//	//	name  string
//	//	args  args
//	//	wantW string
//	//}{
//	//	{}
//	//}
//	// TODO
//	t.Skip()
//
//	fileName := "testdata/append.txt"
//	appendStr := "bar\n"
//
//	t.Run("ファイルの追記を拾う", func(t *testing.T) {
//		w := &bytes.Buffer{}
//		file.ReadRealTime(w, fileName)
//
//		appendStringToFile(t, fileName, appendStr)
//
//		if gotW := w.String(); gotW != "foo\n" {
//			t.Errorf("readRealTime() = %v, want %v", gotW, tt.wantW)
//		}
//	})
//}

func appendStringToFile(t *testing.T, fileName, text string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}
