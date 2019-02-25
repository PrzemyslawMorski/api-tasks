package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func ReadTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if store.GlobalStoreRef == nil {
		log.Fatal(errors.New("database was not initialized"))
	}

	id := ps.ByName("id")
	if id == "" {
		ListTasks(w, r, ps)
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	task := store.GlobalStoreRef.GetTaskById(taskId)
	if task == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(task)
	if err != nil {
		taskJson := fmt.Sprintf("{\"id\":%v,\"title\":\"%v\"}", task.Id, task.Title)
		_, err = w.Write([]byte(taskJson))
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}
