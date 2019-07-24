package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var store map[uuid.UUID]*Timer = make(map[uuid.UUID]*Timer)

type Timer struct {
	Description string    `json:"description"`
	Id          uuid.UUID `json:"id"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop,omitempty"`
}

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", GetAllHandler).Methods("GET")
	r.HandleFunc("/current", GetCurrentTimerHandler).Methods("GET")
	r.HandleFunc("/timer/{id}", GetOneHandler).Methods("GET")

	r.HandleFunc("/", StartTimerHandler).Methods("POST")
	r.HandleFunc("/stop", StopTimer).Methods("PUT")
	r.HandleFunc("/delete/{id}", DeleteTimer).Methods("POST")

	return r
}

func main() {
	router := SetupRouter()
	http.ListenAndServe(":8080", router)
}
