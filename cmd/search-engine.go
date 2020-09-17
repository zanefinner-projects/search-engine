package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zanefinner-projects/search-engine/pkg/crawler"
)

func main() {

	log.Println("Server Start")
	//Routes
	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/", root).Methods(http.MethodGet)
	mainRouter.HandleFunc("/search", search).Methods(http.MethodGet)

	//Serve Routes
	panic(http.ListenAndServe("localhost:8765", mainRouter))

}
func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>`)
}
func search(w http.ResponseWriter, r *http.Request) {

	defer log.Println("[Root] <- [Client]")
	io.WriteString(w, "<!DOCTYPE html>")
	q, ok := r.URL.Query()["q"]

	if !ok || len(q[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		io.WriteString(w, "Search something!")
	} else {
		log.Println("searching", q[0])
		content, links := crawler.Crawl(q[0], 77)
		io.WriteString(w, content)
		for _, link := range links {

			io.WriteString(w, "<a href='http://localhost:8765/search?q="+link+"'>"+link+"</a></br>")
		}
	}
}
