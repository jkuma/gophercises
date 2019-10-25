package task

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

const MainBucket = "tasks"
const CompletedBucket = "completed"

type Task struct {
	Number int
	Name   string
}

type CompletedTask struct {
	Time int64
	Task Task
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
		b, err := tx.CreateBucketIfNotExists([]byte(MainBucket))

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

func DeleteTask(number int, completed bool) (Task, error) {
	var t Task
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(MainBucket))

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

	if err == nil && completed {
		err = markTaskAsCompleted(db, t)
	}

	return t, err
}

func markTaskAsCompleted(db *bolt.DB, task Task) error {
	err := db.Update(func(tx *bolt.Tx) error {
		var buffer bytes.Buffer
		b, err := tx.CreateBucketIfNotExists([]byte(CompletedBucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = gob.NewEncoder(&buffer).Encode(task)

		if err != nil {
			return fmt.Errorf("buffer error: %s", err)
		}

		return b.Put(i64tob(time.Now().Unix()), buffer.Bytes())
	})

	return err
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
		b, err := tx.CreateBucketIfNotExists([]byte(MainBucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = b.ForEach(func(k, v []byte) error {
			tasks = append(tasks, Task{
				Number: btoi(k),
				Name:   string(v),
			})

			return nil
		})

		return err
	})

	return tasks, err
}

func ListCompletedTasks() (map[int]Task, error) {
	tasks := map[int]Task{}
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		var buffer bytes.Buffer
		b, err := tx.CreateBucketIfNotExists([]byte(CompletedBucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = b.ForEach(func(k, v []byte) error {
			var t Task
			_, err = buffer.Write(v)
			err = gob.NewDecoder(&buffer).Decode(&t)
			buffer.Reset()

			tasks[btoi(k)] = t

			return err
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

func i64tob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Returns an integer of byte slice b.
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
