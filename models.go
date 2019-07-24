package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetCurrentTimerHandler(w http.ResponseWriter, r *http.Request) {
	id := getCurrentTimerIndex()
	item, err := json.Marshal(store[id])

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%v\n", string(item))
}

func getCurrentTimerIndex() uuid.UUID {
	for i, timer := range store {
		if timer.Stop.IsZero() {
			return i
		}
	}

	return uuid.UUID{}
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	for _, value := range store {
		fmt.Fprintf(w, "%v\n", value)
	}
}

func GetOneHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := uuid.MustParse(vars["id"])
	fmt.Fprintf(w, "%v\n", store[id])
}

func StartTimerHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	var values map[string]string

	if err = json.Unmarshal(b, &values); err != nil {
		panic(err)
	}

	description := values["description"]

	if description == "" {
		return
	}

	newId := uuid.New()
	newTimer := &Timer{description, newId, time.Now(), time.Time{}}
	store[newId] = newTimer
	fmt.Fprintf(w, "%v\n", newTimer)
}

func StopTimer(w http.ResponseWriter, r *http.Request) {
	id := getCurrentTimerIndex()
	emptyId := uuid.UUID{}
	if id != emptyId {
		store[id].Stop = time.Now()
		body, err := json.Marshal(store[id])
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "%v\n", string(body))
	}
}

func DeleteTimer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := uuid.MustParse(vars["id"])
	delete(store, id)
	fmt.Fprintf(w, "Deleted %s\n", id)
}
