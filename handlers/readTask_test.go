package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/PrzemyslawMorski/json-api/server"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestReadTaskHandler(t *testing.T) {
	DeleteTestDb(t)

	var err error
	store.GlobalStoreRef, err = store.NewStore(testDbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer store.GlobalStoreRef.Close()

	router := httprouter.New()
	router.GET("/:id", ReadTask)
	s := server.Server{Router: router}

	t.Run("invalid id", func(t *testing.T) {
		r, err := http.NewRequest("GET", "/true", nil)
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
		r, err := http.NewRequest("GET", "/59", nil)
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

	t.Run("known task", func(t *testing.T) {
		task, err := store.GlobalStoreRef.CreateTask("created task")
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.NewRequest("GET", "/"+strconv.Itoa(task.Id), nil)
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

		expectedCode := http.StatusOK
		if w.Code != expectedCode {
			t.Errorf("handler returned unexpected status code: got %v want %v",
				w.Code, expectedCode)
		}

		expectedBody, err := json.Marshal(task)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(w.Body.Bytes(), expectedBody) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), string(expectedBody))
		}
	})

	DeleteTestDb(t)
}
