package handlers

import (
	"io"
	"net/http"
)

//Index serve index page
func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
	<!DOCTYPE html>
	<form>
		<input name="q" placeholder="search for something!"></input>
		<button type="submit">Search</button>
	</form>
	`)
}
