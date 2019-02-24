package store

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"strings"
)

type Store struct {
	db *bolt.DB
}

var GlobalStoreRef *Store

const TasksBucketName = "tasks"
const DefaultFileName = "tasks.db"

func NewStore(fileName string) (*Store, error) {
	if fileName == "" {
		return nil, errors.New("fileName cannot be empty")
	}

	handle, err := bolt.Open(fileName, 0600, nil)

	if err != nil {
		return nil, err
	}

	store := &Store{db: handle}

	err = store.initialize()
	if err != nil {
		return nil, err
	}

	GlobalStoreRef = store

	return store, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (s *Store) initialize() error {
	return s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(TasksBucketName))

		if err != nil {
			return err
		}

		return nil
	})
}

func (s *Store) GetTasks() ([]*Task, error) {
	var tasks []*Task

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksBucketName))

		err := b.ForEach(func(k, v []byte) error {
			var t Task
			err := json.Unmarshal(v, &t)

			if err != nil {
				return err
			}

			tasks = append(tasks, &t)
			return nil
		})

		if err != nil {
			return nil
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Store) CreateTask(title string) (*Task, error) {
	if strings.Trim(title, " ") == "" {
		return nil, errors.New("title cannot be empty or all whitespace")
	}

	task := &Task{}

	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksBucketName))

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		task.Id = int(id)
		task.Title = title

		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(itob(task.Id), buf)
	})

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Store) Contains(id int) bool {
	tasks, err := s.GetTasks()
	if err != nil {
		return false
	}

	for _, task := range tasks {
		if task.Id == id {
			return true
		}
	}

	return false
}

func (s *Store) GetTaskById(id int) *Task {
	tasks, err := s.GetTasks()
	if err != nil {
		return nil
	}

	for _, task := range tasks {
		if task.Id == id {
			return task
		}
	}

	return nil
}

func (s *Store) DeleteTask(id int) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksBucketName))
		return b.Delete(itob(id))
	})
}

func (s *Store) UpdateTask(id int, newTitle string) (*Task, error) {
	if strings.Trim(newTitle, " ") == "" {
		return nil, errors.New("title cannot be empty or all whitespace")
	}

	if s.GetTaskById(id) == nil {
		return nil, errors.New("task with id " + string(id) + " doesn't exist")
	}

	task := &Task{Id: id, Title: newTitle}

	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksBucketName))

		buf, err := json.Marshal(task)
		if err != nil {
			return err
		}

		return b.Put(itob(task.Id), buf)
	})

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}
