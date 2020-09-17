package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

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
	io.WriteString(w, `
	<form action="/search" method="GET">
	<input name="q"></input>
	<button type="submit">Go!</button>`)
}
func search(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer log.Println("[Root] <- [Client]")
	io.WriteString(w, "<!DOCTYPE html>")
	q, ok := r.URL.Query()["q"]

	if !ok || len(q[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		io.WriteString(w, "Search something!")
	} else {
		log.Println("searching", q[0])
		content, links := crawler.Crawl(q[0], 77)
		log.Println("Retrieved in", "[", q[0], "]", time.Since(start))
		io.WriteString(w, content)
		for _, link := range links {
			linkSlc := strings.Split(link, "")
			if len(linkSlc) != 0 {

				if linkSlc := strings.Split(link, ""); linkSlc[0] != "/" {
					io.WriteString(w, "<a href='http://localhost:8765/search?q="+link+"'>"+link+"</a></br>")
				} else {

					io.WriteString(w, "<a href='http://localhost:8765/search?q="+q[0]+link+"'>"+link+"</a></br>")
				}
			}
		}
	}
}
