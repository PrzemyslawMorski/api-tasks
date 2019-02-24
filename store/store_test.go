package store

import (
	"os"
	"testing"
)

const testDbFileName = "__test__.db"

func DeleteTestDb(t *testing.T) {
	if _, err := os.Stat(testDbFileName); os.IsNotExist(err) {
		return
	}

	err := os.Remove(testDbFileName)
	if err != nil {
		t.Log(err)
	}
}

func TestStoreIntegration(t *testing.T) {
	DeleteTestDb(t)

	store, err := NewStore(testDbFileName)
	if err != nil {
		DeleteTestDb(t)
		t.Fatal(err)
	}

	// "store is empty after creation"
	tasks, err := store.GetTasks()
	if err != nil {
		DeleteTestDb(t)
		t.Fatal(err)
	}

	if len(tasks) != 0 {
		DeleteTestDb(t)
		t.Fatal("store has tasks after being created")
	}

	// "store remains empty after passing invalid params to CreateTask"
	task, err := store.CreateTask("")
	if err == nil {
		DeleteTestDb(t)
		t.Fatal("store didnt return an error after trying to create a task with empty title")
	}
	if task != nil {
		DeleteTestDb(t)
		t.Fatal("store didnt return a nil task after trying to create a task with empty title")
	}

	tasks, err = store.GetTasks()
	if err != nil {
		DeleteTestDb(t)
		t.Fatal(err)
	}

	if len(tasks) != 0 {
		DeleteTestDb(t)
		t.Fatal("store is not empty after trying to create task with empty title")
	}

	var createdTaskId int
	taskTitle := "my test task"

	// "store no longer empty after passing valid params to CreateTask"
	task, err = store.CreateTask(taskTitle)
	if err != nil {
		DeleteTestDb(t)
		t.Fatal(err)
	}

	if task == nil {
		DeleteTestDb(t)
		t.Fatal("store didnt return a task after trying to create a task with valid title")
	}

	if task.Title != taskTitle {
		DeleteTestDb(t)
		t.Fatal("store created a task with title other than a title passes to CreateTask")
	}

	tasks, err = store.GetTasks()
	if err != nil {
		DeleteTestDb(t)
		t.Fatal(err)
	}

	if len(tasks) != 1 {
		DeleteTestDb(t)
		t.Fatal("store remains empty after trying to create task with valid title")
	}

	createdTaskId = task.Id

	// "store contains created task"
	contains := store.Contains(createdTaskId)
	if !contains {
		DeleteTestDb(t)
		t.Fatal("store Contains returned false on a task that was just created")
	}

	// "store doesnt contain non-existent task"
	contains = store.Contains(59)
	if contains {
		DeleteTestDb(t)
		t.Fatal("store Contains task with id 59")
	}

	// "getTaskById returns created task"
	task = store.GetTaskById(createdTaskId)
	if task.Title != taskTitle {
		DeleteTestDb(t)
		t.Fatalf("getTaskById returned unexpected task; got %v want %v", task.Title, taskTitle)
	}

	// "getTaskById returns nil on non-existent task"
	task = store.GetTaskById(59)
	if task != nil {
		DeleteTestDb(t)
		t.Fatal("getTaskById didnt return a nil task on non-existent task")
	}

	// "UpdateTask doesnt update task when passed an empty new title"
	task, err = store.UpdateTask(createdTaskId, "")
	if err == nil {
		DeleteTestDb(t)
		t.Fatal("UpdateTask didnt return an error when passed an empty title")
	}

	if task != nil {
		DeleteTestDb(t)
		t.Fatal("UpdateTask didnt return a nil task when passed an empty title")
	}

	// "UpdateTask doesnt update task when passed an unknown task id"
	task, err = store.UpdateTask(59, "random text")
	if err == nil {
		DeleteTestDb(t)
		t.Fatal("UpdateTask didnt return an error when passed an unknown task id")
	}

	if task != nil {
		DeleteTestDb(t)
		t.Fatal("UpdateTask didnt return a nil task when passed an unknown task id")
	}

	taskNewTitle := "my test task - updated"

	// "UpdateTask updates task when passed a known task id and a non-empty new title"
	task, err = store.UpdateTask(createdTaskId, taskNewTitle)
	if err != nil {
		DeleteTestDb(t)
		t.Fatal(err)
	}

	updatedTask := store.GetTaskById(createdTaskId)
	if updatedTask == nil {
		DeleteTestDb(t)
		t.Fatalf("UpdateTask caused a task deletion")
	}

	if updatedTask.Title != taskNewTitle {
		DeleteTestDb(t)
		t.Fatalf("UpdateTask didnt update title; got %v want %v", updatedTask.Title, taskNewTitle)
	}

	if task.Title != taskNewTitle {
		DeleteTestDb(t)
		t.Fatalf("UpdateTask didnt return task with updated title; got %v want %v", updatedTask.Title, taskNewTitle)
	}

	// "DeleteTask does nothing when passed an unknown task id"
	err = store.DeleteTask(59)
	if err != nil {
		DeleteTestDb(t)
		t.Fatal(err)
	}

	// "DeleteTask deletes a task when passed an known task id"
	err = store.DeleteTask(createdTaskId)
	if err != nil {
		DeleteTestDb(t)
		t.Fatal(err)
	}

	DeleteTestDb(t)
}
