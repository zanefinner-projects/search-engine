package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Server Start")
	//Routes
	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/", root).Methods(http.MethodGet)

	//Serve Routes
	panic(http.ListenAndServe("localhost:8765", mainRouter))
}
func root(w http.ResponseWriter, r *http.Request) {
	defer log.Println("[Root] <- [Client]")
	io.WriteString(w, "root")
}
