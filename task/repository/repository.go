package repository

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/jkuma/gophercises/task/database"
	bolt "go.etcd.io/bbolt"
	"time"
)

const MainBucket = "tasks"
const CompletedBucket = "completed"
const DBName = "my.db"

type Task struct {
	Number int64
	Name   string
}

type CompletedTask struct {
	Time int64
	Task Task
}

func AddTask(task string) error {
	db, _ := database.DB(DBName)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(MainBucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()

		// Persist bytes to users bucket.
		return b.Put(i64tob(int64(id)), []byte(task))
	})

	return err
}

func DeleteTask(number int64) (Task, error) {
	var t Task
	db, _ := database.DB(DBName)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(MainBucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		v := string(b.Get(i64tob(number)))

		if v != "" {
			t.Number, t.Name = number, v
		}

		// Persist bytes to users bucket.
		return b.Delete(i64tob(number))
	})

	return t, err
}

func MarkTaskAsCompleted(task Task) error {
	db, _ := database.DB(DBName)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		var buffer bytes.Buffer
		b, err := tx.CreateBucketIfNotExists([]byte(CompletedBucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = gob.NewEncoder(&buffer).Encode(CompletedTask{
			Time: time.Now().Unix(),
			Task: task,
		})

		if err != nil {
			return fmt.Errorf("buffer error: %s", err)
		}

		id, _ := b.NextSequence()

		return b.Put(i64tob(int64(id)), buffer.Bytes())
	})

	return err
}

func ListTasks() ([]Task, error) {
	var tasks []Task

	db, _ := database.DB(DBName)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(MainBucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = b.ForEach(func(k, v []byte) error {
			tasks = append(tasks, Task{
				Number: btoi64(k),
				Name:   string(v),
			})

			return nil
		})

		return err
	})

	return tasks, err
}

func ListCompletedTasks() ([]CompletedTask, error) {
	var tasks []CompletedTask

	db, _ := database.DB(DBName)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		var buffer bytes.Buffer
		b, err := tx.CreateBucketIfNotExists([]byte(CompletedBucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = b.ForEach(func(k, v []byte) error {
			var t CompletedTask
			_, err = buffer.Write(v)
			err = gob.NewDecoder(&buffer).Decode(&t)
			buffer.Reset()

			tasks = append(tasks, t)

			return err
		})

		return err
	})

	return tasks, err
}

func i64tob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Returns an integer of byte slice b.
func btoi64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}
