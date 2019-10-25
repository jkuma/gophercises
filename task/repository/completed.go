package repository

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/jkuma/gophercises/task/database"
	bolt "go.etcd.io/bbolt"
	"time"
)

const CompletedBucket = "completed"

type CompletedTask struct {
	Time int64
	Task Task
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

