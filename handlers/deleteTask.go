package handlers

import (
	"errors"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func DeleteTask(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	if store.GlobalStoreRef == nil {
		log.Fatal(errors.New("database was not initialized"))
	}

	id := ps.ByName("id")
	if id == "" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	knownId := store.GlobalStoreRef.Contains(taskId)
	if !knownId {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = store.GlobalStoreRef.DeleteTask(taskId)
	if err != nil {
		// db was created as read-only
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
