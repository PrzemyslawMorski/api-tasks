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
	"testing"
)

func TestUpdateTaskHandler(t *testing.T) {
	DeleteTestDb(t)

	var err error
	store.GlobalStoreRef, err = store.NewStore(testDbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer store.GlobalStoreRef.Close()

	router := httprouter.New()
	router.PUT("/:id", UpdateTask)
	router.PUT("/", UpdateTask)
	s := server.Server{Router: router}

	t.Run("empty id", func(t *testing.T) {
		newTitle := "updated title"
		reqBody, err := json.Marshal(dtos.UpdateTaskDto{NewTitle: newTitle})
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.NewRequest("PUT", "/", bytes.NewReader(reqBody))
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

		expectedCode := http.StatusMethodNotAllowed
		if w.Code != expectedCode {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		newTitle := "updated title"
		reqBody, err := json.Marshal(dtos.UpdateTaskDto{NewTitle: newTitle})
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.NewRequest("PUT", "/true", bytes.NewReader(reqBody))
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		expectedContentTypeHeader := "application/json; charset=utf-8"
		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
				w.Header().Get("Content-Type"), expectedContentTypeHeader)
		}

		expectedCode := http.StatusNotFound
		if w.Code != expectedCode {
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}
	})

	t.Run("unknown task", func(t *testing.T) {
		newTitle := "updated title"
		reqBody, err := json.Marshal(dtos.UpdateTaskDto{NewTitle: newTitle})
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.NewRequest("PUT", "/59", bytes.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		expectedContentTypeHeader := "application/json; charset=utf-8"
		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
				w.Header().Get("Content-Type"), expectedContentTypeHeader)
		}

		expectedCode := http.StatusNotFound
		if w.Code != expectedCode {
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}
	})

	t.Run("known task - empty title", func(t *testing.T) {
		task, err := store.GlobalStoreRef.CreateTask("created task")
		if err != nil {
			log.Fatal(err)
		}

		reqBody, err := json.Marshal(dtos.UpdateTaskDto{NewTitle: ""})
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.NewRequest("PUT", "/"+strconv.Itoa(task.Id), bytes.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		expectedContentTypeHeader := "application/json; charset=utf-8"
		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
				w.Header().Get("Content-Type"), expectedContentTypeHeader)
		}

		expectedCode := http.StatusBadRequest
		if w.Code != expectedCode {
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}

		expectedBody, err := json.Marshal(dtos.ErrorResponseDto{
			Code:  http.StatusBadRequest,
			Error: "title cannot be empty or all whitespace",
		})
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(w.Body.Bytes(), expectedBody) {
			t.Errorf("handler returned unexpected error response: got %v want %v",
				w.Body.String(), string(expectedBody))
		}
	})

	t.Run("known task - correct title", func(t *testing.T) {
		task, err := store.GlobalStoreRef.CreateTask("created task")
		if err != nil {
			log.Fatal(err)
		}

		newTitle := "updated title"
		reqBody, err := json.Marshal(dtos.UpdateTaskDto{NewTitle: newTitle})
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.NewRequest("PUT", "/"+strconv.Itoa(task.Id), bytes.NewReader(reqBody))
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()

		s.ServeHTTP(w, r)

		expectedContentTypeHeader := "application/json; charset=utf-8"
		if w.Header().Get("Content-Type") != expectedContentTypeHeader {
			t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
				w.Header().Get("Content-Type"), expectedContentTypeHeader)
		}

		expectedCode := http.StatusNoContent
		if w.Code != expectedCode {
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}

		if store.GlobalStoreRef.GetTaskById(task.Id).Title != newTitle {
			DeleteTestDb(t)
			t.Error("task title was not updated in db")
		}
	})

	DeleteTestDb(t)
}
