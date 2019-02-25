package handlers

import (
	"encoding/json"
	"errors"
	"github.com/PrzemyslawMorski/json-api/dtos"
	"github.com/PrzemyslawMorski/json-api/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func UpdateTask(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if store.GlobalStoreRef == nil {
		w.WriteHeader(http.StatusInternalServerError)
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

	t := dtos.UpdateTaskDto{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&t)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, err := json.Marshal(dtos.ErrorResponseDto{
			Code:  http.StatusBadRequest,
			Error: "title has to be a non-empty string",
		})
		if err != nil {
			_, err = w.Write([]byte("{\"code\":404, \"error\":\"title has to be a non-empty string\"}"))
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = w.Write([]byte(response))
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	if strings.Trim(t.NewTitle, " ") == "" {
		log.Printf("User with ip %v wanted to update a task with empty title", r.RemoteAddr)
		w.WriteHeader(http.StatusBadRequest)

		response, err := json.Marshal(dtos.ErrorResponseDto{
			Code:  http.StatusBadRequest,
			Error: "title cannot be empty or all whitespace",
		})
		if err != nil {
			_, err = w.Write([]byte("{\"code\":404, \"error\":\"title cannot be empty or all whitespace\"}"))
			if err != nil {
				log.Fatal(err)
			}
		}

		_, err = w.Write([]byte(response))
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	_, err = store.GlobalStoreRef.UpdateTask(taskId, t.NewTitle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
