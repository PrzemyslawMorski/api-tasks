package main

import (
	"github.com/PrzemyslawMorski/json-api/handlers"
	"github.com/PrzemyslawMorski/json-api/server"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	var err error
	store.GlobalStoreRef, err = store.NewStore(store.DefaultFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer store.GlobalStoreRef.Close()

	router := httprouter.New()
	router.GET("/", handlers.ListTasks)
	router.GET("/:id", handlers.ReadTask)
	router.POST("/", handlers.CreateTask)
	router.PUT("/:id", handlers.UpdateTask)
	router.DELETE("/:id", handlers.DeleteTask)

	log.Fatal(http.ListenAndServe(":8000", &server.Server{Router: router}))
}
