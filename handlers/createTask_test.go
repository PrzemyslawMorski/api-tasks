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
	"strconv"
	"strings"
	"testing"
)

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

	t.Run("invalid title", func(t *testing.T) {
		reqBody := "{\"title\":5}"
		if err != nil {
			DeleteTestDb(t)
			log.Fatal(err)
		}

		r, err := http.NewRequest("POST", "/", strings.NewReader(reqBody))
		if err != nil {
			log.Fatal(err)
		}
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

		expectedErrorResponse := dtos.ErrorResponseDto{Code: expectedCode, Error: "title has to be a string"}

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

	t.Run("empty title", func(t *testing.T) {
		taskTitle := ""

		reqBody, err := json.Marshal(dtos.CreateTaskDto{Title: taskTitle})
		if err != nil {
			DeleteTestDb(t)
			log.Fatal(err)
		}

		r, err := http.NewRequest("POST", "/", bytes.NewReader(reqBody))
		if err != nil {
			log.Fatal(err)
		}
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

		expectedErrorResponse := dtos.ErrorResponseDto{Code: expectedCode, Error: "title cannot be empty or all whitespace"}

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

	t.Run("correct title", func(t *testing.T) {
		taskTitle := "eat some cheese"

		reqBody, err := json.Marshal(dtos.CreateTaskDto{Title: taskTitle})
		if err != nil {
			DeleteTestDb(t)
			log.Fatal(err)
		}

		r, err := http.NewRequest("POST", "/", bytes.NewReader(reqBody))
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		expectedContentTypeHeader := "application/json; charset=utf-8"
		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
				w.Header().Get("Content-Type"), expectedContentTypeHeader)
		}

		expectedCode := http.StatusCreated
		if w.Code != expectedCode {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}

		if w.Header().Get("Location") == "" {
			DeleteTestDb(t)
			t.Error("handler did not set the Location header")
		}

		createdTaskId, err := strconv.Atoi(w.Header().Get("Location")[1:])
		if err != nil {
			DeleteTestDb(t)
			t.Errorf("handler did not set the proper Location header; got %v", w.Header().Get("Location"))
		}

		createdTask := store.GlobalStoreRef.GetTaskById(createdTaskId)
		if createdTask == nil {
			if err != nil {
				DeleteTestDb(t)
				t.Errorf("handler did not set the proper Location header; got %v; task with this id does not exist in store",
					w.Header().Get("Location"))
			}
		}

		if createdTask.Title != taskTitle {
			t.Errorf("a task was created with an incorrect title; got %v want %v",
				createdTask.Title, taskTitle)
		}
	})

	DeleteTestDb(t)
}
