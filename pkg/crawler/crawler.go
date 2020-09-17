package crawler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jackdanger/collectlinks"
)

//Crawl sends a get request and visits included link recursively
func Crawl(location string, depth int) (string, []string) {
	data := get(location)
	dataLinks := collectlinks.All(data.Body)
	defer data.Body.Close()
	output, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Println(err)
	}
	return string(output), dataLinks
}

func get(location string) *http.Response {
	resp, err := http.Get(location)
	if err != nil {
		log.Println(location, "->", err)
	}
	return resp
}
