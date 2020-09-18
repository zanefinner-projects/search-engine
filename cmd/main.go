package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zanefinner-projects/search-engine/pkg/handlers"
)

func main() {
	//Crawl assets/sites
	//Crawl all 5 letter combinations
	//Use all links with robots.txt in mind ^
	//Create new go routine
	////Add to db using titles/metas
	////further crawl
	////further add

	//Serve search page
	//Search through db and deliver
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Index)
	log.Println("Serving at http://localhost:8776")
	panic(http.ListenAndServe("localhost:8776", router))

}
