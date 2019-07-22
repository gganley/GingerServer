package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetAllHandler).Methods("GET")
	r.HandleFunc("/{id}", GetOneHandler).Methods("GET")

	r.HandleFunc("/", PostHandler).Methods("POST")
	r.HandleFunc("/stop", StopTimer).Methods("POST")
	r.HandleFunc("/delete/{id}", DeleteTimer).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GET")
}

func GetOneHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "%s %s\n", "GET", id)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST")
}

func StopTimer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Stopped timer")
}

func DeleteTimer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Delete %s\n", id)
}
