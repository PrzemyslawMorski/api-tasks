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
	"testing"
)

func TestListTasksHandler(t *testing.T) {
	DeleteTestDb(t)

	var err error
	store.GlobalStoreRef, err = store.NewStore(testDbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer store.GlobalStoreRef.Close()

	task1, err := store.GlobalStoreRef.CreateTask("created task1")
	if err != nil {
		log.Fatal(err)
	}

	task2, err := store.GlobalStoreRef.CreateTask("created task2")
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.GET("/", ListTasks)
	s := server.Server{Router: router}

	r, _ := http.NewRequest("GET", "/", nil)
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

	expectedBody, err := json.Marshal([]store.Task{*task1, *task2})
	if !bytes.Equal(w.Body.Bytes(), expectedBody) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expectedBody))
	}

	DeleteTestDb(t)
}
