package repository

import (
	"github.com/dgraph-io/badger"
	"github.com/jkuma/gophercises/fizzbuzz/database"
	"net/http"
)

func Update(r *http.Request) error {
	db := database.Get()
	defer db.Close()

	key := []byte(r.URL.RequestURI())

	return db.Update(func(txn *badger.Txn) error {
		val, err := MergedValue(key)

		switch err {
		case badger.ErrKeyNotFound:
			e := badger.NewEntry(key, DefaultValue())
			return txn.SetEntry(e)
		case nil:
			return txn.Set(key, val)
		default:
			return err
		}
	})
}
