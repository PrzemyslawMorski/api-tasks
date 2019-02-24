package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/PrzemyslawMorski/json-api/dtos"
	"github.com/PrzemyslawMorski/json-api/server"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

func TestCreateTasksHandler(t *testing.T) {
	DeleteTestDb(t)

	var err error
	store.GlobalStoreRef, err = store.NewStore(testDbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer store.GlobalStoreRef.Close()

	router := httprouter.New()
	router.POST("/", CreateTask)
	s := server.Server{Router: router}

	t.Run("correct title", func(t *testing.T) {
		taskTitle := "eat some cheeeeeese"

		reqBody, err := json.Marshal(dtos.CreateTaskDto{Title: taskTitle})
		if err != nil {
			DeleteTestDb(t)
			log.Fatal(err)
		}

		r, _ := http.NewRequest("POST", "/", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		expectedContentTypeHeader := "application/json; charset=utf-8"
		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
				w.Header().Get("Content-Type"), expectedContentTypeHeader)
		}

		expectedCode := http.StatusOK
		if w.Code != expectedCode {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}

		returnedTask := &store.Task{}
		err = json.Unmarshal(w.Body.Bytes(), returnedTask)

		createdTask := store.GlobalStoreRef.GetTaskById(returnedTask.Id)
		if createdTask == nil {
			DeleteTestDb(t)
			t.Error("task was not created in db")
		}

		if createdTask.Title != taskTitle {
			t.Errorf("a task was created with an incorrect title; got %v want %v",
				createdTask.Title, taskTitle)
		}

		if returnedTask.Title != taskTitle {
			t.Errorf("handler returned a task with unexpected title; got %v want %v",
				createdTask.Title, taskTitle)
		}
	})

	t.Run("empty title", func(t *testing.T) {
		taskTitle := ""

		reqBody, err := json.Marshal(dtos.CreateTaskDto{Title: taskTitle})
		if err != nil {
			DeleteTestDb(t)
			log.Fatal(err)
		}

		r, _ := http.NewRequest("POST", "/", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		expectedContentTypeHeader := "application/json; charset=utf-8"
		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
				w.Header().Get("Content-Type"), expectedContentTypeHeader)
		}

		expectedCode := http.StatusBadRequest
		if w.Code != expectedCode {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}

		expectedErrorResponse := dtos.ErrorResponseDto{Code: expectedCode, Message: "title cannot be empty or all whitespace"}

		expectedBody, err := json.Marshal(expectedErrorResponse)
		if err != nil {
			DeleteTestDb(t)
			t.Error(err)
		}

		if !bytes.Equal(w.Body.Bytes(), expectedBody) {
			t.Errorf("handler returned an unexpected error response; got %v want %+v",
				w.Body.String(), expectedErrorResponse)
		}
	})

	t.Run("invalid title", func(t *testing.T) {
		reqBody := "{\"title\":5}"
		if err != nil {
			DeleteTestDb(t)
			log.Fatal(err)
		}

		r, _ := http.NewRequest("POST", "/", strings.NewReader(reqBody))
		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		expectedContentTypeHeader := "application/json; charset=utf-8"
		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
				w.Header().Get("Content-Type"), expectedContentTypeHeader)
		}

		expectedCode := http.StatusBadRequest
		if w.Code != expectedCode {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}

		expectedErrorResponse := dtos.ErrorResponseDto{Code: expectedCode, Message: "title has to be a string"}

		expectedBody, err := json.Marshal(expectedErrorResponse)
		if err != nil {
			DeleteTestDb(t)
			t.Error(err)
		}

		if !bytes.Equal(w.Body.Bytes(), expectedBody) {
			t.Errorf("handler returned an unexpected error response; got %v want %+v",
				w.Body.String(), expectedErrorResponse)
		}
	})

	DeleteTestDb(t)
}
