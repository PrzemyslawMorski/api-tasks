package handlers

import (
	"errors"
	"fmt"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func DeleteTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if store.GlobalStoreRef == nil {
		log.Fatal(errors.New("database was not initialized"))
	}

	_, _ = fmt.Fprint(w, "DeleteTask\n")
}
