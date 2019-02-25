package handlers

import (
	"github.com/PrzemyslawMorski/json-api/server"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDeleteTaskHandler(t *testing.T) {
	DeleteTestDb(t)

	var err error
	store.GlobalStoreRef, err = store.NewStore(testDbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer store.GlobalStoreRef.Close()

	router := httprouter.New()
	router.DELETE("/:id", DeleteTask)
	router.DELETE("/", DeleteTask)
	s := server.Server{Router: router}

	t.Run("empty id", func(t *testing.T) {
		r, err := http.NewRequest("DELETE", "/", nil)
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
		r, err := http.NewRequest("DELETE", "/true", nil)
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

		expectedCode := http.StatusNotFound
		if w.Code != expectedCode {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}
	})

	t.Run("unknown task", func(t *testing.T) {
		r, err := http.NewRequest("DELETE", "/59", nil)
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

		expectedCode := http.StatusNotFound
		if w.Code != expectedCode {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}
	})

	t.Run("known task", func(t *testing.T) {
		task, err := store.GlobalStoreRef.CreateTask("created task")
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.NewRequest("DELETE", "/"+strconv.Itoa(task.Id), nil)
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

		expectedCode := http.StatusOK
		if w.Code != expectedCode {
			DeleteTestDb(t)
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}

		if store.GlobalStoreRef.GetTaskById(task.Id) != nil {
			DeleteTestDb(t)
			t.Error("task was not deleted from db")
		}
	})

	DeleteTestDb(t)
}
