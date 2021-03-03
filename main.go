package main

import (
	"datasite/userprojects"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", userprojects.GetUserProjects).Methods("GET")
	fmt.Println("Server listening! Woohoo!")
	http.ListenAndServe(":7777", r)
}
