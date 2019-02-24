package handlers

import (
	"encoding/json"
	"errors"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func ListTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if store.GlobalStoreRef == nil {
		log.Fatal(errors.New("database was not initialized"))
	}

	tasks, err := store.GlobalStoreRef.GetTasks()
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
