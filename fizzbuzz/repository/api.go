package repository

import (
	"github.com/dgraph-io/badger"
	"github.com/jkuma/gophercises/fizzbuzz/database"
	"net/http"
)

func Update(r *http.Request) error {
	db := database.Get()

	key := []byte(GetUri(r))

	return db.Update(func(txn *badger.Txn) error {
		val, err := Increment(key)

		if err == badger.ErrKeyNotFound {
			return txn.Set(key, val)
		}

		return err
	})
}

func Get(key []byte) ([]byte, error) {
	var val []byte
	db := database.Get()

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		return item.Value(func(v []byte) error {
			val = append([]byte{}, v...)
			return err
		})
	})

	return val, err
}

func HighScore() (key []byte, score int, err error) {
	db := database.Get()

	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				if int(bytesToUint64(v)) > score {
					score = int(bytesToUint64(v))
					key = item.Key()
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return key, score, err
}
