package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func UpdateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//if store.GlobalStoreRef == nil {
	//	log.Fatal(errors.New("database was not initialized"))
	//}

	_, _ = fmt.Fprint(w, "UpdateTask\n")
}
