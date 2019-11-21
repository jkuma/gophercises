package fcnt

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type FileContent struct {
	Content []byte
	Size    int64
}

func (f *FileContent) Write(p []byte) (n int, err error) {
	for _, b := range p {
		f.Content = append(f.Content, b)
	}

	return len(p), nil
}

func (f *FileContent) String() string {
	return string(f.Content)
}

func GetFileContent(filename string) (*FileContent, error) {
	fc := new(FileContent)
	f, err := os.Open(filename)

	if err != nil {
		return fc, errors.New(fmt.Sprintf("The file %v cannot been opened. Err: %v", filename, err))
	}

	fc.Size, err = io.Copy(fc, f)

	return fc, err
}
