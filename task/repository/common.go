package repository

import "encoding/binary"

const DBName = "my.db"

func i64tob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Returns an integer of byte slice b.
func btoi64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}
