package repository

import (
	"fmt"
	"github.com/jkuma/gophercises/task/database"
	bolt "go.etcd.io/bbolt"
)

const Bucket = "task"

type Task struct {
	Number int64
	Name   string
}

func AddTask(task string) error {
	db, _ := database.DB(DBName)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(Bucket))

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

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
		b, err := tx.CreateBucketIfNotExists([]byte(Bucket))

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

func ListTasks() ([]Task, error) {
	var tasks []Task

	db, _ := database.DB(DBName)
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(Bucket))

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
