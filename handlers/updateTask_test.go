package handlers

//
//import (
//	"bytes"
//	"encoding/json"
//	"errors"
//	"github.com/PrzemyslawMorski/json-api/dtos"
//	"github.com/PrzemyslawMorski/json-api/server"
//	"github.com/PrzemyslawMorski/json-api/store"
//	"github.com/julienschmidt/httprouter"
//	"log"
//	"math/rand"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestUpdateTaskHandler(t *testing.T) {
//	var err error
//	store.GlobalStoreRef, err = store.NewStore()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer store.GlobalStoreRef.Close()
//
//	router := httprouter.New()
//	router.PUT("/:id", UpdateTask)
//	s := server.Server{Router: router}
//
//	t.Run("invalid id", func(t *testing.T) {
//
//	})
//
//	t.Run("unknown task", func(t *testing.T) {
//		unknownTaskId := rand.Int()
//		for store.GlobalStoreRef.Contains(unknownTaskId) {
//			unknownTaskId = rand.Int()
//		}
//
//		newTitle := "updated title"
//		reqBody, err := json.Marshal(dtos.UpdateTaskDto{NewTitle: newTitle})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		r, _ := http.NewRequest("POST", "/"+string(unknownTaskId), bytes.NewReader(reqBody))
//
//		w := httptest.NewRecorder()
//
//		s.ServeHTTP(w, r)
//
//		expectedContentTypeHeader := "application/json; charset=utf-8"
//		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
//			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
//				w.Header().Get("Content-Type"), expectedContentTypeHeader)
//		}
//
//		expectedCode := http.StatusNotFound
//		if w.Code != expectedCode {
//			t.Errorf("handler returned unexpected status code: got %v want %v",
//				w.Code, expectedCode)
//		}
//	})
//
//	t.Run("known task - correct title", func(t *testing.T) {
//		tasksInDb, err := store.GlobalStoreRef.GetTasks()
//		if err != nil {
//			t.Fatal(err)
//		}
//		if len(tasksInDb) == 0 {
//			t.Fatal(errors.New("no tasks in db"))
//		}
//
//		existingTask := tasksInDb[0]
//		oldTitle := existingTask.Title
//		newTitle := oldTitle + " - updated"
//
//		reqBody, err := json.Marshal(dtos.UpdateTaskDto{NewTitle: newTitle})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		r, _ := http.NewRequest("POST", "/"+string(existingTask.Id), bytes.NewReader(reqBody))
//		w := httptest.NewRecorder()
//
//		s.ServeHTTP(w, r)
//
//		expectedContentTypeHeader := "application/json; charset=utf-8"
//		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
//			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
//				w.Header().Get("Content-Type"), expectedContentTypeHeader)
//		}
//
//		expectedCode := http.StatusOK
//		if w.Code != expectedCode {
//			t.Errorf("handler returned unexpected status code: got %v want %v",
//				w.Code, expectedCode)
//		}
//
//		updatedTask := store.Task{Id: existingTask.Id, Title: newTitle}
//
//		expectedBody, err := json.Marshal(updatedTask)
//		if err != nil {
//			t.Fatal(err)
//		}
//		expectedBodyBuffer := bytes.NewBuffer(expectedBody)
//
//		if w.Body.String() != expectedBodyBuffer.String() {
//			t.Errorf("handler returned unexpected success response: got %v want %v",
//				w.Body.String(), expectedBodyBuffer.String())
//		}
//
//		updatedTaskFromDb, err := store.GlobalStoreRef.GetTaskById(existingTask.Id)
//		if err != nil {
//			t.Fatal(err)
//		}
//		if updatedTaskFromDb.Title != newTitle {
//			t.Errorf("task was not updated in db: got %+v want %+v",
//				updatedTaskFromDb, updatedTask)
//		}
//	})
//
//	t.Run("known task - empty title", func(t *testing.T) {
//		tasksInDb, err := store.GlobalStoreRef.GetTasks()
//		if err != nil {
//			t.Fatal(err)
//		}
//		if len(tasksInDb) == 0 {
//			t.Fatal(errors.New("no tasks in db"))
//		}
//
//		existingTask := tasksInDb[0]
//
//		reqBody, err := json.Marshal(dtos.UpdateTaskDto{NewTitle: ""})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		r, err := http.NewRequest("POST", "/"+string(existingTask.Id), bytes.NewReader(reqBody))
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		w := httptest.NewRecorder()
//
//		s.ServeHTTP(w, r)
//
//		expectedContentTypeHeader := "application/json; charset=utf-8"
//		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
//			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
//				w.Header().Get("Content-Type"), expectedContentTypeHeader)
//		}
//
//		expectedCode := http.StatusBadRequest
//		if w.Code != expectedCode {
//			t.Errorf("handler returned unexpected status code: got %v want %v",
//				w.Code, expectedCode)
//		}
//
//		expectedBody, err := json.Marshal(dtos.ErrorResponseDto{
//			Code:    http.StatusBadRequest,
//			Message: "title cannot be empty",
//		})
//		if err != nil {
//			t.Fatal(err)
//		}
//		expectedBodyBuffer := bytes.NewBuffer(expectedBody)
//
//		if w.Body.String() != expectedBodyBuffer.String() {
//			t.Errorf("handler returned unexpected error response: got %v want %v",
//				w.Body.String(), expectedBodyBuffer.String())
//		}
//	})
//
//	t.Run("known task - invalid title", func(t *testing.T) {
//
//	})
//}
