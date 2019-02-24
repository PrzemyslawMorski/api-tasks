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
//func TestReadTaskHandler(t *testing.T) {
//	var err error
//	store.GlobalStoreRef, err = store.NewStore()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer store.GlobalStoreRef.Close()
//
//	router := httprouter.New()
//	router.GET("/:id", ReadTask)
//	s := server.Server{Router: router}
//
//	t.Run("known task", func(t *testing.T) {
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
//		r, _ := http.NewRequest("GET", "/"+string(existingTask.Id), nil)
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
//		expectedBody, err := json.Marshal(existingTask)
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
//	t.Run("unknown task", func(t *testing.T) {
//		taskId := rand.Int()
//		for store.GlobalStoreRef.Contains(taskId) {
//			taskId = rand.Int()
//		}
//
//		r, _ := http.NewRequest("GET", "/"+string(taskId), nil)
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
//	t.Run("invalid id", func(t *testing.T) {
//		r, _ := http.NewRequest("GET", "/true", nil)
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
//			Message: "id has to be an integer",
//		})
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		expectedBodyBuffer := bytes.NewBuffer(expectedBody)
//
//		if w.Body.String() != expectedBodyBuffer.String() {
//			t.Errorf("handler returned unexpected error response: got %v want %v",
//				w.Body.String(), expectedBodyBuffer.String())
//		}
//	})
//}
