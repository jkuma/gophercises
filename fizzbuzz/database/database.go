package database

import (
	"github.com/dgraph-io/badger"
)

var db *badger.DB

func Get() *badger.DB {
	return db
}

func Init(filename string) error {
	d, err := badger.Open(badger.DefaultOptions(filename))
	if err != nil {
		return err
	}

	db = d

	return nil
}
