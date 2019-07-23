package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var store map[string]*Timer = make(map[string]*Timer)
var currentId = 0

type Timer struct {
	Description string
	Start       string
	Stop        string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", GetAllHandler).Methods("GET")
	r.HandleFunc("/current", GetCurrentTimerHandler).Methods("GET")
	r.HandleFunc("/timer/{id}", GetOneHandler).Methods("GET")

	r.HandleFunc("/", StartTimerHandler).Methods("POST")
	r.HandleFunc("/stop", StopTimer).Methods("POST")
	r.HandleFunc("/delete/{id}", DeleteTimer).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}

func GetCurrentTimerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current timer")
	// id := getCurrentTimerIndex()
	// if id == "" {
	// 	fmt.Fprintln(w, "No current timer")
	// } else {
	// 	fmt.Fprintf(w, "Current timer is: %v\n", store[id])
	// }
}

func getCurrentTimerIndex() string {
	for i, timer := range store {
		if timer.Stop == "" {
			return i
		}
	}

	return ""
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, store)
}

func GetOneHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "%v\n", store[id])
}

func StartTimerHandler(w http.ResponseWriter, r *http.Request) {
	newTimer := &Timer{"Something", "Now", ""}
	store[string(currentId)] = newTimer
	fmt.Fprintf(w, "Started %v\n", newTimer)
	currentId++
}

func StopTimer(w http.ResponseWriter, r *http.Request) {
	id := getCurrentTimerIndex()
	if id == "" {
		fmt.Fprintln(w, "No running timer")
	} else {
		store[id].Stop = "Also now"
		fmt.Fprintln(w, "Stopped timer")
	}
}

func DeleteTimer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	delete(store, id)
	fmt.Fprintf(w, "Deleted %s\n", id)
}
