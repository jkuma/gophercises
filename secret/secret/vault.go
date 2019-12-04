package secret

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type VaultAccess interface {
	Set(key, val string) error
	Get(key string) (string, error)
}

type FileVault struct {
	Key      []byte
	Filename string
	Mutex    sync.Mutex
}

func (f FileVault) Set(key, val string) error {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()
	m, err := fileArray(f.Filename)

	if err == nil {
		m[unique(key)] = encrypt(f.Key, val)
		_, err = writeFile(f.Filename, m)
	}

	return err
}

func (f FileVault) Get(key string) (string, error) {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()
	var res string
	m, err := fileArray(f.Filename)

	if err == nil {
		if val, ok := m[unique(key)]; ok {
			return decrypt(f.Key, val), nil
		}

		return res, errors.New(fmt.Sprintf("The key %s is not found", key))
	}

	return res, err
}

func fileArray(filename string) (map[string][]byte, error) {
	result := make(map[string][]byte)
	file, err := os.Open(filename)

	if err == nil {
		defer file.Close()
		s := bufio.NewScanner(file)

		for s.Scan() {
			cl := strings.Split(s.Text(), "=")

			if len(cl) != 2 {
				return result, errors.New(fmt.Sprintf("Line %s does not contains any '='", s.Text()))
			}

			b, err := hex.DecodeString(cl[1])

			if err == nil {
				result[cl[0]] = b
			}
		}
	}

	return result, err
}

func writeFile(filename string, m map[string][]byte) (int64, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC, os.ModePerm)

	if err == nil {
		defer file.Close()
		var buf bytes.Buffer

		for key, val := range m {
			buf.Write([]byte(fmt.Sprintf("%s=%s\n", unique(key), hex.EncodeToString(val))))
		}

		return io.Copy(file, &buf)
	}

	return 0, err
}

func unique(key string) string {
	return strings.ToUpper(key)
}
