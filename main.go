package main

import (
	"fmt"
	"log"
	"net/http"

	"go-rest-api/events"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/create/", events.CreateEvent)
	router.HandleFunc("/all/", events.GetAllEvents)
	log.Fatal(http.ListenAndServe(":8080", router))
}
