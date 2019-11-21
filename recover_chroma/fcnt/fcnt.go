package fcnt

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileContent struct {
	Content   []byte
	Extension string
	Size      int64
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
	f, err := os.Open(filename)

	if err != nil {
		return &FileContent{}, errors.New(fmt.Sprintf("The file %v cannot been opened. Err: %v", filename, err))
	}

	fc := &FileContent{Extension: filepath.Ext(filename)}

	fc.Size, err = io.Copy(fc, f)

	return fc, err
}
