package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	memoryDatabase = &memDatabase{data: make(map[string]string)}
	infoLogger     = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	debugLogger    = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
)

func main() {
	r := mux.NewRouter()
	//Matches 8 string hex hash for a redirect
	r.HandleFunc("/{hash:[a-fA-F0-9]{8,8}}", RedirectHandler(memoryDatabase)).Methods(http.MethodGet)
	// Serve the home page
	r.Handle("/", http.FileServer(http.Dir(".")))
	// Serve static files
	statics := r.PathPrefix("/static/")
	statics.Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.ListenAndServe(":8080", r)
}
