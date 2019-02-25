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

func CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if store.GlobalStoreRef == nil {
		log.Fatal(errors.New("database was not initialized"))
	}

	t := dtos.CreateTaskDto{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&t)
	if err != nil {
		log.Printf("User with ip %v passed %+v to CreateTask", r.RemoteAddr, r.Body)

		w.WriteHeader(http.StatusBadRequest)
		response, err := json.Marshal(dtos.ErrorResponseDto{
			Code:  http.StatusBadRequest,
			Error: "title has to be a string",
		})
		if err != nil {
			_, err = w.Write([]byte("{\"code\":404, \"error\":\"title has to be a string\"}"))
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

	if strings.Trim(t.Title, " ") == "" {
		log.Printf("User with ip %v wanted to create a task with empty title", r.RemoteAddr)
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

	task, err := store.GlobalStoreRef.CreateTask(t.Title)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", "/"+strconv.Itoa(task.Id))
}
