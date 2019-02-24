package handlers

//
//import (
//	"github.com/PrzemyslawMorski/json-api/server"
//	"github.com/PrzemyslawMorski/json-api/store"
//	"github.com/julienschmidt/httprouter"
//	"log"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestListTasksHandler(t *testing.T) {
//	var err error
//	store.GlobalStoreRef, err = store.NewStore()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer store.GlobalStoreRef.Close()
//
//	router := httprouter.New()
//	router.GET("/", ListTasks)
//	s := server.Server{Router:router}
//
//	r, _ := http.NewRequest("GET", "/", nil)
//	w := httptest.NewRecorder()
//
//	s.ServeHTTP(w, r)
//
//	expectedContentTypeHeader := "application/json; charset=utf-8"
//	if w.Header().Get("Content-Type") != expectedContentTypeHeader {
//		t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
//			w.Header().Get("Content-Type"), expectedContentTypeHeader)
//	}
//
//	expectedCode := http.StatusOK
//	if w.Code != expectedCode {
//		t.Errorf("handler returned unexpected status code: got %v want %v",
//			w.Code, expectedCode)
//	}
//}
