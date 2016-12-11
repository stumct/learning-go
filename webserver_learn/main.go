package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", indexHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	w.Write([]byte(id))
}
