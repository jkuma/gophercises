package repository

import (
	"encoding/binary"
	"github.com/jkuma/gophercises/fizzbuzz/database"
	"net/http"
	"time"
)

func Increment(key []byte) ([]byte, error) {
	db := database.Get()

	m := db.GetMergeOperator(key, add, 200*time.Millisecond)
	defer m.Stop()

	m.Add(uint64ToBytes(1))

	return m.Get()
}

func GetUri(r *http.Request) string {
	if !r.URL.IsAbs() {
		scheme := "http://"
		if r.TLS != nil {
			scheme = "https://"
		}

		return scheme + r.Host + r.URL.RequestURI()
	}

	return r.URL.String()
}

// Merge function to add two uint64 numbers
func add(existing, new []byte) []byte {
	return uint64ToBytes(bytesToUint64(existing) + bytesToUint64(new))
}

func uint64ToBytes(i uint64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], i)
	return buf[:]
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
