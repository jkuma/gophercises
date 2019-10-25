package task

import (
	"encoding/binary"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
)

const BucketName = "tasks"

type Task struct {
	Number int
	Name   string
}

func AddTask(task string) error {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()

		// Persist bytes to users bucket.
		return b.Put(itob(int(id)), []byte(task))
	})

	return err
}

func DeleteTask(number int) (Task, error) {
	var t Task
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		v := string(b.Get(itob(number)))

		if v != "" {
			t.Number, t.Name = number, v
		}

		// Persist bytes to users bucket.
		return b.Delete(itob(number))
	})

	return t, err
}

func ListTasks() ([]Task, error) {
	var tasks []Task
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = b.ForEach(func(k, v []byte) error {
			tasks = append(tasks, Task{
				Number: int(binary.BigEndian.Uint64(k)),
				Name:   string(v),
			})

			return nil
		})

		return err
	})

	return tasks, err
}

// Returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
