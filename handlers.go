package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//RedirectHandler returns a function that handles all redirects
//A closure is used here so database can be mocked for tests
func RedirectHandler(db database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		trimHash := vars["hash"]
		// Try to get the actual URL for the hash
		actualURL, err := db.retrieve(trimHash)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Redirecting now to %s", actualURL)
	}
}
