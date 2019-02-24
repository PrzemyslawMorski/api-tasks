package handlers

//
//import (
//	"github.com/PrzemyslawMorski/json-api/store"
//	"github.com/julienschmidt/httprouter"
//	"log"
//	"testing"
//)
//
//func TestDeleteTaskHandler(t *testing.T) {
//	var err error
//	store.GlobalStoreRef, err = store.NewStore()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer store.GlobalStoreRef.Close()
//
//	router := httprouter.New()
//	router.DELETE("/:id", DeleteTask)
//	//s := server.Server{Router: router}
//
//	t.Run("known task", func(t *testing.T) {
//		// TODO
//	})
//
//	t.Run("unknown task", func(t *testing.T) {
//		// TODO
//	})
//
//	t.Run("invalid id", func(t *testing.T) {
//		// TODO
//	})
//}
